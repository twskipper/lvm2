package lvm2

import (
	"fmt"
	"strconv"
)

func ParseLvmSizeToBytes(s string) (int, error) {
	var size int
	var err error
	if s[len(s)-1] == 'B' {
		size, err = strconv.Atoi(s[:len(s)-1])
	} else {
		return 0, fmt.Errorf("invalid size: %s", s)
	}
	if err != nil {
		return 0, err
	}
	return size, nil
}
