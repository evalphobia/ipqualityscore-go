package main

import (
	"flag"
	"fmt"

	"github.com/evalphobia/ipqualityscore-go/config"
	"github.com/evalphobia/ipqualityscore-go/ipqs"
)

// nolint
func main() {
	var ipaddr string
	flag.StringVar(&ipaddr, "ipaddr", "", "set target ip address")
	flag.Parse()

	conf := config.Config{
		APIKey: "",
		Debug:  true,
	}

	svc, err := ipqs.New(conf)
	if err != nil {
		panic(err)
	}

	resp, err := svc.IPReputation(ipaddr, ipqs.IPReputationOption{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36",
	})
	fmt.Printf("[%+v]\n", resp)
	if err != nil {
		panic(err)
	}
}
