package monitor

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/miekg/dns"
)

func MonitorDns(path string, dnsType int, dnsServer string) (checkSuccess bool, response string, elapsed int64) {
	c := dns.Client{
		Timeout: 5 * time.Second,
	}
	m := dns.Msg{}
	m.SetQuestion(path+".", dns.TypeA)
	r, rtt, err := c.Exchange(&m, dnsServer+":53")
	if err != nil {
		fmt.Println(err)
		return false, "", 0
	}
	var dst []string
	for _, ans := range r.Answer {
		if dnsType == 1 {
			record, isType := ans.(*dns.A)
			if isType {
				dst = append(dst, record.A.String())
			}
		}
		if dnsType == 2 {
			record, isType := ans.(*dns.AAAA)
			if isType {
				dst = append(dst, record.AAAA.String())
			}
		}
		if dnsType == 3 {
			record, isType := ans.(*dns.CAA)
			if isType {
				dst = append(dst, record.String())
			}
		}
		if dnsType == 4 {
			record, isType := ans.(*dns.CNAME)
			if isType {
				dst = append(dst, record.String())
			}
		}
		if dnsType == 5 {
			record, isType := ans.(*dns.MX)
			if isType {
				dst = append(dst, record.String())
			}
		}
		if dnsType == 6 {
			record, isType := ans.(*dns.NS)
			if isType {
				dst = append(dst, record.String())
			}
		}
		if dnsType == 7 {
			record, isType := ans.(*dns.PTR)
			if isType {
				dst = append(dst, record.String())
			}
		}
		if dnsType == 8 {
			record, isType := ans.(*dns.SOA)
			if isType {
				dst = append(dst, record.String())
			}
		}
		if dnsType == 9 {
			record, isType := ans.(*dns.SRV)
			if isType {
				dst = append(dst, record.String())
			}
		}
		if dnsType == 10 {
			record, isType := ans.(*dns.TXT)
			if isType {
				dst = append(dst, record.String())
			}
		}

	}

	var checkSuccessBool bool = false
	if len(dst) > 0 {
		checkSuccessBool = true
	}
	responseByte, _ := json.Marshal(dst)
	return checkSuccessBool, string(responseByte), rtt.Milliseconds()
}
