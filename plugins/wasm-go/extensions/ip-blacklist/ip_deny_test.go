package main

import (
	"fmt"
)

func test_ip_deny() {
	var ipIntervalRoot *IPIntervalNode

	minIp, maxIp, _ := getIPIntRange("1.1.1.1")
	ipIntervalRoot = IPIntervalTreeInsert(ipIntervalRoot, IPInterval{minIp, maxIp})

	minIp2, maxIp2, _ := getIPIntRange("2.2.2.2/16")
	ipIntervalRoot = IPIntervalTreeInsert(ipIntervalRoot, IPInterval{minIp2, maxIp2})

	minIp3, maxIp3, _ := getIPIntRange("2001:0db8:85a3:0000:0000:8a2e:0000:0000")
	ipIntervalRoot = IPIntervalTreeInsert(ipIntervalRoot, IPInterval{minIp3, maxIp3})

	minIp4, maxIp4, _ := getIPIntRange("2023:0db8::/32")
	ipIntervalRoot = IPIntervalTreeInsert(ipIntervalRoot, IPInterval{minIp4, maxIp4})

	// true
	res1 := IPSearch(ipIntervalRoot, minIp) != nil
	fmt.Println("Contains 1.1.1.1", res1)

	// false
	otherIp, _, _ := getIPIntRange("10.23.2.2")
	res2 := IPSearch(ipIntervalRoot, otherIp) != nil
	fmt.Println("Contains 10.23.2.2", res2)

	// true
	ip3, _, _ := getIPIntRange("2.2.1.1")
	res3 := IPSearch(ipIntervalRoot, ip3) != nil
	fmt.Println("Contains 2.2.1.1", res3)

	// true
	ip4, _, _ := getIPIntRange("2001:0db8:85a3:0000:0000:8a2e:0000:0000")
	res4 := IPSearch(ipIntervalRoot, ip4) != nil
	fmt.Println("Contains 2001:0db8:85a3:0000:0000:8a2e:0000:0000", res4)

	// true
	ip5, _, _ := getIPIntRange("2023:0db8:85a3:0000:0000:8a2e:0000:0000")
	res5 := IPSearch(ipIntervalRoot, ip5) != nil
	fmt.Println("Contains 2023:0db8:85a3:0000:0000:8a2e:0000:0000", res5)
}
