package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tcnksm/go-httpstat"
)

func main() {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	// Create a httpstat powered context
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)
	// Send request by default HTTP client
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	
	// Show the results
	log.Printf("DNS lookup: %d ms", int(result.DNSLookup/time.Millisecond))
	log.Printf("TCP connection: %d ms", int(result.TCPConnection/time.Millisecond))
	log.Printf("TLS handshake: %d ms", int(result.TLSHandshake/time.Millisecond))
	log.Printf("Server processing: %d ms", int(result.ServerProcessing/time.Millisecond))
	log.Printf("Content transfer: %d ms", int(result.ContentTransfer(time.Now())/time.Millisecond))
}
