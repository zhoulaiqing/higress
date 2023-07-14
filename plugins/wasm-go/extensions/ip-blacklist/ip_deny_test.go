package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ip-tools"
	"testing"
)

func TestIpDeny(t *testing.T) {
	var ipIntervalRootV4 *ip_tools.IPIntervalNode
	var ipIntervalRootV6 *ip_tools.IPIntervalNode

	_, minIp, maxIp, _ := ip_tools.GetIPIntRange("1.1.1.1")
	ipIntervalRootV4 = ip_tools.IPIntervalTreeInsert(ipIntervalRootV4, ip_tools.IPInterval{minIp, maxIp})

	_, minIp2, maxIp2, _ := ip_tools.GetIPIntRange("2.2.2.2/16")
	ipIntervalRootV4 = ip_tools.IPIntervalTreeInsert(ipIntervalRootV4, ip_tools.IPInterval{minIp2, maxIp2})

	_, minIp3, maxIp3, _ := ip_tools.GetIPIntRange("2001:0db8:85a3:0000:0000:8a2e:0000:0000")
	ipIntervalRootV6 = ip_tools.IPIntervalTreeInsert(ipIntervalRootV6, ip_tools.IPInterval{minIp3, maxIp3})

	_, minIp4, maxIp4, _ := ip_tools.GetIPIntRange("2023:0db8::/32")
	ipIntervalRootV6 = ip_tools.IPIntervalTreeInsert(ipIntervalRootV6, ip_tools.IPInterval{minIp4, maxIp4})

	// true
	res1 := ip_tools.IPSearch(ipIntervalRootV4, minIp) != nil
	fmt.Println("Contains 1.1.1.1", res1)
	assert.True(t, res1)
	// false
	_, otherIp, _, _ := ip_tools.GetIPIntRange("10.23.2.2")
	res2 := ip_tools.IPSearch(ipIntervalRootV4, otherIp) != nil
	fmt.Println("Contains 10.23.2.2", res2)
	assert.False(t, res2)
	// true
	_, ip3, _, _ := ip_tools.GetIPIntRange("2.2.1.1")
	res3 := ip_tools.IPSearch(ipIntervalRootV4, ip3) != nil
	fmt.Println("Contains 2.2.1.1", res3)
	assert.True(t, res3)
	// true
	_, ip4, _, _ := ip_tools.GetIPIntRange("2001:0db8:85a3:0000:0000:8a2e:0000:0000")
	res4 := ip_tools.IPSearch(ipIntervalRootV6, ip4) != nil
	fmt.Println("Contains 2001:0db8:85a3:0000:0000:8a2e:0000:0000", res4)
	assert.True(t, res4)
	// true
	_, ip5, _, _ := ip_tools.GetIPIntRange("2023:0db8:85a3:0000:0000:8a2e:0000:0000")
	res5 := ip_tools.IPSearch(ipIntervalRootV6, ip5) != nil
	fmt.Println("Contains 2023:0db8:85a3:0000:0000:8a2e:0000:0000", res5)
	assert.True(t, res5)
}
