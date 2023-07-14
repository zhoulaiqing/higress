package ip_tools

import (
	"fmt"
	"net"
	"strings"
)

func GetIPIntRange(ipStr string) (bool, *IPInt, *IPInt, error) {
	// 一个IP地址
	if !strings.Contains(ipStr, "/") {
		ip := net.ParseIP(ipStr)
		if ip == nil {
			return false, nil, nil, fmt.Errorf("Invalid IP address")
		}

		var ipInt *IPInt
		var isV4 bool
		if ip.To4() != nil {
			isV4 = true
			ipInt = IPIntFromIP(ip.To4())
		} else {
			isV4 = false
			ipInt = IPIntFromIP(ip.To16())
		}

		return isV4, ipInt, ipInt, nil
	}

	ip, ipNet, err := net.ParseCIDR(ipStr)
	if err != nil {
		return false, nil, nil, err
	}

	var isV4 bool
	var ipStartInt, ipEndInt *IPInt
	if ip.To4() != nil {
		isV4 = true
		ipStartInt = IPIntFromIP(ip.To4())
	} else {
		isV4 = false
		ipStartInt = IPIntFromIP(ip.To16())
	}

	ipMask := IPIntFromIPMask(ipNet.Mask, false)
	ipMaskInverse := IPIntFromIPMask(ipNet.Mask, true)

	ipStartInt, _ = ipStartInt.BitAnd(ipMask)
	ipEndInt, _ = ipStartInt.BitOr(ipMaskInverse)

	return isV4, ipStartInt, ipEndInt, nil
}

func IPIntFromIP(ip net.IP) *IPInt {
	result := &IPInt{integers: nil}
	if ip.To4() != nil {
		intVal := (uint32(ip[0]) << 24) | (uint32(ip[1]) << 16) | (uint32(ip[2]) << 8) | uint32(ip[3])
		result.integers = []uint32{intVal}
	} else {

		var intVar1, intVar2, intVar3, intVar4 uint32
		for i := 0; i < 4; i++ {
			intVar1 |= uint32(ip[i]) << (4 - i - 1) * 8
			intVar2 |= uint32(ip[i+4]) << (4 - i - 1) * 8
			intVar3 |= uint32(ip[i+8]) << (4 - i - 1) * 8
			intVar4 |= uint32(ip[i+12]) << (4 - i - 1) * 8
		}

		result.integers = []uint32{intVar1, intVar2, intVar3, intVar4}
	}

	return result
}

func IPIntFromIPMask(mask net.IPMask, inverse bool) *IPInt {
	result := &IPInt{integers: nil}
	if len(mask) == 4 {
		var intVal uint32
		if inverse {
			intVal = (uint32(^mask[0]) << 24) | (uint32(^mask[1]) << 16) | (uint32(^mask[2]) << 8) | uint32(^mask[3])
		} else {
			intVal = (uint32(mask[0]) << 24) | (uint32(mask[1]) << 16) | (uint32(mask[2]) << 8) | uint32(mask[3])
		}
		result.integers = []uint32{intVal}
	} else if len(mask) == 16 {
		var intVar1, intVar2, intVar3, intVar4 uint32
		for i := 0; i < 4; i++ {
			maskI := mask[i]
			if inverse {
				maskI = ^mask[i]
			}
			maskIPlus4 := mask[i+4]
			if inverse {
				maskIPlus4 = ^mask[i+4]
			}
			maskIPlus8 := mask[i+8]
			if inverse {
				maskIPlus8 = ^mask[i+8]
			}
			maskIPlus12 := mask[i+12]
			if inverse {
				maskIPlus12 = ^mask[i+12]
			}
			intVar1 |= uint32(maskI) << (4 - i - 1) * 8
			intVar2 |= uint32(maskIPlus4) << (4 - i - 1) * 8
			intVar3 |= uint32(maskIPlus8) << (4 - i - 1) * 8
			intVar4 |= uint32(maskIPlus12) << (4 - i - 1) * 8
		}

		result.integers = []uint32{intVar1, intVar2, intVar3, intVar4}
	}

	return result
}
