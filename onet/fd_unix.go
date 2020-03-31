package onet

import (
	"go-scratch/ointernal/opoll"
)

// Network file descriptor
type netFD struct {
	pfd opoll.FD

	// immutable until Close
	family int
	sotype int
	net string
	laddr Addr
}

func newFD(sysfd, family, sotype int, net string) (*netFD, error) {
	ret := &netFD{
		pfd: opoll.FD{
			Sysfd: sysfd,
		},
		family: family,
		sotype: sotype,
		net: net,
	}
	return ret, nil
}


func (fd *netFD) accept() (netfd *netFD, err error) {
	d, _, errcall, err := fd.pfd.Accept()
	if err != nil {
		if errcall != "" {
		}
		return nil, err
	}
	netfd, err = newFD(d, fd.family, fd.sotype, fd.net)
	if err != nil {
		return nil, err
	}
	return netfd, nil
}
