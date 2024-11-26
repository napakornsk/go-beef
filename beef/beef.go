package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	beef "github.com/napakornsk/go-beef/proto"
	"google.golang.org/grpc"
)

func fetchFromExternal() (string, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (s *BeefServiceServer) GetAllBeef(req *empty.Empty, stream grpc.ServerStreamingServer[beef.GetAllBeefResponse]) error {
	fmt.Println("The GetAllBeef service was invoked...")
	res, err := fetchFromExternal()
	if err != nil {
		return fmt.Errorf("failed to fetch data from external API: %v", err)
	}
	words := strings.Fields(res)

	fmt.Println("Sending beef through stream...")
	for _, w := range words {
		if err := stream.Send(&beef.GetAllBeefResponse{
			Data: strings.Trim(w, ",."),
		}); err != nil {
			return fmt.Errorf("failed to a beef word through stream: %v", err)
		}
	}
	fmt.Println("SUCCESS: all beef have been sent through stream")

	return nil
}
