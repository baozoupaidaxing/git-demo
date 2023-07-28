package grpc

import (
	"context"
	"fmt"
	pb "goProject/student/lib/grpc/server"
	"goProject/student/server/models/psql"
	"google.golang.org/grpc"
	"net"
)

var grpConn *grpc.Server

type Student struct {
}

func Start(addr string) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("failed to listen: %v", err)
		panic(err)
	}

	grpConn = grpc.NewServer()

	pb.RegisterServerServer(grpConn, &Student{})
	fmt.Println("listen to ", addr)
	grpConn.Serve(listen)
}

func Stop() {
	if grpConn != nil {
		grpConn.Stop()
	}
}

func (s *Student) GetAgeById(ctx context.Context, in *pb.SIdReq) (*pb.AgeRes, error) {

	sId := in.SId
	student, err := psql.SelectById(int(sId))
	if err != nil {
		return nil, err
	}

	return &pb.AgeRes{SAge: int64(student.Age)}, nil
}
