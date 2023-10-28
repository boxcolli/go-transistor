package mockserver

import (
	"sync"
	"math/rand"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/server"
	"github.com/boxcolli/go-transistor/types"
)

type mockServer struct {
	chsiz	int	// Channel size
	ch		map[int]chan *types.Message // Subscribe channel
	chmx	sync.RWMutex

	pb.UnimplementedTransistorServiceServer
}

func NewMockServer(chsiz int) server.Server {
	return &mockServer{
		chsiz: chsiz,
	}
}

// Command implements server.Server.
func (*mockServer) Command(*pb.CommandRequest, pb.TransistorService_CommandServer) error {
	panic("unimplemented")
}

// Publish implements server.Server.
func (s *mockServer) Publish(stream pb.TransistorService_PublishServer) error {
	for {
		// Get message
		req, err := stream.Recv()
		if err != nil {
			break
		}

		// Publish message
		s.chmx.RLock()
		for _, ch := range s.ch {
			msg := new(types.Message)
			msg.Unmarshal(req.GetMsg())
			ch <- msg
		}
		s.chmx.RUnlock()
	}

	return nil
}

// Subscribe implements server.Server.
func (s *mockServer) Subscribe(stream pb.TransistorService_SubscribeServer) error {
	// Put my subscribe channel
	ch := make(chan *types.Message, s.chsiz)
	var me int
	{
		s.chmx.Lock()
		for {
			me = rand.Int()
			if _, ok := s.ch[me]; !ok {
				s.ch[me] = ch
				break
			}
		}
		s.chmx.Unlock()
	}

	// Read channel
	for {
		msg := <- ch

		// Send message
		err := stream.Send(&pb.SubscribeResponse{
			Msg: msg.Marshal(),
		})
		if err != nil {
			break
		}
	}

	// Delete my channel
	{
		s.chmx.Lock()
		delete(s.ch, me)
		s.chmx.Unlock()
		close(ch)
	}

	return nil
}
