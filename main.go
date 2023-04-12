package main

import (
	"context"
	"flag"
	"fmt"
	"log"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "listen address for service")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))
	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()

	price, err := svc.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(price)
}
