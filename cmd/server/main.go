package main

import (
	"context"
	"flag"
	"log"
	"net"
	// "fmt"

	// pb "go-test-002/grpc"
	//hellopb "server-test/pkg/grpc"

	// hellopb "github.com/yuuki100/server-test/pkg/grpc"
	pb "github.com/yuuki100/server-test/pkg/grpc"

	// "github.com/yuuki100/server-test/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}


// func (s *Service) TemporaryYuuki(ctx context.Context, r *a.Empty) (*a.Empty, error) {
// 	// fmt.Println("============================= TEST START =========================")
// 	// return &a.Empty{}, nil
// }

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// リフレクションの追加
	reflection.Register(s)

	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}



// func NewMyServer() *myServer {
// 	return &myServer{}
// }

// func main() {
// 	// 1. 8080番portのLisnterを作成
// 	port := 50010
// 	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 2. gRPCサーバーを作成
// 	s := grpc.NewServer()

// 	// 3. gRPCサーバーにGreetingServiceを登録
// 	hellopb.RegisterGreeterServer(s, NewMyServer())

// 	// リフレクションの追加
// 	reflection.Register(s)

// 	// 4. 作成したgRPCサーバーを、8080番ポートで稼働させる
// 	go func() {
// 		log.Printf("start gRPC server port: %v", port)
// 		s.Serve(listener)
// 	}()

	
	// 5.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// log.Println("stopping gRPC server...")
	// s.GracefulStop()





	// flag.Parse()
	// //lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	// lis, err := net.Listen("tcp", "localhost:50051")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// s := grpc.NewServer()

	// // リフレクションの追加
	// reflection.Register(s)

	// pb.RegisterGreeterServer(s, &server{})
	// log.Printf("server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
//}
