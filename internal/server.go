package internal

type Server struct{}

func (s *Server) Run() error {
	return nil
}

func NewServer() Server {
	return Server{}
}
