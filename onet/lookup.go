package onet

import "go-scratch/ocontext"

// DefaultResolver is the resolver used by the package-level Lookup
// functions and by Dialers without a specified Resolver.
var DefaultResolver = &Resolver{}

// A Resolver looks up names and numbers.
//
// A nil *Resolver is equivalent to a zero Resolver.
type Resolver struct {

}

func (r *Resolver) LookupPort(ctx ocontext.Context, network, service string) (port int, err error) {
	port, needslookup := parsePort(service)
	if needslookup {
		switch network {
		case "tcp":
		default:
			return 0, &AddrError{Err: "unknown network", Addr: network}
		}
	}
	return port, nil
}

// lookupIPAddr looks up host using the local resolver and particular network.
// It returns a slice of that host's IPv4 and IPv6 addresses.
func (r *Resolver) lookupIPAddr(ctx ocontext.Context, network, host string) ([]IPAddr, error) {
	// Make sure that no matter what we do later, host=="" is rejected.
	// parseIP, for example, does accept empty strings.
	if ip, zone := parseIPZone(host); ip != nil {
		return []IPAddr{{IP: ip, Zone: zone}}, nil
	}
	return []IPAddr{{}}, nil
}
