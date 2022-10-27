package lvm2

type LvmReport struct {
	LVReport     *LVReportAll  `json:"lv_report,omitempty"`
	VGReport     *VGReportAll  `json:"vg_report,omitempty"`
	PVReport     *PVReportAll  `json:"pv_report,omitempty"`
	Blockdevices *Blockdevices `json:"blockdevices,omitempty"`
	Mnt          *Mnt          `json:"mnt,omitempty"`
	DF           *DFReport     `json:"df,omitempty"`
}

func GetLvmReportAll() (*LvmReport, error) {
	var lr LvmReport
	lvr, err := GetLVReportAll()
	if err != nil {
		return nil, err
	}
	vgr, err := GetVGReportAll()
	if err != nil {
		return nil, err
	}
	pvr, err := GetPVReportAll()
	if err != nil {
		return nil, err
	}
	blks, err := GetBlockdevices()
	if err != nil {
		return nil, err
	}
	mnt, err := GetMnt()
	if err != nil {
		return nil, err
	}
	df, err := GetDFReport()
	if err != nil {
		return nil, err
	}
	lr.LVReport = lvr
	lr.VGReport = vgr
	lr.PVReport = pvr
	lr.Blockdevices = blks
	lr.Mnt = mnt
	lr.DF = df
	return &lr, nil
}

func GetLvmReportAllOverSSH(sshClient *Client) (*LvmReport, error) {
	var lr LvmReport
	lvr, err := GetLVReportAllOverSSH(sshClient)
	if err != nil {
		return nil, err
	}
	vgr, err := GetVGReportAllOverSSH(sshClient)
	if err != nil {
		return nil, err
	}
	pvr, err := GetPVReportAllOverSSH(sshClient)
	if err != nil {
		return nil, err
	}

	mnt, err := GetMntOverSSH(sshClient)
	if err != nil {
		return nil, err
	}
	blks, err := GetBlockdevicesOverSSH(sshClient)
	if err != nil {
		return nil, err
	}
	lr.LVReport = lvr
	lr.VGReport = vgr
	lr.PVReport = pvr
	lr.Blockdevices = blks
	lr.Mnt = mnt
	return &lr, nil
}
