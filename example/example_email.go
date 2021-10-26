package main

import (
	"flag"
	"fmt"

	"github.com/evalphobia/ipqualityscore-go/config"
	"github.com/evalphobia/ipqualityscore-go/ipqs"
)

// nolint
func main() {
	var email string
	flag.StringVar(&email, "email", "", "set target email address")
	flag.Parse()

	conf := config.Config{
		APIKey: "",
		Debug:  true,
	}

	svc, err := ipqs.New(conf)
	if err != nil {
		panic(err)
	}

	resp, err := svc.EmailValidation(email)
	fmt.Printf("[%+v]\n", resp)
	if err != nil {
		panic(err)
	}
}
