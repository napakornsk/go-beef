package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	beef "github.com/napakornsk/go-beef/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BeefRes struct {
	Beef map[string]int32
}

type BeefService struct {
	client beef.BeefServiceClient
}

func InitBeefService(grpcAddress string) (*BeefService, error) {
	conn, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	client := beef.NewBeefServiceClient(conn)

	return &BeefService{client: client}, nil
}

func (s *BeefService) GetBeefMap(c *gin.Context) {
	stream, err := s.client.GetAllBeef(context.Background(), &empty.Empty{})
	if err != nil {
		log.Printf("failed to call GetAllBeef: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to call GetAllBeef: %v", err)})
		return
	}

	beefMap := make(map[string]int32)
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error while receiving from stream: %v\n", err)
		}

		amount, isFound := beefMap[res.Data]
		if isFound {
			beefMap[res.Data] = amount + 1
		} else {
			beefMap[res.Data] = 1
		}
	}

	c.JSON(http.StatusOK, BeefRes{Beef: beefMap})
}
