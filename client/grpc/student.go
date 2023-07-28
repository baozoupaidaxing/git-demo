package grpc

import (
	"context"
	"fmt"
	pb "goProject/student/lib/grpc/server"
	"google.golang.org/grpc"
)

func GetAgeById(sId int) (int, error) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer conn.Close()

	c := pb.NewServerClient(conn)
	idReq := &pb.SIdReq{SId: int64(sId)}
	ageRes, err := c.GetAgeById(context.Background(), idReq)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int(ageRes.SAge), nil
}
