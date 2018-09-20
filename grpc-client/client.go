package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/iamgoangle/go-grpc-kafka-hbase/api/proto/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "topic.json"
)

func parseFile(file string) (*pb.Topic, error) {
	var topic *pb.Topic
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &topic)
	return topic, err
}
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewTopicServiceClient(conn)
	file := defaultFilename

	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	topic, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}
	r, err := client.CreateTopic(context.Background(), topic)

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetTopics(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list topics: %v", err)
	}

	for _, v := range getAll.Topics {
		log.Println(v)
	}
}
