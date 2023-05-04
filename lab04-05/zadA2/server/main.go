package main

import (
	// context "context"
	// "fmt"
	"log"
	"math"
	"net"
	"sync"

	// "sync"
	"time"

	"grpcproject/grpcproject"

	"google.golang.org/grpc"
)

var (
	runningTime   = int(time.Now().Unix())
	definedEvents [4]*notification
	clients       map[string][]grpcproject.Notification
	clientsMutex  sync.Mutex
)

type server struct {
	grpcproject.UnimplementedGrpcProjectServer
}

type notification struct {
	SubscribtionId int32
	Message        string
	Time           int32
	Times          []int32
	TestEnum       *grpcproject.TestEnum
}

// func (s *server) SayHello(ctx context.Context, in *grpcproject.HelloRequest) (*grpcproject.HelloReply, error) {
// 	fmt.Printf("Recived message: %s\n", in.GetName())
// 	return &grpcproject.HelloReply{Message: "Hello, " + in.GetName()}, nil
// }

func runNotification(
	in *grpcproject.SubscribeRequest,
	srv grpcproject.GrpcProject_SubscribeServer,
	wg *sync.WaitGroup,
	times []int32,
	event notification,
) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		resp := grpcproject.Notification{
			SubscribtionId: event.SubscribtionId,
			Message:        event.Message,
			Time:           event.Time,
			Times:          times,
			TestEnum:       *event.TestEnum,
		}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
			clientsMutex.Lock()
			clients[in.Name] = append(clients[in.Name], resp)
			clientsMutex.Unlock()
		}
	}()
}

func (s *server) Subscribe(in *grpcproject.SubscribeRequest, srv grpcproject.GrpcProject_SubscribeServer) error {

	log.Printf("Starting new client sub for ID: %d at %d\n", in.SubscribtionId, runningTime)

	var event = definedEvents[in.SubscribtionId]
	var clientName = in.Name
	var times []int32
	var wg sync.WaitGroup
	var oldRunningTime = 0
	clientsMutex.Lock()
	clients[clientName] = []grpcproject.Notification{}
	clientsMutex.Unlock()

	for {
		//clientsMutex.Lock()
		//if len(clients[clientName]) > 0 {
		//	resp := clients[clientName][0]
		//	if err := srv.Send(&resp); err != nil {
		//		//log.Printf("send error %v", err)
		//		clients[clientName] = append(clients[clientName], resp)
		//	} else {
		//		clients[clientName] = append(clients[clientName][:0], clients[clientName][1:]...)
		//	}
		//}
		//clientsMutex.Unlock()

		if runningTime%int(event.Time) == 0 && oldRunningTime != runningTime {
			log.Printf("Trying to send running time %v on subscription %v to client %v\n", runningTime, in.SubscribtionId, clientName)
			times = append(times, int32(runningTime))
			runNotification(in, srv, &wg, times, *event)
			oldRunningTime = runningTime
			// if unsubscribe
			// then break
		}
	}

	wg.Wait()
	return nil
}

func newNotification(SubscribtionId int32, Message string, Time int32, Times []int32, TestEnum *grpcproject.TestEnum) *notification {
	return &notification{
		SubscribtionId,
		Message,
		Time,
		Times,
		TestEnum,
	}
}

func main() {
	// based on https://itnext.io/build-grpc-server-with-golang-go-step-by-step-b3f5abcf9e0e
	clients = make(map[string][]grpcproject.Notification)

	log.Println("defining events")
	definedEvents[0] = newNotification(0, "event0", int32(math.Pow(2.0, 0)), []int32{}, new(grpcproject.TestEnum))
	definedEvents[1] = newNotification(1, "event1", int32(math.Pow(2.0, 1)), []int32{}, new(grpcproject.TestEnum))
	definedEvents[2] = newNotification(2, "event2", int32(math.Pow(2.0, 2)), []int32{}, new(grpcproject.TestEnum))
	definedEvents[3] = newNotification(3, "event2", int32(math.Pow(2.0, 3)), []int32{}, new(grpcproject.TestEnum))

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	log.Println("Start run timer")
	go func() {
		for {
			time.Sleep(1 * time.Second)
			runningTime = int(time.Now().Unix())
			log.Println(runningTime)
		}
	}()

	log.Println("Starting server …")
	s := grpc.NewServer()
	grpcproject.RegisterGrpcProjectServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
