// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package local

import (
	"github.com/opspec-io/opctl/pkg/node"
	"os/exec"
	"syscall"
)

func (this nodeProvider) CreateNode() (nodeInfo *node.InfoView, err error) {
	nodeCmd := exec.Command(
		"opctl",
		"node",
		"create",
	)

	nodeCmd.SysProcAttr = &syscall.SysProcAttr{
		// ensure node gets it's own process group
		Setpgid: true,
	}

	err = nodeCmd.Start()
	if nil != err {
		panic(err)
	}

	nodeInfo = &node.InfoView{}

	return
}
