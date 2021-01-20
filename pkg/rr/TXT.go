package rr

import (
	"github.com/miekg/dns"
)

func TypeTXT(fqdn string, value string) (Rr *dns.TXT) {
	Rr = new(dns.TXT)
	Rr.Hdr = dns.RR_Header{
		Name:   fqdn,
		Rrtype: dns.TypeTXT,
		Class:  dns.ClassINET,
		Ttl:    3600}
	Rr.Txt[0] = value
	return Rr
}
