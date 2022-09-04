package lvm2

import (
	"encoding/json"
	"os/exec"
	"strings"
)

var (
	VgsAllCmd = []string{
		"vgs",
		"-o",
		"all",
		"--reportformat",
		"json",
		"--units",
		"b",
	}
)

type VGReportAll struct {
	Report []VGReport `json:"report"`
}

func (v VGReportAll) GetAllVG() []Vg {
	var vg []Vg
	for _, r := range v.Report {
		vg = append(vg, r.Vg...)
	}
	return vg
}
func (v VGReportAll) IsVG(name string) bool {
	vgs := v.GetAllVG()
	for _, vg := range vgs {
		if vg.VgName == name || vg.VgUUID == name {
			return true
		}
	}
	return false
}
func (v VGReportAll) GetVG(name string) *Vg {
	vgs := v.GetAllVG()
	for _, vg := range vgs {
		if vg.VgName == name || vg.VgUUID == name {
			return &vg
		}
	}
	return nil
}
func (v VGReportAll) GetTotalFreeSize() string {
	var total string
	for _, r := range v.Report {
		for _, vg := range r.Vg {
			total += vg.VgFree
		}
	}
	return total
}

type Vg struct {
	VgFmt              string `json:"vg_fmt"`
	VgUUID             string `json:"vg_uuid"`
	VgName             string `json:"vg_name"`
	VgAttr             string `json:"vg_attr"`
	VgPermissions      string `json:"vg_permissions"`
	VgExtendable       string `json:"vg_extendable"`
	VgExported         string `json:"vg_exported"`
	VgPartial          string `json:"vg_partial"`
	VgAllocationPolicy string `json:"vg_allocation_policy"`
	VgClustered        string `json:"vg_clustered"`
	VgShared           string `json:"vg_shared"`
	VgSize             string `json:"vg_size"`
	VgFree             string `json:"vg_free"`
	VgSysid            string `json:"vg_sysid"`
	VgSystemid         string `json:"vg_systemid"`
	VgLockType         string `json:"vg_lock_type"`
	VgLockArgs         string `json:"vg_lock_args"`
	VgExtentSize       string `json:"vg_extent_size"`
	VgExtentCount      string `json:"vg_extent_count"`
	VgFreeCount        string `json:"vg_free_count"`
	MaxLv              string `json:"max_lv"`
	MaxPv              string `json:"max_pv"`
	PvCount            string `json:"pv_count"`
	VgMissingPvCount   string `json:"vg_missing_pv_count"`
	LvCount            string `json:"lv_count"`
	SnapCount          string `json:"snap_count"`
	VgSeqno            string `json:"vg_seqno"`
	VgTags             string `json:"vg_tags"`
	VgProfile          string `json:"vg_profile"`
	VgMdaCount         string `json:"vg_mda_count"`
	VgMdaUsedCount     string `json:"vg_mda_used_count"`
	VgMdaFree          string `json:"vg_mda_free"`
	VgMdaSize          string `json:"vg_mda_size"`
	VgMdaCopies        string `json:"vg_mda_copies"`
}
type VGReport struct {
	Vg []Vg `json:"vg"`
}

func GetVGReportAll() (*VGReportAll, error) {
	c := VgsAllCmd
	cmd := exec.Command("sudo", c...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var r VGReportAll
	err = json.Unmarshal(out, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func GetVGReportAllOverSSH(ssh *Client) (*VGReportAll, error) {
	c := "sudo " + strings.Join(VgsAllCmd, " ")
	out, err := ssh.Cmd(c).Output()
	if err != nil {
		return nil, err
	}
	var r VGReportAll
	err = json.Unmarshal(out, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
