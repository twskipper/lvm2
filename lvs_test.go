package lvm2

import (
	"fmt"
	"testing"
)

func TestGetLVS(t *testing.T) {
	r, err := GetLVReportAll()
	if err != nil {
		t.Error(err)
	}
	isl := r.IsLv("/dev/mapper/vgubuntu-root")
	fmt.Println(isl)
	lv := r.GetLv("/dev/mapper/vgubuntu-root")
	fmt.Printf("%+v\n", lv)
	fmt.Println(lv.LvLayout, lv.VgName)
}
