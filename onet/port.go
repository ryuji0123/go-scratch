package onet

// parsePort parses service as a decimal integer and returns the
// corresponding value as port. It is the caller's responsibility to
// parse service as a non-decimal integer when needsLookup is true.
//
// Some system resolvers will return a valid port number when given a number
// over 65536 (see https://golang.org/issues/11715). Alas, the parser
// can't bail early on numbers > 65536. Therefore reasonably large/small
// numbers are parsed in full and rejected if invalid.
func parsePort(service string) (port int, needslookup bool) {
	const (
		max = uint32(1 << 32 - 1)
		cutoff = uint32(1 << 30)
	)
	//neg := false
	var n uint32
	for _, d := range service {
		if '0' <= d && d <= '9' {
			d -= '0'
		}
		if n >= cutoff {
			n = max
			break
		}
		n *= 10
		nn := n + uint32(d)
		if nn < n || nn > max {
			n = max
			break
		}
		n = nn
	}
	port = int(n)
	return port, false
}
