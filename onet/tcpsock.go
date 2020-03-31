package onet

// TCPAddr represents the address of a TCP end point.
type TCPAddr struct {
	IP IP
	Port int
	Zone string // IPv6 scoped addressing zone
}

func (a *TCPAddr) String() string {
	if a == nil {
		return "<nil>"
	}
	return "developing"
}


// TCPConn is an implementation of the Conn interface for TCP network
// connections.
type TCPConn struct {
	conn
}

func newTCPConn(fd *netFD) *TCPConn {
	c := &TCPConn{conn{fd:fd}}
	return c
}

// TCPListener is a TCP network listener. Clients should typically
// use variables of type Listener instead of assuming TCP.
type TCPListener struct {
	fd *netFD
}

// Accept implements the Accept method in the Listener interface; it
// waits for the next call and returns a generic Conn.
//func (l *TCPListener) Accept() (Conn, error) {

//}

// AcceptTCP accepts the next incoming call and returns the new
// connection.
func (l *TCPListener) AcceptTCP() (*TCPConn, error) {
	c, err := l.accept()
	if err != nil {
		return nil, &OpError{Op: "accept", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
	}
	return c, nil
}
