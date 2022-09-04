package lvm2

import (
	"encoding/json"
	"os/exec"
	"strings"
)

var (
	LvsAllCmd = []string{"lvs", "-o", "all", "--reportformat", "json", "--units", "b"}
)

type LVReportAll struct {
	Report []LVReport `json:"report"`
}

func (l LVReportAll) GetAllLv() []Lv {
	var lv []Lv
	for _, r := range l.Report {
		lv = append(lv, r.Lv...)
	}
	return lv
}
func (l LVReportAll) IsLv(name string) bool {
	lvs := l.GetAllLv()
	for _, lv := range lvs {
		if lv.LvDmPath == name || lv.LvPath == name || lv.LvName == name {
			return true
		}
	}
	return false
}
func (l LVReportAll) GetLv(name string) *Lv {
	lvs := l.GetAllLv()
	for _, lv := range lvs {
		if lv.LvDmPath == name || lv.LvPath == name || lv.LvName == name {
			return &lv
		}
	}
	return nil
}

type Lv struct {
	LvUUID               string `json:"lv_uuid"`
	LvName               string `json:"lv_name"`
	LvFullName           string `json:"lv_full_name"`
	LvPath               string `json:"lv_path"`
	LvDmPath             string `json:"lv_dm_path"`
	LvParent             string `json:"lv_parent"`
	LvLayout             string `json:"lv_layout"`
	LvRole               string `json:"lv_role"`
	LvInitialImageSync   string `json:"lv_initial_image_sync"`
	LvImageSynced        string `json:"lv_image_synced"`
	LvMerging            string `json:"lv_merging"`
	LvConverting         string `json:"lv_converting"`
	LvAllocationPolicy   string `json:"lv_allocation_policy"`
	LvAllocationLocked   string `json:"lv_allocation_locked"`
	LvFixedMinor         string `json:"lv_fixed_minor"`
	LvSkipActivation     string `json:"lv_skip_activation"`
	LvWhenFull           string `json:"lv_when_full"`
	LvActive             string `json:"lv_active"`
	LvActiveLocally      string `json:"lv_active_locally"`
	LvActiveRemotely     string `json:"lv_active_remotely"`
	LvActiveExclusively  string `json:"lv_active_exclusively"`
	LvMajor              string `json:"lv_major"`
	LvMinor              string `json:"lv_minor"`
	LvReadAhead          string `json:"lv_read_ahead"`
	LvSize               string `json:"lv_size"`
	LvMetadataSize       string `json:"lv_metadata_size"`
	SegCount             string `json:"seg_count"`
	Origin               string `json:"origin"`
	OriginUUID           string `json:"origin_uuid"`
	OriginSize           string `json:"origin_size"`
	LvAncestors          string `json:"lv_ancestors"`
	LvFullAncestors      string `json:"lv_full_ancestors"`
	LvDescendants        string `json:"lv_descendants"`
	LvFullDescendants    string `json:"lv_full_descendants"`
	RaidMismatchCount    string `json:"raid_mismatch_count"`
	RaidSyncAction       string `json:"raid_sync_action"`
	RaidWriteBehind      string `json:"raid_write_behind"`
	RaidMinRecoveryRate  string `json:"raid_min_recovery_rate"`
	RaidMaxRecoveryRate  string `json:"raid_max_recovery_rate"`
	MovePv               string `json:"move_pv"`
	MovePvUUID           string `json:"move_pv_uuid"`
	ConvertLv            string `json:"convert_lv"`
	ConvertLvUUID        string `json:"convert_lv_uuid"`
	MirrorLog            string `json:"mirror_log"`
	MirrorLogUUID        string `json:"mirror_log_uuid"`
	DataLv               string `json:"data_lv"`
	DataLvUUID           string `json:"data_lv_uuid"`
	MetadataLv           string `json:"metadata_lv"`
	MetadataLvUUID       string `json:"metadata_lv_uuid"`
	PoolLv               string `json:"pool_lv"`
	PoolLvUUID           string `json:"pool_lv_uuid"`
	LvTags               string `json:"lv_tags"`
	LvProfile            string `json:"lv_profile"`
	LvLockargs           string `json:"lv_lockargs"`
	LvTime               string `json:"lv_time"`
	LvTimeRemoved        string `json:"lv_time_removed"`
	LvHost               string `json:"lv_host"`
	LvModules            string `json:"lv_modules"`
	LvHistorical         string `json:"lv_historical"`
	LvKernelMajor        string `json:"lv_kernel_major"`
	LvKernelMinor        string `json:"lv_kernel_minor"`
	LvKernelReadAhead    string `json:"lv_kernel_read_ahead"`
	LvPermissions        string `json:"lv_permissions"`
	LvSuspended          string `json:"lv_suspended"`
	LvLiveTable          string `json:"lv_live_table"`
	LvInactiveTable      string `json:"lv_inactive_table"`
	LvDeviceOpen         string `json:"lv_device_open"`
	DataPercent          string `json:"data_percent"`
	SnapPercent          string `json:"snap_percent"`
	MetadataPercent      string `json:"metadata_percent"`
	CopyPercent          string `json:"copy_percent"`
	SyncPercent          string `json:"sync_percent"`
	CacheTotalBlocks     string `json:"cache_total_blocks"`
	CacheUsedBlocks      string `json:"cache_used_blocks"`
	CacheDirtyBlocks     string `json:"cache_dirty_blocks"`
	CacheReadHits        string `json:"cache_read_hits"`
	CacheReadMisses      string `json:"cache_read_misses"`
	CacheWriteHits       string `json:"cache_write_hits"`
	CacheWriteMisses     string `json:"cache_write_misses"`
	KernelCacheSettings  string `json:"kernel_cache_settings"`
	KernelCachePolicy    string `json:"kernel_cache_policy"`
	KernelMetadataFormat string `json:"kernel_metadata_format"`
	LvHealthStatus       string `json:"lv_health_status"`
	KernelDiscards       string `json:"kernel_discards"`
	LvCheckNeeded        string `json:"lv_check_needed"`
	LvMergeFailed        string `json:"lv_merge_failed"`
	LvSnapshotInvalid    string `json:"lv_snapshot_invalid"`
	VdoOperatingMode     string `json:"vdo_operating_mode"`
	VdoCompressionState  string `json:"vdo_compression_state"`
	VdoIndexState        string `json:"vdo_index_state"`
	VdoUsedSize          string `json:"vdo_used_size"`
	VdoSavingPercent     string `json:"vdo_saving_percent"`
	LvAttr               string `json:"lv_attr"`
	VgFmt                string `json:"vg_fmt"`
	VgUUID               string `json:"vg_uuid"`
	VgName               string `json:"vg_name"`
	VgAttr               string `json:"vg_attr"`
	VgPermissions        string `json:"vg_permissions"`
	VgExtendable         string `json:"vg_extendable"`
	VgExported           string `json:"vg_exported"`
	VgPartial            string `json:"vg_partial"`
	VgAllocationPolicy   string `json:"vg_allocation_policy"`
	VgClustered          string `json:"vg_clustered"`
	VgShared             string `json:"vg_shared"`
	VgSize               string `json:"vg_size"`
	VgFree               string `json:"vg_free"`
	VgSysid              string `json:"vg_sysid"`
	VgSystemid           string `json:"vg_systemid"`
	VgLockType           string `json:"vg_lock_type"`
	VgLockArgs           string `json:"vg_lock_args"`
	VgExtentSize         string `json:"vg_extent_size"`
	VgExtentCount        string `json:"vg_extent_count"`
	VgFreeCount          string `json:"vg_free_count"`
	MaxLv                string `json:"max_lv"`
	MaxPv                string `json:"max_pv"`
	PvCount              string `json:"pv_count"`
	VgMissingPvCount     string `json:"vg_missing_pv_count"`
	LvCount              string `json:"lv_count"`
	SnapCount            string `json:"snap_count"`
	VgSeqno              string `json:"vg_seqno"`
	VgTags               string `json:"vg_tags"`
	VgProfile            string `json:"vg_profile"`
	VgMdaCount           string `json:"vg_mda_count"`
	VgMdaUsedCount       string `json:"vg_mda_used_count"`
	VgMdaFree            string `json:"vg_mda_free"`
	VgMdaSize            string `json:"vg_mda_size"`
	VgMdaCopies          string `json:"vg_mda_copies"`
}
type LVReport struct {
	Lv []Lv `json:"lv"`
}

func (r *LVReport) GetLv(lvName string) *Lv {
	for _, lv := range r.Lv {
		if lv.LvName == lvName {
			return &lv
		}
	}
	return nil
}

func (r *LVReport) GetLvUUID(lvUUID string) *Lv {
	for _, lv := range r.Lv {
		if lv.LvUUID == lvUUID {
			return &lv
		}
	}
	return nil
}

func GetLVReportAll() (*LVReportAll, error) {
	cmd := exec.Command("sudo", LvsAllCmd...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var r LVReportAll
	err = json.Unmarshal(out, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func GetLVReportAllOverSSH(ssh *Client) (*LVReportAll, error) {
	c := "sudo " + strings.Join(LvsAllCmd, " ")
	out, err := ssh.Cmd(c).Output()
	if err != nil {
		return nil, err
	}
	var r LVReportAll
	err = json.Unmarshal(out, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
