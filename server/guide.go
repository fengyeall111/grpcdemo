package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcdemo/pb"
	"io"
	"log"
	"net"
)
type guideServer struct {
	pb.UnimplementedGuideServer
}

func (receiver *guideServer) S1R1(ctx context.Context,req *pb.GuideRequest) (*pb.GuideResponse, error){
		log.Printf("Received: %v", req.GetName())
	    return &pb.GuideResponse{Result: "s1r1"}, nil
}
func (receiver *guideServer) S1Rn(req *pb.GuideRequest,stream pb.Guide_S1RnServer)error{
	log.Printf("Received: %v", req.GetName())
	for i := 0; i < 10; i++ {
		if err:= stream.Send(&pb.GuideResponse{Result: "s1rn"});err!=nil{
			return err
		}
	}
	return nil
	// return &pb.GuideResponse{Result: "s1rn"}, nil
}

func (receiver *guideServer) SnR1(stream pb.Guide_SnR1Server)error{
	// log.Printf("Received: %v", req.GetName())
	for{
		req,err := stream.Recv()
		fmt.Println(req)
		if err == io.EOF{
			return stream.SendAndClose(&pb.GuideResponse{Result: "snrn"})
		}
		if err != nil {
			return err
		}
		fmt.Println(req)
	}
}
func (receiver *guideServer) SnRn(stream pb.Guide_SnRnServer)error{
	// log.Printf("Received: %v", req.GetName())
	for{
		in,err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println(in)
		stream.Send(&pb.GuideResponse{Result: "snrn"})
	}
}


func main()  {
	lis, err := net.Listen("tcp", ":50052")
	if err!= nil{
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGuideServer(s,&guideServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}