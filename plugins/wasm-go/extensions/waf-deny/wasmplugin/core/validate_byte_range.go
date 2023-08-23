package core

import (
	"fmt"
	"strconv"
	"strings"
)

type ValidateByteRange struct {
	validBytes [256]bool // array, not slice, so don't pass as-is to functions
}

func NewValidateByRange(byteValue string) (*ValidateByteRange, error) {
	var validBytes [256]bool
	for _, br := range strings.Split(byteValue, ",") {
		br = strings.TrimSpace(br)
		start, end, ok := strings.Cut(br, "-")

		if !ok {
			if b, err := strconv.Atoi(start); err != nil {
				return nil, err
			} else if err := validateByte(b); err != nil {
				return nil, err
			} else {
				validBytes[b] = true
			}
			continue
		}
		s, err := strconv.Atoi(start)
		if err != nil {
			return nil, err
		}
		if err := validateByte(s); err != nil {
			return nil, err
		}
		e, err := strconv.Atoi(end)
		if err != nil {
			return nil, err
		}
		if err := validateByte(e); err != nil {
			return nil, err
		}
		for i := s; i <= e; i++ {
			validBytes[i] = true
		}
	}

	return &ValidateByteRange{validBytes: validBytes}, nil
}

func (o *ValidateByteRange) Validate(value string) bool {
	if value == "" {
		return true
	}

	for i := 0; i < len(value); i++ {
		c := value[i]
		if !o.validBytes[c] {
			return false
		}
	}
	return true
}

func validateByte(b int) error {
	if b < 0 || b > 255 {
		return fmt.Errorf("invalid byte %d", b)
	}
	return nil
}
