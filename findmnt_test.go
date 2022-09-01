package lvm2

import (
	"fmt"
	"testing"
)

func TestGetMNT(t *testing.T) {
	mnt, err := GetMnt()
	if err != nil {
		t.Error(err)
	}
	//fmt.Printf("%+v\n", mnt)
	c, err := mnt.GetInfoOfMountPoint("/")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", c)
}
