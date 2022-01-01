package main

import (
	"context"
	"fmt"
	"time"
)

func exeA1(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("A finish")
			return
		default:
			fmt.Printf("Worker A exe %d seconds: \n", i)
		}
		i++
	}
}

func exeB1(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("B finish")
			return
		default:
			fmt.Printf("Worker B exe %d seconds: \n", i)
		}
		i++
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4 * time.Second)

	go exeA1(ctx)
	go exeB1(ctx)

	//模拟程序运行 - Sleep 5秒
	time.Sleep(6 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("end!!!!")
}

