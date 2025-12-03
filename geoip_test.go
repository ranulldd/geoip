package geoip

import (
	"testing"
)

func TestGetIPAddressRemote(t *testing.T) {
	ip := "111.111.111.111"
	ans := GetIPAddress(ip)
	t.Errorf("%v ", ans)
}
