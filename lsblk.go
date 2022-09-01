package lvm2

import (
	"encoding/json"
	"os/exec"
)

var (
	LSBLKCommand = []string{"lsblk", "-O", "--json"}
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
	Size       string        `json:"size"`
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
	DiscGran   string        `json:"disc-gran"`
	DiscMax    string        `json:"disc-max"`
	DiscZero   bool          `json:"disc-zero"`
	Wsame      string        `json:"wsame"`
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

func (b Blockdevices) GetInfoOfBlock() {

}
func (b Blockdevices) GetFreeBlockDevices() []string {
	var free []string
	for _, v := range b.Blockdevices {
		if v.Type == "disk" && v.Fstype == "" && v.Children == nil {
			free = append(free, v.Path)
		}
	}
	return free
}

func GetBlockdevices() (*Blockdevices, error) {
	cmd := exec.Command(LSBLKCommand[0], LSBLKCommand[1:]...)
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
