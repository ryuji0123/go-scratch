package opoll

import "go-scratch/osyscall"

// FD is a file descriptor. The net and os packages use this type as a
// field of a larger type representing a network connection or OS file.
type FD struct {
	// System file descriptor. Immutable until Close.
	Sysfd int
}

// Accept wraps the accept network call.
func (fd *FD) Accept() (int, osyscall.Sockaddr, string, error) {
	for {
		s, rsa, errcall, err := accept(fd.Sysfd)
		if err == nil {
			return s, rsa, "", err
		}
		return -1, nil, errcall, err
	}
}