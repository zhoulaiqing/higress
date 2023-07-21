package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ip-tools"
	"testing"
)

func TestIpDeny(t *testing.T) {
	var ipIntervalTreeV4 *ip_tools.AVLTree = ip_tools.NewAVLTree()
	var ipIntervalTreeV6 *ip_tools.AVLTree = ip_tools.NewAVLTree()

	_, minIp, maxIp, _ := ip_tools.GetIPIntRange("1.1.1.1")
	ipIntervalTreeV4.Insert(&ip_tools.IPInterval{minIp, maxIp})

	_, minIp2, maxIp2, _ := ip_tools.GetIPIntRange("2.2.2.2/16")
	ipIntervalTreeV4.Insert(&ip_tools.IPInterval{minIp2, maxIp2})

	_, minIp3, maxIp3, _ := ip_tools.GetIPIntRange("2001:0db8:85a3:0000:0000:8a2e:0000:0000")
	ipIntervalTreeV6.Insert(&ip_tools.IPInterval{minIp3, maxIp3})

	_, minIp4, maxIp4, _ := ip_tools.GetIPIntRange("2023:0db8::/32")
	ipIntervalTreeV6.Insert(&ip_tools.IPInterval{minIp4, maxIp4})

	// true
	res1 := ipIntervalTreeV4.SearchSingle(minIp) != nil
	fmt.Println("Contains 1.1.1.1", res1)
	assert.True(t, res1)
	// false
	_, otherIp, _, _ := ip_tools.GetIPIntRange("10.23.2.2")
	res2 := ipIntervalTreeV4.SearchSingle(otherIp) != nil
	fmt.Println("Contains 10.23.2.2", res2)
	assert.False(t, res2)
	// true
	_, ip3, _, _ := ip_tools.GetIPIntRange("2.2.1.1")
	res3 := ipIntervalTreeV4.SearchSingle(ip3) != nil
	fmt.Println("Contains 2.2.1.1", res3)
	assert.True(t, res3)
	// true
	_, ip4, _, _ := ip_tools.GetIPIntRange("2001:0db8:85a3:0000:0000:8a2e:0000:0000")
	res4 := ipIntervalTreeV6.SearchSingle(ip4) != nil
	fmt.Println("Contains 2001:0db8:85a3:0000:0000:8a2e:0000:0000", res4)
	assert.True(t, res4)
	// true
	_, ip5, _, _ := ip_tools.GetIPIntRange("2023:0db8:85a3:0000:0000:8a2e:0000:0000")
	res5 := ipIntervalTreeV6.SearchSingle(ip5) != nil
	fmt.Println("Contains 2023:0db8:85a3:0000:0000:8a2e:0000:0000", res5)
	assert.True(t, res5)
}
