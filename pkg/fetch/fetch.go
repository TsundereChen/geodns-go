package fetch

import (
	"github.com/miekg/dns"
	"strings"
)

var (
	ConfigMap map[string]interface{}
)

func FetchRR(config interface{}) (rrData map[interface{}]interface{}) {
	return config.(map[interface{}]interface{})["rr"].(map[interface{}]interface{})
}

func FetchSubDomainName(fqdn string, domain string) (subDomain string) {
	return strings.Split(strings.Split(fqdn, domain)[0], ".")[0]
}

func FetchRrType(rrData interface{}) (rrType string) {
	return rrData.(map[interface{}]interface{})["type"].(string)
}

func FetchDefaultValue(rrData interface{}) (value string) {
	return rrData.(map[interface{}]interface{})["default"].(string)
}

func FetchDNSType(requestType string) (rrType uint16) {
	switch strings.ToUpper(requestType) {
	case "A":
		return dns.TypeA
	case "AAAA":
		return dns.TypeAAAA
	case "CNAME":
		return dns.TypeCNAME
	case "TXT":
		return dns.TypeTXT
	default:
		return dns.TypeNone
	}
}
