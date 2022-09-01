package lvm2

import (
	"fmt"
	"testing"
)

func TestGetPv(t *testing.T) {
	pv, err := GetPvReportAll()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", pv)
	fmt.Println(pv.IsPv("/dev/sda5"))
	fmt.Println(pv.GetTotalFreeSizeBytes())
}
