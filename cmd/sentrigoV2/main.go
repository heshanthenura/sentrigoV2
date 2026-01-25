package main

import (
	"log"
	"net"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
)

func main() {
	spec, err := ebpf.LoadCollectionSpec("ebpf/drop_packets.o")
	if err != nil {
		log.Fatal(err)
	}

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		log.Fatal(err)
	}
	defer coll.Close()

	log.Println("eBPF program loaded successfully")

	prog, ok := coll.Programs["drop_icmp"]
	if !ok {
		log.Fatal("XDP program not found in collection")
	}

	iface, err := net.InterfaceByName("docker0")
	if err != nil {
		log.Fatal(err)
	}

	xdpLink, err := link.AttachXDP(link.XDPOptions{
		Program:   prog,
		Interface: iface.Index,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer xdpLink.Close()

	log.Println("XDP program attached to", iface.Name)

	select {}
}
