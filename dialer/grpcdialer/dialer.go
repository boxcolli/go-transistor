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
	ctx context.Context
	opts []grpc.DialOption

	// connection holder
	e map[*types.Member]*entry
	emx sync.Mutex
}

// Dial implements dialer.Dialer.
func (d *grpcDialer) Dial(m *types.Member) (io.StreamReader, error) {
	var ent = new(entry)
	{
		var err error
		ent.conn, err = grpc.Dial(m.Address(), d.opts...)
		if err != nil {
			return nil, err
		}
		ent.cl = pb.NewTransistorServiceClient(ent.conn)
		ent.sub, err = ent.cl.Subscribe(d.ctx)	// blocking function
		if err != nil {
			return nil, err
		}
		ent.sr = grpcreader.NewGrpcClientStream(ent.sub)
	}

	d.emx.Lock()
	d.e[m] = ent
	d.emx.Unlock()
	return ent.sr, nil
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

func NewGrpcDialer(ctx context.Context, opts []grpc.DialOption) dialer.Dialer {
	return &grpcDialer{
		ctx: ctx,
		opts: opts,
	}
}
