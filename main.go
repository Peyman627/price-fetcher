package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Peyman627/price-fetcher/client"
	"github.com/Peyman627/price-fetcher/proto"
	"log"
	"time"
)

func main() {
	//cl := client.New("http://localhost:3000")
	//
	//price, err := cl.FetchPrice(context.Background(), "BTC")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%+v\n", price)
	//return

	var (
		jsonAddr = flag.String("jsonAddr", ":3000", "listen address of the json transport")
		grpcAddr = flag.String("grpcAddr", ":4000", "listen address of the grpc transport")
		svc      = NewLoggingService(NewMetricService(&priceFetcher{}))
		ctx      = context.Background()
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		time.Sleep(3 * time.Second)
		resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", resp)
	}()

	go makeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()
}
