package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/beevik/ntp"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	var hosts string
	var timeoutSecond int
	var debug bool
	flag.StringVar(&hosts, "hosts", "time.cloudflare.com,time.aws.com", "comma separated ntp hosts")
	flag.IntVar(&timeoutSecond, "timeout", 5, "timeout seconds")
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()

	if debug {
		log.Printf("hosts: %+v, timeout: %+v, debug: %+v", hosts, timeoutSecond, debug)
	}

	hostsSlice := strings.Split(hosts, ",")

	offset, err := offset(hostsSlice, timeoutSecond)
	if err != nil {
		log.Fatalf("ERROR: offset failed. error: %+v", err)
	}

	if debug {
		log.Printf("offset: %+v ms\n", offset)
	}
	fmt.Printf(`{"clock.offset.ms": %v}`, offset)

}

func offset(hosts []string, timeoutSecond int) (offset float64, err error) {
	for _, host := range hosts {
		options := ntp.QueryOptions{Timeout: time.Duration(timeoutSecond) * time.Second}
		response, err := ntp.QueryWithOptions(strings.Trim(host, " "), options)
		if err != nil {
			log.Printf("ERROR: ntp query failed. host: %+v, response: %+v, error: %+v", host, response, err)
			continue
		}
		return float64(response.ClockOffset.Milliseconds()), nil
	}
	return offset, nil
}
