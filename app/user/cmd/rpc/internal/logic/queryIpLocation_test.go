package logic

import (
	"fmt"
	"testing"
)

func TestQueryIpLocation(t *testing.T) {
	ips := []string{"24.48.0.1", "46.232.121.174", "36.232.43.210", "115.191.200.34", "233"}
	for _, ip := range ips {
		fmt.Println(queryIpLocation(ip))
	}
}
