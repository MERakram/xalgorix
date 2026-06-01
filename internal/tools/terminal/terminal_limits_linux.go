//go:build linux

package terminal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

// setProcessLimitsForPID applies OOM / RLIMIT_AS constraints on Linux.
// The OOM score adjustment is best-effort because it may require elevated privileges.
func setProcessLimitsForPID(pid int, memoryLimited bool, memLimitBytes int64) {
	if pid <= 0 {
		return
	}

	oomPath := fmt.Sprintf("/proc/%d/oom_score_adj", pid)
	if err := os.WriteFile(oomPath, []byte("500"), 0o644); err != nil {
		log.Printf("[RESOURCES] Cannot set OOM score for PID %d: %v", pid, err)
	}

	if memoryLimited && memLimitBytes > 0 {
		newLimit := syscall.Rlimit{
			Cur: uint64(memLimitBytes),
			Max: uint64(memLimitBytes),
		}
		_, _, errno := syscall.RawSyscall6(
			syscall.SYS_PRLIMIT64,
			uintptr(pid),
			uintptr(syscall.RLIMIT_AS),
			uintptr(unsafe.Pointer(&newLimit)),
			0,
			0,
			0,
		)
		if errno != 0 {
			log.Printf("[RESOURCES] Cannot set RLIMIT_AS for PID %d: %v", pid, errno)
		} else {
			log.Printf("[RESOURCES] Tool PID %d: OOM score=500, mem limit=%d MB", pid, memLimitBytes/(1024*1024))
		}
	} else {
		log.Printf("[RESOURCES] PID %d: OOM score set to 500", pid)
	}
}

func setProcessLimits(cmd *exec.Cmd, memoryLimited bool, memLimitBytes int64) {
	if cmd.Process == nil {
		return
	}
	setProcessLimitsForPID(cmd.Process.Pid, memoryLimited, memLimitBytes)
}
