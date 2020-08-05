package query

import (
	_ "fmt"
	_ "net"
	"strings"

	"github.com/TsundereChen/geodns-go/pkg/config"
	"github.com/TsundereChen/geodns-go/pkg/handler"
	"github.com/miekg/dns"
)

func RegisterDomain() {
	var domainList = config.FetchDomain(config.ConfigMap)
	for i := range domainList {
		if domainList[i] != "regions" {
			dns.HandleFunc(domainList[i], HandleFunction)
		}
	}
}

func HandleFunction(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true
	fqdn := m.Question[0].Name
	fqdn = strings.ToLower(fqdn)
	switch r.Question[0].Qtype {
	case dns.TypeTXT:
		rr := handler.DNSHandler(fqdn, "TXT")
		m.Answer = []dns.RR{rr}
		break
	case dns.TypeA:
		rr := handler.DNSHandler(fqdn, "A")
		m.Answer = []dns.RR{rr}
		break
	case dns.TypeAAAA:
		rr := handler.DNSHandler(fqdn, "AAAA")
		m.Answer = []dns.RR{rr}
		break
	case dns.TypeCNAME:
		rr := handler.DNSHandler(fqdn, "CNAME")
		m.Answer = []dns.RR{rr}
		break
	default:
		break
	}
	w.WriteMsg(m)
}
