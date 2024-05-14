package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	pb "7-solutions/beef_service"
	"unicode"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBeefServiceServer
}

func (s *server) GetBeefSummary(ctx context.Context, req *pb.BeefRequest) (*pb.BeefResponse, error) {
	data, err := fetchBeefData()
	if err != nil {
		return nil, err
	}

	response := countBeefOccurrences(data)

	return &pb.BeefResponse{WordCounts: response}, nil
}

func fetchBeefData() (string, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func countBeefOccurrences(input string) map[string]int32 {
	words := make(map[string]int32)
	currentWord := ""

	for _, char := range input {
		if unicode.IsLetter(char) || char == '-' {
			currentWord += string(unicode.ToLower(char))
		} else if currentWord != "" {
			words[currentWord] += 1
			currentWord = ""
		}
	}

	if currentWord != "" {
		words[currentWord] += 1
	}

	return words
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBeefServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
