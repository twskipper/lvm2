package lvm2

import (
	"encoding/json"
	"os/exec"
)

var (
	PvsAllCmd = []string{"pvs", "-o", "all,vg_all", "--reportformat", "json", "--units", "b"}
)

type PvReportAll struct {
	Report []PvReport `json:"report"`
}

func (p PvReportAll) GetAllPv() []Pv {
	var pv []Pv
	for _, r := range p.Report {
		pv = append(pv, r.Pv...)
	}
	return pv
}
func (p PvReportAll) IsPv(name string) bool {
	pvs := p.GetAllPv()
	for _, pv := range pvs {
		if pv.PvName == name || pv.PvUUID == name {
			return true
		}
	}
	return false
}
func (p PvReportAll) GetPv(name string) *Pv {
	pvs := p.GetAllPv()
	for _, pv := range pvs {
		if pv.PvName == name || pv.PvUUID == name {
			return &pv
		}
	}
	return nil
}
func (p PvReportAll) GetPvByVg(vg string) []Pv {
	var pv []Pv
	for _, r := range p.Report {
		for _, p := range r.Pv {
			if p.VgName == vg {
				pv = append(pv, p)
			}
		}
	}
	return pv
}

func (p PvReportAll) GetTotalFreeSizeBytes() (int, error) {
	var total int

	for _, r := range p.Report {
		for _, p := range r.Pv {
			size, err := ParseLvmSizeToBytes(p.PvFree)
			if err != nil {
				return 0, err
			}
			total += size
		}
	}
	return total, nil
}

type Pv struct {
	PvFmt              string `json:"pv_fmt"`
	PvUUID             string `json:"pv_uuid"`
	DevSize            string `json:"dev_size"`
	PvName             string `json:"pv_name"`
	PvMajor            string `json:"pv_major"`
	PvMinor            string `json:"pv_minor"`
	PvMdaFree          string `json:"pv_mda_free"`
	PvMdaSize          string `json:"pv_mda_size"`
	PvExtVsn           string `json:"pv_ext_vsn"`
	PeStart            string `json:"pe_start"`
	PvSize             string `json:"pv_size"`
	PvFree             string `json:"pv_free"`
	PvUsed             string `json:"pv_used"`
	PvAttr             string `json:"pv_attr"`
	PvAllocatable      string `json:"pv_allocatable"`
	PvExported         string `json:"pv_exported"`
	PvMissing          string `json:"pv_missing"`
	PvPeCount          string `json:"pv_pe_count"`
	PvPeAllocCount     string `json:"pv_pe_alloc_count"`
	PvTags             string `json:"pv_tags"`
	PvMdaCount         string `json:"pv_mda_count"`
	PvMdaUsedCount     string `json:"pv_mda_used_count"`
	PvBaStart          string `json:"pv_ba_start"`
	PvBaSize           string `json:"pv_ba_size"`
	PvInUse            string `json:"pv_in_use"`
	PvDuplicate        string `json:"pv_duplicate"`
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
type PvReport struct {
	Pv []Pv `json:"pv"`
}

func GetPvReportAll() (*PvReportAll, error) {
	c := PvsAllCmd
	cmd := exec.Command("sudo", c...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var r PvReportAll
	err = json.Unmarshal(out, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
