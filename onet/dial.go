package onet

import (
	"go-scratch/ocontext"
	"go-scratch/osyscall"
)

type ListenConfig struct {
	// If Control is not nil, it is called after creating the network
	// connection but before binding it to the operating system.
	//
	// Network and address parameters passed to Control method are not
	// necessarily the ones passed to Listen. For example, passing "tcp" to
	// Listen will cause the Control function to be called with "tcp4" or "tcp6".
	Control func(network, address string, c osyscall.RawConn) error
}

func parseNetwork(ctx ocontext.Context, network string, needsProto bool) (afnet string, proto int, err error) {
	i := last(network, ':')
	if i < 0 {
		switch network {
		case "tcp":
		default:
			return "", 0, UnknownNetworkError(network)
		}
		return network, 0, nil
	}
	return "", 0, UnknownNetworkError(network)
}
//
func (r *Resolver) resolveAddrList(ctx ocontext.Context, op, network, addr string, hint Addr) (addrList, error) {
	afnet, _, err := parseNetwork(ctx, network, true)
	if err != nil {
		return nil, err
	}
	addrs, err := r.internetAddrList(ctx, afnet, addr)
	if err != nil || op != "dial" || hint == nil {
		return addrs, err
	}
	return addrs, err
}

//func (lc *ListenConfig) Listen(ctx ocontext.Context, network, address string) (Listener, error) {
//	addrs, err := DefaultResolver.resolveAddrList(ctx, "listen", network, address, nil)
//	return Listener(), err
//}
//
//func Listen(network, address string) (Listener, error) {
//	var lc ListenConfig
//	return lc.Listen(ocontext.Background(), network, address)
//}
