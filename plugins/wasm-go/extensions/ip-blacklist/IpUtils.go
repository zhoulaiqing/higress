package main

import (
	"fmt"
	"math/big"
	"net"
	"strings"
)

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

func getIPRange(ipString string) (*big.Int, *big.Int, error) {
	// 一个IP地址
	if !strings.Contains(ipString, "/") {
		ip := net.ParseIP(ipString)
		if ip == nil {
			return nil, nil, fmt.Errorf("Invalid IP address")
		}

		ipInt := new(big.Int).SetUint64(0)
		if ip.To4() != nil {
			ipInt.SetBytes(ip.To4())
		} else {
			ipInt.SetBytes(ip.To16())
		}

		return ipInt, ipInt, nil
	}

	// 一个IP段
	ip, ipNet, err := net.ParseCIDR(ipString)
	if err != nil {
		return nil, nil, err
	}

	start := big.NewInt(0)
	end := big.NewInt(0)

	if ip.To4() != nil {
		// IPv4
		ipStart := big.NewInt(0)
		ipStart.SetBytes(ip.To4())

		ipMask := big.NewInt(0)
		ipMask.SetBytes(ipNet.Mask)
		ipMaskNot := big.NewInt(0).SetUint64(big.NewInt(0).Not(ipMask).Uint64())

		ipStart.And(ipStart, ipMask)

		ipEnd := big.NewInt(0).Or(ipStart, ipMaskNot)

		start.Set(ipStart)
		end.Set(ipEnd)
	} else {
		// IPv6
		ipStart := big.NewInt(0)
		ipStart.SetBytes(ip.To16())

		ipMask := big.NewInt(0)
		ipMask.SetBytes(ipNet.Mask)
		ipMaskNot := big.NewInt(0).Not(ipMask)

		ipStart.And(ipStart, ipMask)
		ipEnd := big.NewInt(0).Or(ipStart, ipMaskNot)

		start.Set(ipStart)
		end.Set(ipEnd)
	}

	return start, end, nil
}

func getIPIntRange(ipStr string) (*IPInt, *IPInt, error) {
	// 一个IP地址
	if !strings.Contains(ipStr, "/") {
		ip := net.ParseIP(ipStr)
		if ip == nil {
			return nil, nil, fmt.Errorf("Invalid IP address")
		}

		var ipInt *IPInt
		if ip.To4() != nil {
			ipInt = IPIntFromIP(ip.To4())
		} else {
			ipInt = IPIntFromIP(ip.To16())
		}

		return ipInt, ipInt, nil
	}

	ip, ipNet, err := net.ParseCIDR(ipStr)
	if err != nil {
		return nil, nil, err
	}

	var ipStartInt, ipEndInt *IPInt
	if ip.To4() != nil {
		ipStartInt = IPIntFromIP(ip.To4())
	} else {
		ipStartInt = IPIntFromIP(ip.To16())
	}

	ipMask := IPIntFromIPMask(ipNet.Mask)
	// todo: 这里得到的值太大了
	ipMaskInverse, _ := ipMask.BitInverse()

	ipStartInt, _ = ipStartInt.BitAnd(ipMask)
	ipEndInt, _ = ipStartInt.BitOr(ipMaskInverse)

	return ipStartInt, ipEndInt, nil
}

func IPIntFromIP(ip net.IP) *IPInt {
	result := &IPInt{integers: nil}
	if ip.To4() != nil {
		//intVal := (uint64(ip[0]) << 24) | (uint64(ip[1]) << 16) | (uint64(ip[2]) << 8) | uint64(ip[3])
		intVal2 := ipToUint64(ip)
		result.integers = []uint64{intVal2}
	} else {

		var intVar1, intVar2 uint64
		for i := 0; i < 8; i++ {
			intVar1 |= uint64(ip[i]) << (8 - i - 1) * 8
			intVar2 |= uint64(ip[i+8]) << (8 - i - 1) * 8
		}

		result.integers = []uint64{intVar1, intVar2}
	}

	return result
}

func IPIntFromIPMask(mask net.IPMask) *IPInt {
	result := &IPInt{integers: nil}
	if len(mask) == 4 {
		intVal := (uint64(mask[0]) << 24) | (uint64(mask[1]) << 16) | (uint64(mask[2]) << 8) | uint64(mask[3])
		result.integers = []uint64{intVal}
	} else if len(mask) == 16 {
		var intVar1, intVar2 uint64
		for i := 0; i < 8; i++ {
			intVar1 |= uint64(mask[i]) << (8 - i - 1) * 8
			intVar2 |= uint64(mask[i+8]) << (8 - i - 1) * 8
		}

		result.integers = []uint64{intVar1, intVar2}
	}

	return result
}

func intToIP(ipInt uint32) net.IP {
	ip := make(net.IP, 4)
	ip[0] = byte(ipInt >> 24)
	ip[1] = byte(ipInt >> 16)
	ip[2] = byte(ipInt >> 8)
	ip[3] = byte(ipInt)
	return ip
}
