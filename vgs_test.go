package lvm2

import (
	"fmt"
	"testing"
)

func TestGetVGS(t *testing.T) {
	vgs, err := GetVGReportAll()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(vgs)
	fmt.Println(vgs.IsVG("vgubuntu"))
	vg := vgs.GetVG("vgubuntu")
	fmt.Printf("%+v\n", vg)

}
