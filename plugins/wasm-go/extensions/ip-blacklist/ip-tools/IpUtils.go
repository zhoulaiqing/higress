package ip_tools

import (
	"fmt"
	"net"
	"strings"
)

func GetIPIntRange(ipStr string) (*IPInt, *IPInt, error) {
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

	ipMask := IPIntFromIPMask(ipNet.Mask, false)
	ipMaskInverse := IPIntFromIPMask(ipNet.Mask, true)

	ipStartInt, _ = ipStartInt.BitAnd(ipMask)
	ipEndInt, _ = ipStartInt.BitOr(ipMaskInverse)

	return ipStartInt, ipEndInt, nil
}

func IPIntFromIP(ip net.IP) *IPInt {
	result := &IPInt{integers: nil}
	if ip.To4() != nil {
		intVal := (uint64(ip[0]) << 24) | (uint64(ip[1]) << 16) | (uint64(ip[2]) << 8) | uint64(ip[3])
		result.integers = []uint64{intVal}
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

func IPIntFromIPMask(mask net.IPMask, inverse bool) *IPInt {
	result := &IPInt{integers: nil}
	if len(mask) == 4 {
		var intVal uint64
		if inverse {
			intVal = (uint64(^mask[0]) << 24) | (uint64(^mask[1]) << 16) | (uint64(^mask[2]) << 8) | uint64(^mask[3])
		} else {
			intVal = (uint64(mask[0]) << 24) | (uint64(mask[1]) << 16) | (uint64(mask[2]) << 8) | uint64(mask[3])
		}
		result.integers = []uint64{intVal}
	} else if len(mask) == 16 {
		var intVar1, intVar2 uint64
		for i := 0; i < 8; i++ {
			maskI := mask[i]
			if inverse {
				maskI = ^mask[i]
			}

			maskIPlus8 := mask[i+8]
			if inverse {
				maskIPlus8 = ^mask[i+8]
			}
			intVar1 |= uint64(maskI) << (8 - i - 1) * 8
			intVar2 |= uint64(maskIPlus8) << (8 - i - 1) * 8
		}

		result.integers = []uint64{intVar1, intVar2}
	}

	return result
}
