
if err := http.ListenAndServe(":8000", nil); err != nil {
fmt.Println("start http server fail:", err)
}
// net/http/server.go:L3002-3005
func ListenAndServe(addr string, handler Handler) error {
server := &Server{Addr: addr, Handler: handler}
return server.ListenAndServe()
}
// net/http/server.go:L2752-2765
func (srv *Server) ListenAndServe() error {
// ... 省略代码
ln, err := net.Listen("tcp", addr) // <-----看这里listen
if err != nil {
return err
}
return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}
// net/http/server.go:L2805-2853
func (srv *Server) Serve(l net.Listener) error {
// ... 省略代码
for {
rw, e := l.Accept() // <----- 看这里accept
if e != nil {
select {
case <-srv.getDoneChan():
return ErrServerClosed
default:
}
if ne, ok := e.(net.Error); ok && ne.Temporary() {
if tempDelay == 0 {
tempDelay = 5 * time.Millisecond
} else {
tempDelay *= 2
}
if max := 1 * time.Second; tempDelay > max {
tempDelay = max
}
srv.logf("http: Accept error: %v; retrying in %v", e, tempDelay)
time.Sleep(tempDelay)
continue
}
return e
}
tempDelay = 0
c := srv.newConn(rw)
c.setState(c.rwc, StateNew) // before Serve can return
go c.serve(ctx) // <--- 看这里
}
}
// net/http/server.go:L1739-1878
func (c *conn) serve(ctx context.Context) {
// ... 省略代码
serverHandler{c.server}.ServeHTTP(w, w.req)
w.cancelCtx()
if c.hijacked() {
return
}
w.finishRequest()
// ... 省略代码
}
// net/http/server.go:L2733-2742
func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
handler := sh.srv.Handler
if handler == nil {
handler = DefaultServeMux
}
if req.RequestURI == "*" && req.Method == "OPTIONS" {
handler = globalOptionsHandler{}
}
handler.ServeHTTP(rw, req)
}
// net/http/server.go:L2352-2362
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
if r.RequestURI == "*" {
if r.ProtoAtLeast(1, 1) {
w.Header().Set("Connection", "close")
}
w.WriteHeader(StatusBadRequest)
return
}
h, _ := mux.Handler(r) // <--- 看这里
h.ServeHTTP(w, r)
}
go复制代码// net/http/server.go:L1963-1965
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
f(w, r)
}
