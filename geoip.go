package geoip

import (
	"log"
	"net/netip"
	"strings"

	"github.com/oschwald/maxminddb-golang/v2"
)

var cityDB, _ = maxminddb.Open("GeoLite2-City.mmdb")
var cnDB, _ = maxminddb.Open("GeoCN.mmdb")

func GetIPAddress(ip string) string {
	addr := netip.MustParseAddr(ip)
	var cityRecord struct {
		Country struct {
			Names map[string]string `maxminddb:"names"`
		} `maxminddb:"country"`
	}
	err := cityDB.Lookup(addr).Decode(&cityRecord)
	if err != nil {
		log.Println(err)
		return ""
	}
	ans := cityRecord.Country.Names["zh-CN"]
	if ans != "中国" {
		return ans
	}

	var cnRecord struct {
		Province string `maxminddb:"province"`
		City     string `maxminddb:"city"`
		ISP      string `maxminddb:"isp"`
	}
	err = cnDB.Lookup(addr).Decode(&cnRecord)
	if err != nil {
		log.Println(err)
		return ""
	}
	ans = strings.Join([]string{cnRecord.Province, cnRecord.City, cnRecord.ISP}, "|")

	return ans
}
