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
	fqdn := strings.ToLower(m.Question[0].Name)
    rr := handler.DNSHandler(fqdn, r.Question[0].Qtype)
    m.Answer = []dns.RR{rr}
	w.WriteMsg(m)
}
