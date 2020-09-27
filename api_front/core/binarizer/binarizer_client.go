package binarizer

import (
	"context"
	"math"
	"time"

	pb "github.com/Eitol/document_color_meter/api_front/core/binarizer/pb"

	"google.golang.org/grpc"
	"log"
)

type Binarizer interface {
	Binarize(images [][]byte, method BinarizationMethod) ([][]byte, string, error)
}

func NewBinarizerInstance(serverAddress string, timeout time.Duration) Binarizer {
	return &remoteBinarizer{
		ServerAddress: serverAddress,
		Timeout:       timeout,
	}
}

type remoteBinarizer struct {
	ServerAddress string
	Timeout       time.Duration
}

func (client remoteBinarizer) newGRPCClient() (*grpc.ClientConn, context.CancelFunc, pb.BinarizerClient, context.Context) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(client.ServerAddress, grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxUint32),
		grpc.MaxCallSendMsgSize(math.MaxUint32)), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewBinarizerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), client.Timeout)
	return conn, cancel, c, ctx
}

func closeConnection(conn *grpc.ClientConn, cancel context.CancelFunc) {
	err := conn.Close()
	if err != nil {
		log.Print(err)
	}
	cancel()
}

type BinarizationMethod int

const (
	UNDEFINED = BinarizationMethod(0)
	SAUVOLA   = BinarizationMethod(1)
	NIBLACK   = BinarizationMethod(2)
	OTSU      = BinarizationMethod(3)
)

func (client remoteBinarizer) Binarize(images [][]byte, method BinarizationMethod) ([][]byte, string, error) {
	conn, cancel, c, ctx := client.newGRPCClient()
	defer func() {
		closeConnection(conn, cancel)
	}()
	r, err := c.Binarize(ctx, &pb.BinarizeRequest{
		Images:             images,
		BinarizationMethod: pb.BinarizationMethod(method),
	})
	if err != nil || r == nil {
		return nil, "", err
	}
	return r.Images, r.OutPath, nil
}
