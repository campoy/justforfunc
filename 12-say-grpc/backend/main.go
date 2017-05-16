package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/Sirupsen/logrus"
	pb "github.com/campoy/justforfunc/12-say-grpc/api"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Infof("listening on port %d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, new(server))
	logrus.Fatal(s.Serve(lis))
}

type server struct{}

func (s *server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create tmp file: %v", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close newly created tmp file: %v", err)
	}
	defer os.Remove(f.Name())

	out, err := exec.Command("flite", "-t", text.Text, "-o", f.Name()).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("flite failed: %s", out)
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read tmp file %s: %v", f.Name(), err)
	}
	logrus.Debugf("read %d bytes from %s", len(data), f.Name())

	return &pb.Speech{Sound: data}, nil
}
