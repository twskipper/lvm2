package lvm2

import (
	"encoding/json"
	"errors"
	"os/exec"
)

var FindmntCommand = []string{
	"findmnt",
	"--real",
	"--json",
}

type Mnt struct {
	Filesystems []Filesystems `json:"filesystems"`
}

type Filesystems struct {
	Target   string     `json:"target"`
	Source   string     `json:"source"`
	Fstype   string     `json:"fstype"`
	Options  string     `json:"options"`
	Children []Children `json:"children"`
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
		for _, child := range fs.Children {
			if child.Target == mp {
				c.Fstype = child.Fstype
				c.Options = child.Options
				c.Source = child.Source
				c.Target = child.Target
				return &c, nil
			}
		}
	}
	return nil, errors.New("mount point not found")
}

func GetMnt() (*Mnt, error) {
	cmd := exec.Command(FindmntCommand[0], FindmntCommand[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var mnt Mnt
	err = json.Unmarshal(out, &mnt)
	if err != nil {
		return nil, err
	}
	return &mnt, nil

}
