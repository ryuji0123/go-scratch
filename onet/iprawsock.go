package onet

// IPAddr represents the address of an IP end point.
type IPAddr struct {
	IP IP
	Zone string //IPv6 scoped addressing zone
}