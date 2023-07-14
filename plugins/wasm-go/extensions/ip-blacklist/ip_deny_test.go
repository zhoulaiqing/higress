package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ip-tools"
	"testing"
)

func TestIpDeny(t *testing.T) {
	var ipIntervalRoot *ip_tools.IPIntervalNode

	minIp, maxIp, _ := ip_tools.GetIPIntRange("1.1.1.1")
	ipIntervalRoot = ip_tools.IPIntervalTreeInsert(ipIntervalRoot, ip_tools.IPInterval{minIp, maxIp})

	minIp2, maxIp2, _ := ip_tools.GetIPIntRange("2.2.2.2/16")
	ipIntervalRoot = ip_tools.IPIntervalTreeInsert(ipIntervalRoot, ip_tools.IPInterval{minIp2, maxIp2})

	minIp3, maxIp3, _ := ip_tools.GetIPIntRange("2001:0db8:85a3:0000:0000:8a2e:0000:0000")
	ipIntervalRoot = ip_tools.IPIntervalTreeInsert(ipIntervalRoot, ip_tools.IPInterval{minIp3, maxIp3})

	minIp4, maxIp4, _ := ip_tools.GetIPIntRange("2023:0db8::/32")
	ipIntervalRoot = ip_tools.IPIntervalTreeInsert(ipIntervalRoot, ip_tools.IPInterval{minIp4, maxIp4})

	// true
	res1 := ip_tools.IPSearch(ipIntervalRoot, minIp) != nil
	fmt.Println("Contains 1.1.1.1", res1)
	assert.True(t, res1)
	// false
	otherIp, _, _ := ip_tools.GetIPIntRange("10.23.2.2")
	res2 := ip_tools.IPSearch(ipIntervalRoot, otherIp) != nil
	fmt.Println("Contains 10.23.2.2", res2)
	assert.False(t, res2)
	// true
	ip3, _, _ := ip_tools.GetIPIntRange("2.2.1.1")
	res3 := ip_tools.IPSearch(ipIntervalRoot, ip3) != nil
	fmt.Println("Contains 2.2.1.1", res3)
	assert.True(t, res3)
	// true
	ip4, _, _ := ip_tools.GetIPIntRange("2001:0db8:85a3:0000:0000:8a2e:0000:0000")
	res4 := ip_tools.IPSearch(ipIntervalRoot, ip4) != nil
	fmt.Println("Contains 2001:0db8:85a3:0000:0000:8a2e:0000:0000", res4)
	assert.True(t, res4)
	// true
	ip5, _, _ := ip_tools.GetIPIntRange("2023:0db8:85a3:0000:0000:8a2e:0000:0000")
	res5 := ip_tools.IPSearch(ipIntervalRoot, ip5) != nil
	fmt.Println("Contains 2023:0db8:85a3:0000:0000:8a2e:0000:0000", res5)
	assert.True(t, res5)
}
