package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"

	pb "github.com/campoy/justforfunc/say-grpc/api"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "output.wav", "wav file where the output will be written")
	flag.Parse()

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s: %v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)

	text := &pb.Text{Text: "hello, there"}
	res, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatalf("could not say %s: %v", text.Text, err)
	}

	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("could not write to %s: %v", *output, err)
	}
}
