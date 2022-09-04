package lvm2

import (
	"encoding/json"
	"errors"
	"os/exec"
	"strings"
)

var (
	LSBLKCmd = []string{"lsblk", "-O", "--json", "-b"}
)

type Blockdevices struct {
	Blockdevices []Blockdevice `json:"blockdevices"`
}

type Blockdevice struct {
	Name       string        `json:"name"`
	Kname      string        `json:"kname"`
	Path       string        `json:"path"`
	MajMin     string        `json:"maj:min"`
	Fsavail    string        `json:"fsavail"`
	Fssize     string        `json:"fssize"`
	Fstype     string        `json:"fstype"`
	Fsused     string        `json:"fsused"`
	Fsuse      string        `json:"fsuse%"`
	Mountpoint string        `json:"mountpoint"`
	Label      interface{}   `json:"label"`
	UUID       string        `json:"uuid"`
	Ptuuid     string        `json:"ptuuid"`
	Pttype     string        `json:"pttype"`
	Parttype   interface{}   `json:"parttype"`
	Partlabel  interface{}   `json:"partlabel"`
	Partuuid   interface{}   `json:"partuuid"`
	Partflags  interface{}   `json:"partflags"`
	Ra         int           `json:"ra"`
	Ro         bool          `json:"ro"`
	Rm         bool          `json:"rm"`
	Hotplug    bool          `json:"hotplug"`
	Model      string        `json:"model"`
	Serial     string        `json:"serial"`
	Size       int           `json:"size"`
	State      string        `json:"state"`
	Owner      string        `json:"owner"`
	Group      string        `json:"group"`
	Mode       string        `json:"mode"`
	Alignment  int           `json:"alignment"`
	MinIo      int           `json:"min-io"`
	OptIo      int           `json:"opt-io"`
	PhySec     int           `json:"phy-sec"`
	LogSec     int           `json:"log-sec"`
	Rota       bool          `json:"rota"`
	Sched      string        `json:"sched"`
	RqSize     int           `json:"rq-size"`
	Type       string        `json:"type"`
	DiscAln    int           `json:"disc-aln"`
	DiscGran   int           `json:"disc-gran"`
	DiscMax    int           `json:"disc-max"`
	DiscZero   bool          `json:"disc-zero"`
	Wsame      int           `json:"wsame"`
	Wwn        interface{}   `json:"wwn"`
	Rand       bool          `json:"rand"`
	Pkname     interface{}   `json:"pkname"`
	Hctl       interface{}   `json:"hctl"`
	Tran       string        `json:"tran"`
	Subsystems string        `json:"subsystems"`
	Rev        string        `json:"rev"`
	Vendor     string        `json:"vendor"`
	Zoned      string        `json:"zoned"`
	Children   []Blockdevice `json:"children,omitempty"`
}

func SearchBlk(blk *Blockdevice, path string) *Blockdevice {
	if blk.Path == path {
		return blk
	}
	for _, v := range blk.Children {
		r := SearchBlk(&v, path)
		if r != nil {
			return r
		}
	}
	return nil
}

func (b *Blockdevices) GetBlkByPath(s string) (*Blockdevice, error) {
	for _, v := range b.Blockdevices {
		b := SearchBlk(&v, s)
		if b != nil {
			return b, nil
		}
	}
	return nil, errors.New("not found")

}
func (b Blockdevices) GetFreeBlockDevices() []Blockdevice {
	var free []Blockdevice
	for _, v := range b.Blockdevices {
		if v.Type == "disk" && v.Fstype == "" && v.Children == nil {
			free = append(free, v)
		}
	}
	return free
}

func GetBlockdevices() (*Blockdevices, error) {
	cmd := exec.Command(LSBLKCmd[0], LSBLKCmd[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var b Blockdevices
	if err := json.Unmarshal(out, &b); err != nil {
		return nil, err
	}
	return &b, nil
}

func GetBlockdevicesOverSSH(ssh *Client) (*Blockdevices, error) {
	c := "sudo " + strings.Join(LSBLKCmd, " ")
	out, err := ssh.Cmd(c).Output()
	if err != nil {
		return nil, err
	}
	var r Blockdevices
	err = json.Unmarshal(out, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
