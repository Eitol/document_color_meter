// Package main implements a client for Greeter service.
package pdf_to_image_converter

import (
	"context"
	"github.com/Eitol/document_color_meter/api_front/core/pdf_to_image_converter/pb"
	"log"
	"math"
	"time"

	"google.golang.org/grpc"
)

const (
	defaultAddress = "localhost:50052"
	defaultTimeout = time.Second * 20
)

type remoteConverter struct {
	ServerAddress string
	Timeout       time.Duration
}

func NewRemoteConverter(serverAddress string, timeout time.Duration) Converter {
	c := remoteConverter{
		ServerAddress: serverAddress,
		Timeout:       timeout,
	}
	c.init()
	return Converter(c)
}

func (client *remoteConverter) init() {
	if client.ServerAddress == "" {
		client.ServerAddress = defaultAddress
	}
	if client.Timeout == 0 {
		client.Timeout = defaultTimeout
	}
}

type imageFormat int32

const (
	ImageFormatUNKNOWN = imageFormat(0)
	ImageFormatJPEG    = imageFormat(1)
	ImageFormatPNG     = imageFormat(2)
	ImageFormatSVG     = imageFormat(3)
)

func (client remoteConverter) newGRPCClient() (*grpc.ClientConn, context.CancelFunc, pb.PdfToImageServiceClient, context.Context) {
	// Set up a connection to the server.

	conn, err := grpc.Dial(client.ServerAddress,
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxUint32),
			grpc.MaxCallSendMsgSize(math.MaxUint32)),
		grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewPdfToImageServiceClient(conn)
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

func (client remoteConverter) Convert(file []byte, format imageFormat) ([][]byte, error) {
	conn, cancel, c, ctx := client.newGRPCClient()
	defer func() {
		closeConnection(conn, cancel)
	}()
	r, err := c.Convert(ctx, &pb.PdfToImageRequest{
		File:   file,
		Format: pb.ImageFormat(format),
	})
	if err != nil || r == nil {
		return nil, err
	}
	return r.Pages, nil
}
