package onet

import "go-scratch/ocontext"

// An addrList represents a list of network endpoint addresses.
type addrList []Addr

// SplitHostPort splits a network address of the form "host:port",
// "host%zone:port", "[host]:port" or "[host%zone]:port" into host or
// host%zone and port.
//
// A literal IPv6 address in hostport must be enclosed in square
// brackets, as in "[::1]:80", "[::1%lo0]:80".
//
// See func Dial for a description of the hostport parameter, and host
// and port results.
func SplitHostPort(hostport string) (host, port string, err error) {
	const (
		missingPort   = "missing port in address"
		tooManyColons = "too many colons in address"
	)
	addrErr := func(addr, why string) (host, port string, err error) {
		return "", "", &AddrError{Err: why, Addr: addr}
	}

	// The port starts after the last colon.
	i := last(hostport, ':')

	if i < 0 {
		return addrErr(hostport, missingPort)
	}

	if hostport[0] == '[' {

	} else {
		host = hostport[:i]
	}

	port = hostport[i+1:]
	return host, port, nil
}

// internetAddrList resolves addr, which may be a literal IP
// address or a DNS name, and returns a list of internet protocol
// family addresses. The result contains at least one address when
// error is nil.
func (r *Resolver) internetAddrList(ctx ocontext.Context, net, addr string) (addrList, error) {
	var (
		err error
		host, port string
		portnum int
	)
	switch net {
	case "tcp":
		if addr != "" {
			if host, port, err = SplitHostPort(addr); err != nil {
				return nil, err
			}
			if portnum, err = r.LookupPort(ctx, net, port); err != nil {
				return nil, err
			}
		}
	default:
		return nil, UnknownNetworkError(net)
	}
	inetaddr := func(ip IPAddr) Addr {
		switch net {
		case "tcp":
			return &TCPAddr{IP: ip.IP, Port: portnum, Zone: ip.Zone}
		default:
			panic("unexpected network: " + net)
		}
	}
	if host == "" {
		return addrList{inetaddr(IPAddr{})}, nil
	}
	return addrList{inetaddr(IPAddr{})}, nil
}
