package lvm2

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	b, err := GetBlockdevices()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", b)
	fb := b.GetFreeBlockDevices()
	fmt.Printf("%+v\n", fb)
}
