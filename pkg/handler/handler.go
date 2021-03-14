package handler

import (
	"github.com/TsundereChen/geodns-go/pkg/config"
	"github.com/TsundereChen/geodns-go/pkg/fetch"
	"github.com/miekg/dns"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func DNSHandler(fqdn string, questionType uint16, sourceAddress net.IP, IPv4 bool) (rr dns.RR) {
	// Get the subdomain information first
	record, err := config.GeoDB.Country(sourceAddress)
	if err != nil {
		log.Panic(err)
	}
	if *(config.Debug) == true {
		log.Printf("Source IP => %s\n", sourceAddress.String())
		log.Printf("Source Country => %s\n", record.Country.IsoCode)
		log.Printf("handler.DNSHandler handling request %s, question type %s\n", fqdn, dns.TypeToString[questionType])
	}
	var value string
	value = ""
	for k := range config.ConfigMap {
		if strings.Contains(fqdn, k) {
			// Split FQDN into domain and subdomain
			subdomain := fetch.FetchSubDomainName(fqdn, k)
			if *(config.Debug) == true {
				log.Printf("Searching subdomain %s in %s\n", subdomain, k)
			}
			// Now find if record exists.
			rrData := fetch.FetchRR(config.ConfigMap[k])
			for rrName := range rrData {
				if rrName == subdomain {
					// Check if rr type match
					rrType := fetch.FetchRrType(rrData[rrName])
					if typeChecker(rrType, questionType) {
						// Match
						if record.Country.IsoCode == "" {
							// If country ISO code is empty
							// Return default value
							value = fetch.FetchDefaultValue(rrData[rrName])
						} else if config.ServerMapping[record.Country.IsoCode] == "" {
							// If there's no value in ServerMapping
							// Also return default value
							value = fetch.FetchDefaultValue(rrData[rrName])
						} else {
							// Locate the region first
							region := config.ServerMapping[record.Country.IsoCode]
							// Check if there's any rule matches the request
							rules := fetch.FetchRules(rrData[rrName])
							// Set value as default value first
							// Overwrite the default value if we found matching rule
							value = fetch.FetchDefaultValue(rrData[rrName])
							for rule := range rules {
								recordRegion := rules[rule].(map[interface{}]interface{})["region"].(string)
								if region != recordRegion {
									// If region doesn't match
									// Then check next rule
									continue
								}
								ruleTime := rules[rule].(map[interface{}]interface{})["time"].(string)
								timeArr := strings.Split(ruleTime, "-")
								timeStart, _ := strconv.Atoi(timeArr[0])
								timeEnd, _ := strconv.Atoi(timeArr[1])
								currHour, _, _ := time.Now().Clock()
								if timeStart <= currHour && currHour <= timeEnd {
									// Time is valid, check the record and return the result
									hashString := config.Hash(rrName.(string) + k + recordRegion + ruleTime)
									value = config.WeightedRR[hashString].Next().(string)
								} else {
									continue
								}
							}
						}
					} else {
						continue
					}
				}
			}
		}
	}
	rr = RrGenerator(questionType, fqdn, value)
	return rr
}

func typeChecker(rrType string, questionType uint16) (res bool) {
	return rrType == dns.TypeToString[questionType]
}
