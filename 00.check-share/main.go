package main

import (
	"fmt"
	"rdma-share/utils"

	"github.com/vishvananda/netlink"
)

// 查看RDMA是否是shared
// rdma system show netns
func main() {
	mode, err := netlink.RdmaSystemGetNetnsMode()
	utils.Must(err)

	fmt.Println(mode)
}
