package main

import (
	"context"
	"flag"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	pb "github.com/campoy/justforfunc/12-say-grpc/api"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "output.wav", "path of the output file")
	flag.Parse()

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect to %s: %v", *backend, err)
	}
	defer conn.Close()

	text := "hello"

	client := pb.NewTextToSpeechClient(conn)
	res, err := client.Say(context.Background(), &pb.Text{Text: text})
	if err != nil {
		logrus.Fatalf("could not say %s: %v", text, err)
	}

	if err := ioutil.WriteFile(*output, res.Sound, 0666); err != nil {
		logrus.Fatalf("could not write to %s: %v", *output, err)
	}
}
