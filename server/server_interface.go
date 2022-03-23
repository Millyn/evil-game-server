package server

type Server interface {
	Init()
	Start()
	Shutdown()
}
