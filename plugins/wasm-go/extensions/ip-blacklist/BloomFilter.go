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
	data, _ := ipv4ToInt(ip)

	bf.bits[0] = 1 << (data % 64)
}

func (bf *BloomFilter) AddIpv6(ip string) {

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

func IPRange(ipStr string) (uint32, uint32, error) {
	var ip net.IP
	var mask net.IPMask

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
		mask = net.IPv4Mask(255, 255, 255, 255)
	}

	ipInt := ipToUint32(ip)
	maskInt := ipMaskToUint32(mask)

	networkInt := ipInt & maskInt
	minIPInt := networkInt
	maxIPInt := networkInt | ^maskInt

	return minIPInt, maxIPInt, nil
}

func ipToUint32(ip net.IP) uint32 {
	return (uint32(ip[0]) << 24) | (uint32(ip[1]) << 16) | (uint32(ip[2]) << 8) | uint32(ip[3])
}

func ipMaskToUint32(mask net.IPMask) uint32 {
	if len(mask) == 4 {
		return (uint32(mask[0]) << 24) | (uint32(mask[1]) << 16) | (uint32(mask[2]) << 8) | uint32(mask[3])
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
