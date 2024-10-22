package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, h asMx, hasSPF, sprRecord, h asDMARC, dmarcRecord\n")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	err := scanner.Err()
	handleError(err)
}

func checkDomain(domain string) {
	var haxMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	if len(mxRecords) > 0 {
		haxMX = true
	}
	txtRecords, err := net.LookupTXT(domain)
	handleError(err)
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	handleError(err)

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Println(domain, haxMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

func handleError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
