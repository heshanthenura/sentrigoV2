#!/bin/bash
clang -O2 -g -target bpf -c ebpf/interceptor.c -o ebpf/interceptor.o
go build -o main cmd/sentrigoV2/main.go
sudo ./main