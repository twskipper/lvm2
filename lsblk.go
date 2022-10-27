package lvm2

import (
	"bufio"
	"errors"
	"os/exec"
	"strings"
)

var (

	//LSBLKCmd = []string{"lsblk", "-O", "--json", "-b"}
	//lsblk -p -n -b -r -s -o NAME,SIZE,TYPE,FSTYPE
	LSBLKCmd = []string{"lsblk", "-p", "-n", "-b", "-r", "-s", "-o", "NAME,SIZE,TYPE,FSTYPE"}
)

type Blockdevices struct {
	Blockdevices []Blockdevice `json:"blockdevices"`
}

// type Blockdevice struct {
// 	Name       string        `json:"name"`
// 	Kname      string        `json:"kname"`
// 	Path       string        `json:"path"`
// 	MajMin     string        `json:"maj:min"`
// 	Fsavail    string        `json:"fsavail"`
// 	Fssize     string        `json:"fssize"`
// 	Fstype     string        `json:"fstype"`
// 	Fsused     string        `json:"fsused"`
// 	Fsuse      string        `json:"fsuse%"`
// 	Mountpoint string        `json:"mountpoint"`
// 	Label      interface{}   `json:"label"`
// 	UUID       string        `json:"uuid"`
// 	Ptuuid     string        `json:"ptuuid"`
// 	Pttype     string        `json:"pttype"`
// 	Parttype   interface{}   `json:"parttype"`
// 	Partlabel  interface{}   `json:"partlabel"`
// 	Partuuid   interface{}   `json:"partuuid"`
// 	Partflags  interface{}   `json:"partflags"`
// 	Ra         int           `json:"ra"`
// 	Ro         bool          `json:"ro"`
// 	Rm         bool          `json:"rm"`
// 	Hotplug    bool          `json:"hotplug"`
// 	Model      string        `json:"model"`
// 	Serial     string        `json:"serial"`
// 	Size       int           `json:"size"`
// 	State      string        `json:"state"`
// 	Owner      string        `json:"owner"`
// 	Group      string        `json:"group"`
// 	Mode       string        `json:"mode"`
// 	Alignment  int           `json:"alignment"`
// 	MinIo      int           `json:"min-io"`
// 	OptIo      int           `json:"opt-io"`
// 	PhySec     int           `json:"phy-sec"`
// 	LogSec     int           `json:"log-sec"`
// 	Rota       bool          `json:"rota"`
// 	Sched      string        `json:"sched"`
// 	RqSize     int           `json:"rq-size"`
// 	Type       string        `json:"type"`
// 	DiscAln    int           `json:"disc-aln"`
// 	DiscGran   int           `json:"disc-gran"`
// 	DiscMax    int           `json:"disc-max"`
// 	DiscZero   bool          `json:"disc-zero"`
// 	Wsame      int           `json:"wsame"`
// 	Wwn        interface{}   `json:"wwn"`
// 	Rand       bool          `json:"rand"`
// 	Pkname     interface{}   `json:"pkname"`
// 	Hctl       interface{}   `json:"hctl"`
// 	Tran       string        `json:"tran"`
// 	Subsystems string        `json:"subsystems"`
// 	Rev        string        `json:"rev"`
// 	Vendor     string        `json:"vendor"`
// 	Zoned      string        `json:"zoned"`
// 	Children   []Blockdevice `json:"children,omitempty"`
// }

//  lsblk -p -n -b -r -s -o NAME,SIZE,TYPE,FSTYPE
type Blockdevice struct {
	Name   string `json:"name"`
	Fstype string `json:"fstype"`
	Size   string `json:"size"`
	Type   string `json:"type"`
}

func (b *Blockdevices) GetBlkByPath(s string) (*Blockdevice, error) {
	for _, v := range b.Blockdevices {
		if v.Name == s {
			return &v, nil
		}
	}
	return nil, errors.New("not found")

}
func (b Blockdevices) GetFreeBlockDevices() []Blockdevice {
	var free []Blockdevice
	for _, v := range b.Blockdevices {
		if v.Type == "disk" && v.Fstype == "" && !b.CalculateChildren(v.Name) {
			free = append(free, v)
		}
	}
	return free
}
func (b Blockdevices) CalculateChildren(path string) bool {
	var re bool
	for _, v := range b.Blockdevices {
		if v.Name != path && strings.HasPrefix(v.Name, path) {
			re = true
		}
	}
	return re
}

func GetBlockdevices() (*Blockdevices, error) {
	cmd := exec.Command(LSBLKCmd[0], LSBLKCmd[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	r, err := DecodeLineTextFromLSBLK(out)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetBlockdevicesOverSSH(ssh *Client) (*Blockdevices, error) {
	c := "sudo " + strings.Join(LSBLKCmd, " ")
	out, err := ssh.Cmd(c).Output()
	if err != nil {
		return nil, err
	}
	r, err := DecodeLineTextFromLSBLK(out)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// lsblk -p -n -b -r -s -o NAME,SIZE,TYPE,FSTYPE
func DecodeLineTextFromLSBLK(output []byte) (*Blockdevices, error) {
	var r Blockdevices
	var blk []Blockdevice
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		var b Blockdevice
		text := scanner.Text()
		r := strings.Fields(text)
		if len(r) < 3 {
			return nil, errors.New("invalid line")
		} else if len(r) == 3 {
			b.Name = r[0]
			b.Size = r[1]
			b.Type = r[2]
			b.Fstype = ""
		} else {
			b.Name = r[0]
			b.Size = r[1]
			b.Type = r[2]
			b.Fstype = r[3]
		}
		blk = append(blk, b)
	}
	r.Blockdevices = blk
	return &r, nil
}
