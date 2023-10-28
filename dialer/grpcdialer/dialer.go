package grpcdialer

import (
	"context"
	"sync"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/dialer"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/io/reader/grpcreader"
	"github.com/boxcolli/go-transistor/types"
	"google.golang.org/grpc"
)

type entry struct {
	conn	*grpc.ClientConn
	cl		pb.TransistorServiceClient
	sub		pb.TransistorService_SubscribeClient
	sr		io.StreamReader
}

type grpcDialer struct {
	// grpc request information
	opts []grpc.DialOption

	// connection holder
	e map[*types.Member]*entry
	emx sync.RWMutex
}

func NewGrpcDialer(opts []grpc.DialOption) dialer.Dialer {
	return &grpcDialer{
		opts: opts,
	}
}

// Dial implements dialer.Dialer.
func (d *grpcDialer) Dial(ctx context.Context, m *types.Member, c *types.Change) (io.StreamReader, error) {
	d.emx.Lock()
	defer d.emx.Unlock()

	if ent, ok := d.e[m]; ok {
		return ent.sr, nil
	}

	var ent = new(entry)
	{
		var err error
		// Create connection
		ent.conn, err = grpc.Dial(m.Address(), d.opts...)
		if err != nil {
			return nil, err
		}

		// Create client
		ent.cl = pb.NewTransistorServiceClient(ent.conn)

		// Call RPC
		ent.sub, err = ent.cl.Subscribe(ctx)	// blocking function
		if err != nil {
			return nil, err
		}

		// Send initial Change
		if c != nil {
			req := &pb.SubscribeRequest{
				Change: c.Marshal(),
			}
			err := ent.sub.Send(req)
			if err != nil {
				return nil, err
			}
		}

		// Create a StreamReader
		ent.sr = grpcreader.NewGrpcClientStream(ent.sub)
	}

	d.e[m] = ent
	
	return ent.sr, nil
}

func (d *grpcDialer) Apply(m *types.Member, c *types.Change) error {
	d.emx.RLock()
	defer d.emx.RUnlock()

	ent, ok := d.e[m]
	if !ok {
		return dialer.ErrMemberNotFound
	}

	return ent.sub.Send(&pb.SubscribeRequest{ Change: c.Marshal() })
}

// Stop implements dialer.Dialer.
func (d *grpcDialer) Close(m *types.Member) error {
	d.emx.Lock()
	defer d.emx.Unlock()

	ent, ok := d.e[m]
	if !ok {
		// The connection is already closed.
		return nil
	} else {
		// Close connection
		err := ent.conn.Close()
		if err != nil {
			return err
		}
		delete(d.e, m)
	}
	
	return nil
}

// StopAll implements dialer.Dialer.
func (d *grpcDialer) CloseAll() {
	d.emx.Lock()
	defer d.emx.Unlock()

	// Close connection
	for _, v := range d.e {
		v.conn.Close()
	}

	// Wipe out member set
	d.e = map[*types.Member]*entry{}
}
