package lvm2

import (
	"fmt"
	"testing"
)

func TestGetLVS(t *testing.T) {
	r, err := GetLvReport()
	if err != nil {
		t.Error(err)
	}

	//fmt.Printf("%+v\n", r)
	allLv := r.GetAllLv()
	for _, lv := range allLv {
		if lv.LvDmPath == "/dev/mapper/vgubuntu-root" {
			fmt.Printf("%+v\n", lv)
			fmt.Println(lv.VgFree, lv.VgName,lv.LvName,lv.VgUUID,lv.LvDmPath,lv.LvPath,lv.LvParent)
		}

	}
}
