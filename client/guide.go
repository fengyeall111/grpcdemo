package main

import (
	"context"
	"google.golang.org/grpc"
	"grpcdemo/pb"
	"io"
	"log"
	"time"
)

func main()  {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c:=pb.NewGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.S1R1(ctx, &pb.GuideRequest{Name:"s1r1:client"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResult())
	stream, err1 := c.S1Rn(ctx, &pb.GuideRequest{Name:"s1rn:client"})
	if err1 != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf(".ListFeatures(_) = _, %v", err)
		}
		log.Println(res)
	}
	stream1, err := c.SnR1(context.Background())
	var reqs []*pb.GuideRequest
	for i := 0; i < 10; i++ {
         reqs = append(reqs,&pb.GuideRequest{Name: "snr1"})
	}
	log.Printf("Traversing %d points.", len(reqs))
	snrnclient, err := c.SnRn(context.Background())
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", snrnclient, err, nil)
	}
	waitc := make(chan struct{})
	go func() {
		for  {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("in:%v",in)
		}


	}()
	for _, req := range reqs {
		if err := stream1.Send(req); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream1, req, err)
		}
		if err := snrnclient.Send(req); err != nil {
			log.Fatalf("%v.Send(%v) = %v", snrnclient, req, err)
		}

	}
	reply, err := stream1.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)
	snrnclient.CloseSend()
	<-waitc
}
