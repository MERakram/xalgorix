//go:build !linux

package terminal

import "os/exec"

func setProcessLimitsForPID(pid int, memoryLimited bool, memLimitBytes int64) {
	_ = pid
	_ = memoryLimited
	_ = memLimitBytes
}

func setProcessLimits(cmd *exec.Cmd, memoryLimited bool, memLimitBytes int64) {
	if cmd.Process == nil {
		return
	}
	setProcessLimitsForPID(cmd.Process.Pid, memoryLimited, memLimitBytes)
}
