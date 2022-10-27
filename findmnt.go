package lvm2

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

// findmnt  -l -n -o TARGET,SOURCE,FSTYPE,OPTIONS
var FindmntCommand = []string{
	"findmnt",
	"-n",
	"-r",
	"-o",
	"TARGET,SOURCE,FSTYPE,OPTIONS",
}

type Mnt struct {
	Filesystems []Filesystems `json:"filesystems"`
}

type Filesystems struct {
	Target  string `json:"target"`
	Source  string `json:"source"`
	Fstype  string `json:"fstype"`
	Options string `json:"options"`
	//Children []Children `json:"children"`
}

type Children struct {
	Target  string `json:"target"`
	Source  string `json:"source"`
	Fstype  string `json:"fstype"`
	Options string `json:"options"`
}

func (m Mnt) GetInfoOfMountPoint(mp string) (*Children, error) {
	var c Children
	for _, fs := range m.Filesystems {
		if fs.Target == mp {
			c.Fstype = fs.Fstype
			c.Options = fs.Options
			c.Source = fs.Source
			c.Target = fs.Target
			return &c, nil
		}
	}
	return nil, fmt.Errorf("%s mount point not found", mp)
}

func GetMnt() (*Mnt, error) {
	cmd := exec.Command(FindmntCommand[0], FindmntCommand[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	mnt := DecodeLineTextFromFindmnt(out)
	if err != nil {
		return nil, err
	}
	return &mnt, nil

}

// $ findmnt  -l -o TARGET,SOURCE,FSTYPE,OPTIONS
// header: TARGET                     SOURCE     FSTYPE     OPTIONS
// /sys                       sysfs      sysfs      rw,nosuid,nodev,noexec,relatime
// /proc                      proc       proc       rw,nosuid,nodev,noexec,relatime
// /dev                       devtmpfs   devtmpfs   rw,nosuid,size=3992644k,nr_inodes=998161,mode=755
// /sys/kernel/security       securityfs securityfs rw,nosuid,nodev,noexec,relatime
func DecodeLineTextFromFindmnt(output []byte) Mnt {
	var mnt Mnt
	var fs []Filesystems
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		s := strings.Split(scanner.Text(), " ")
		var f Filesystems
		f.Target = s[0]
		f.Source = s[1]
		f.Fstype = s[2]
		f.Options = s[3]
		fs = append(fs, f)
	}
	mnt.Filesystems = fs
	return mnt
}

func GetMntOverSSH(ssh *Client) (*Mnt, error) {
	c := "sudo " + strings.Join(FindmntCommand, " ")
	out, err := ssh.Cmd(c).Output()
	if err != nil {
		return nil, err
	}
	mnt := DecodeLineTextFromFindmnt(out)
	if err != nil {
		return nil, err
	}
	return &mnt, nil
}
