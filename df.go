package lvm2

import (
	"bufio"
	"errors"
	"os/exec"
	"strings"
)

// df -l -B k --output=source,size,used,pcent,target
// Filesystem                   1K-blocks        Used Use% Mounted on
// /dev/mapper/VG00-root        50866176K    6724740K  14% /
// devtmpfs                     16377416K          0K   0% /dev
// tmpfs                        16389472K          0K   0% /dev/shm
// tmpfs                        16389472K    1622992K  10% /run
// tmpfs                        16389472K          0K   0% /sys/fs/cgroup
// /dev/sda1                     1020580K     148440K  15% /boot
// /dev/mapper/VG02-mysqllogs 1467960324K  359836172K  25% /mysqllogs
// /dev/mapper/VG01-mysqldata 6157179908K 4084786224K  67% /mysqldata
// tmpfs                         3277896K          0K   0% /run/user/607674
func DecodeLineTextFromDF(output []byte) *DFReport {
	var dfs []DFInfo
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Filesystem") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 5 {
			continue
		}

		dfs = append(dfs, DFInfo{
			Filesystem:  fields[0],
			Blocks1Kb:   fields[1],
			UsedInKb:    fields[2],
			UsedPercent: fields[3],
			MountedOn:   fields[4],
		})
	}

	return &DFReport{Info: dfs}
}

// df -l -B k --output=source,size,used,pcent,target

var dfCommand = []string{"df", "-l", "-B", "k", "--output=source,size,used,pcent,target"}

type DFReport struct {
	Info []DFInfo
}

func (d DFReport) GetInfoByMountPoint(s string) (*DFInfo, error) {
	for _, v := range d.Info {
		if v.MountedOn == s {
			return &v, nil
		}
	}
	return nil, errors.New("not found this mount point")
}

type DFInfo struct {
	Filesystem  string
	Blocks1Kb   string
	UsedInKb    string
	UsedPercent string
	MountedOn   string
}

func GetDFReport() (*DFReport, error) {
	cmd := exec.Command(dfCommand[0], dfCommand[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return DecodeLineTextFromDF(out), nil
}

func GetDFReportOverSSH(ssh *Client) (*DFReport, error) {
	c := "sudo " + strings.Join(dfCommand, " ")
	out, err := ssh.Cmd(c).Output()
	if err != nil {
		return nil, err
	}
	res := DecodeLineTextFromDF(out)
	if err != nil {
		return nil, err
	}
	return res, nil
}
