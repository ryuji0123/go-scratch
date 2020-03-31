package onet


// IP address lengths (bytes).
const (
	IPv4len = 4
)


// An IP is a single IP address, a slice of bytes.
// Functions in this package accept either 4-byte (IPv4)
// or 16-byte (IPv6) slices as input.
//
// Note that in this documentation, referring to an
// IP address as an IPv4 address or an IPv6 address
// is a semantic property of the address, not just the
// length of the byte slice: a 16-byte slice can still
// be an IPv4 address.
type IP []byte

// Parse IPv4 address (d.d.d.d).
func parseIPv4(s string) IP {
	//var p [IPv4len]byte
	for i := 0; i < IPv4len; i++ {
		if len(s) == 0 {
			// Missing octets.
			return nil
		}
		if i > 0 {

		}
	}
	return nil
}

func parseIPZone(s string) (IP, string) {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return parseIPv4(s), ""
		}
	}
	return nil, ""
}
