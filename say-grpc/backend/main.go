package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/Sirupsen/logrus"
	pb "github.com/campoy/justforfunc/say-grpc/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	logrus.Infof("listening to port %d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *port, err)
	}

	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("could not serve: %v", err)
	}
}

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	return nil, fmt.Errorf("not implemented")
}

// cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")
// cmd.Stdout = os.Stdout
// cmd.Stderr = os.Stderr
// if err := cmd.Run(); err != nil {
// 	log.Fatal(err)
// }
