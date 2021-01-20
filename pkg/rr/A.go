package rr

import (
	"github.com/miekg/dns"
	"net"
)

func TypeA(fqdn string, value string) (Rr *dns.A) {
	Rr = new(dns.A)
	Rr.Hdr = dns.RR_Header{
		Name:   fqdn,
		Rrtype: dns.TypeA,
		Class:  dns.ClassINET,
		Ttl:    3600}
	Rr.A = net.ParseIP(value)
	return Rr
}
