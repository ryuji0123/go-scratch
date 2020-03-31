package opoll

import "go-scratch/osyscall"

// Wrapper around the accept system call that marks the returned file
// descriptor as nonblocking and close-on-exec.
func accept(s int) (int, osyscall.Sockaddr, string, error) {
	// See ../syscall/exec_unix.go for description of ForkLock.
	// It is probably okay to hold the lock across syscall.Accept
	// because we have put fd.sysfd into non-blocking mode.
	// However, a call to the File method will put it back into
	// blocking mode. We can't take that risk, so no use of ForkLock here.
	ns, sa, err := AcceptFunc(s)
	if err != nil {
		return -1, nil, "accept", err
	}
	return ns, sa, "", nil
}
