package ohttp

import (
	"errors"
	"go-scratch/onet"
	"net/http"
)

type Handler interface {
	ServeHTTP(ResponseWriter, *http.Request)
}

type ResponseWriter interface {
	Write([]byte) (int, error)
}

type response struct {
	written int64
	status int
}

func (w *response) Write(data []byte) (n int, err error) {
	return w.write(len(data), data, "")
}

func (w *response) write(lenData int, dataB []byte, dataS string) (n int, err error) {
	return
}

type ServeMux struct {
	m map[string]muxEntry
}

type muxEntry struct {
	h Handler
	pattern string
}

var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux

// Handle registers the handler for the given pattern.
// If a handler already exists for pattern, Handle panics.
func (mux *ServeMux) Handle(pattern string, handler Handler) {

	if pattern == "" {
		panic("http: invalid pattern")
	}

	if handler == nil {
		panic("http: nil handler")
	}

	if _, exist := mux.m[pattern]; exist {
		panic("http: multiple registrations for " + pattern)
	}

	if mux.m == nil {
		mux.m = make(map[string]muxEntry)
	}

	e := muxEntry{
		h:       handler,
		pattern: pattern,
	}
	mux.m[pattern] = e

}


type HandlerFunc func(ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *http.Request) {
	f(w, r)
}

type Server struct {
	Addr string
	Handler Handler
}

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*onet.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (onet.Conn, error) {
	tc, _ := ln.AcceptTCP()
	return tc, nil
}

var (
	// ServerContextKey is a context key. It can be used in HTTP
	// handlers with context.WithValue to access the server that
	// started the handler. The associated value will be of
	// type *Server.
	ServerContextKey = &contextKey{name: "http-server"}
)

// onceCloseListener wraps a net.Listener, protecting it from
// multiple Close calls.
type onceCloseListener struct {
	onet.Listener
	closeErr error
}

// Serve accepts incoming connections on the Listener l, creating a
// new service goroutine for each. The service goroutines read requests and
// then call srv.Handler to reply to them.
//
// HTTP/2 support is only enabled if the Listener returns *tls.Conn
// connections and they were configured with "h2" in the TLS
// Config.NextProtos.
//
// Serve always returns a non-nil error and closes l.
// After Shutdown or Close, the returned error is ErrServerClosed.
func (srv *Server) Serve(l onet.Listener) error {
	l = &onceCloseListener{Listener: l}
	//baseCtx := ocontext.Background() // base is always background, per Issue 16220 baseCtx: unknown empty Context
	//ctx := ocontext.WithValue(baseCtx, ServerContextKey, srv)
	//for {
	//	rw, e := l.Accept()

	//}
	return errors.New("")
}


func (srv *Server) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	//ln, err := onet.Listen()
	//if err != nil {
	//	return err
	//}
	//return srv.Serve(tcpKeepAliveListener{ln.(*onet.TCPListener)})
	return errors.New("developing")
}

// Handle registers the handler for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func Handle(pattern string, handler Handler) {DefaultServeMux.Handle(pattern, handler)}