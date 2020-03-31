package opoll

import "go-scratch/osyscall"

// AcceptFunc is used to hook the accept call.
var AcceptFunc func(int) (int, osyscall.Sockaddr, error) = osyscall.Accept