package capture

import "net"

type InterfaceInfo struct {
	Name  string   `json:"name"`
	Index int      `json:"index"`
	Flags string   `json:"flags"`
	Addrs []string `json:"addrs"`
}

func ListInterfaces() ([]InterfaceInfo, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var result []InterfaceInfo
	for _, iface := range ifaces {
		var addrs []string
		addrList, _ := iface.Addrs()
		for _, addr := range addrList {
			addrs = append(addrs, addr.String())
		}

		result = append(result, InterfaceInfo{
			Name:  iface.Name,
			Index: iface.Index,
			Flags: iface.Flags.String(),
			Addrs: addrs,
		})
	}

	return result, nil
}
