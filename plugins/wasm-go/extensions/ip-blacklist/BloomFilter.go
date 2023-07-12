package main

import (
	"fmt"
	"hash/fnv"
	"net"
	"strconv"
	"strings"
)

type BloomFilter struct {
	bits [2]uint64
}

func NewBloomFilter() *BloomFilter {
	return &BloomFilter{}
}

func (bf *BloomFilter) Add(data string) {
	hash1, hash2 := hashData(data)
	bf.bits[hash1%64] |= (1 << (hash1 % 64))
	bf.bits[1+hash2%64] |= (1 << (hash2 % 64))
}

func (bf *BloomFilter) AddIpv4(ip string) {
}

func (bf *BloomFilter) AddIpv6(ip string) {

}

func setBitsInRange(c, a, b int) int {
	// 首先创建一个掩码，将目标范围内的位设置为1，其他位保持为0
	mask := (1 << (b - a + 1)) - 1<<a

	// 将掩码与c进行按位或运算，将目标范围内的位设置为1
	result := c | mask
	return result
}

func (bf *BloomFilter) Contains(data string) bool {
	hash1, hash2 := hashData(data)
	if bf.bits[hash1%64]&(1<<(hash1%64)) == 0 {
		return false
	}
	if bf.bits[1+hash2%64]&(1<<(hash2%64)) == 0 {
		return false
	}
	return true
}

func hashData(data string) (uint64, uint64) {
	h := fnv.New64a()
	h.Write([]byte(data))
	hash := h.Sum64()
	return uint64(hash >> 32), uint64(hash)
}

func ipv4ToInt(ip string) (uint64, error) {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return 0, fmt.Errorf("Invalid IP address format")
	}

	var ipInt uint64
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return 0, fmt.Errorf("Invalid IP address format")
		}
		ipInt = (ipInt << 8) + uint64(num)
	}

	return ipInt, nil
}

func IPRange(ipStr string) (uint64, uint64, error) {
	var ip net.IP
	var mask net.IPMask
	isFixedIp := false

	if strings.Contains(ipStr, "/") {
		// 处理带掩码的IP地址
		pureIp, ipNet, err := net.ParseCIDR(ipStr)
		if err != nil {
			return 0, 0, fmt.Errorf("Invalid IP or mask address")
		}
		ip = pureIp.To4()
		mask = ipNet.Mask
	} else {
		// 处理固定IP地址
		ip = net.ParseIP(ipStr)
		if ip == nil {
			return 0, 0, fmt.Errorf("Invalid IP address")
		}
		ip = ip.To4()
		isFixedIp = true
	}

	ipInt := ipToUint64(ip)
	if isFixedIp {
		return ipInt, ipInt, nil
	}

	maskInt := ipMaskToUint64(mask)

	networkInt := ipInt & maskInt
	minIPInt := networkInt
	maxIPInt := networkInt | ^maskInt

	return minIPInt, maxIPInt, nil
}

func ipToUint64(ip net.IP) uint64 {
	return (uint64(ip[0]) << 24) | (uint64(ip[1]) << 16) | (uint64(ip[2]) << 8) | uint64(ip[3])
}

func ipMaskToUint64(mask net.IPMask) uint64 {
	if len(mask) == 4 {
		return (uint64(mask[0]) << 24) | (uint64(mask[1]) << 16) | (uint64(mask[2]) << 8) | uint64(mask[3])
	}
	return 0
}

func intToIP(ipInt uint32) net.IP {
	ip := make(net.IP, 4)
	ip[0] = byte(ipInt >> 24)
	ip[1] = byte(ipInt >> 16)
	ip[2] = byte(ipInt >> 8)
	ip[3] = byte(ipInt)
	return ip
}
