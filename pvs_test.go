package lvm2

import (
	"fmt"
	"testing"
)

func TestGetPv(t *testing.T) {
	pv, err := GetPVReportAll()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", pv)
	fmt.Println(pv.IsPv("/dev/sda5"))
	fmt.Println(pv.GetAllFreePv())
}
