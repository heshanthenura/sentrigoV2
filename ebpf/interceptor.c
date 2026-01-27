#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>

#define ETH_P_IP 0x0800

struct
{
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 1024);
    __type(key, __u32); // Host byte order IP
    __type(value, __u8);
} blocked_ips SEC(".maps");

SEC("xdp")
int block_ips(struct xdp_md *ctx)
{
    void *data = (void *)(long)ctx->data;
    void *data_end = (void *)(long)ctx->data_end;

    // 1. Parse Ethernet Header
    struct ethhdr *eth = data;
    if ((void *)(eth + 1) > data_end)
        return XDP_PASS;

    if (bpf_ntohs(eth->h_proto) != ETH_P_IP)
        return XDP_PASS;

    // 2. Parse IP Header
    struct iphdr *ip = data + sizeof(*eth);
    if ((void *)(ip + 1) > data_end)
        return XDP_PASS;

    // 3. Lookup blocked IPs (store map keys in host byte order)
    __u32 src_ip = bpf_ntohl(ip->saddr); // convert to host order
    __u8 *exists = bpf_map_lookup_elem(&blocked_ips, &src_ip);

    // 4. Drop packet if IP exists
    if (exists)
    {
        bpf_printk("DROP IP: %u\n", src_ip);
        return XDP_DROP;
    }

    return XDP_PASS;
}

char LICENSE[] SEC("license") = "GPL";
