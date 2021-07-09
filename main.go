package main

import (
	"log"
	"net"
	"strings"

	"github.com/miekg/dns"
)

type handler struct{}

func (this *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := &dns.Msg{}
	msg.SetReply(r)
	log.Println("Recieved query for " + r.Question[0].String())
	if strings.HasSuffix(strings.TrimRight(r.Question[0].Name, "."), ".docker") && r.Question[0].Qtype == dns.TypeA {
		msg.Authoritative = true
		domain := msg.Question[0].Name
		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
			A:   net.ParseIP("127.0.0.1"),
		})
	} else {
		msg, _ = dns.Exchange(r, "1.1.1.1:53")
	}
	w.WriteMsg(msg)
}

func main() {
	addr := "127.0.0.1:53"
	srv := &dns.Server{Addr: addr, Net: "udp"}
	srv.Handler = &handler{}
	log.Println("Starting DNS server on " + addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
