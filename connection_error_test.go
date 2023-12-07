package grpcstatusunknown

import (
	context "context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	status "google.golang.org/grpc/status"
)

type simpleService struct {
	t *testing.T
	UnimplementedSimpleServiceServer
}

var synCh chan bool = make(chan bool)

func (s *simpleService) Subscribe(req *SimpleRequest, stream SimpleService_SubscribeServer) error {
	defer func() {
		synCh <- true // Signal test body to continue
	}()
	<-synCh // Wait signal to continue
	err := stream.Send(&SimpleResponse{Text: ""})
	st, _ := status.FromError(err)
	if !assert.NotEqual(s.t, codes.Unknown, st.Code()) {
		s.t.Fatal(err)
	}
	return err
}

func setupClient(lisGrpc net.Listener, t *testing.T) (*grpc.ClientConn, SimpleService_SubscribeClient) {
	conn, err := grpc.Dial(lisGrpc.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	require.NotNil(t, conn)

	client := NewSimpleServiceClient(conn)
	subClient, err := client.Subscribe(context.Background(), &SimpleRequest{})
	require.NoError(t, err)
	return conn, subClient
}

func setupServer(t *testing.T) net.Listener {
	lisGrpc, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("listen grpc: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterSimpleServiceServer(grpcServer, &simpleService{t: t})
	go func() {
		grpcServer.Serve(lisGrpc)
	}()
	return lisGrpc
}

func TestGrpcTransmitClosingPackagingServices(t *testing.T) {
	lisGrpc := setupServer(t)
	conn, subClient := setupClient(lisGrpc, t)

	go func() {
		time.Sleep(10 * time.Millisecond) // Wait things settle down.
		err := conn.Close()
		assert.NoError(t, err)
		time.Sleep(1 * time.Millisecond) // Wait a little bit for closing the connection.
		synCh <- true                    // Signal server to send first message on the stream.
	}()

	var err error
	for err == nil {
		_, err = subClient.Recv()
	}
	st, _ := status.FromError(err)
	if !assert.NotEqual(t, codes.Canceled, st) {
		t.Log(err)
	}
	<-synCh // Wait signal from the server
}
