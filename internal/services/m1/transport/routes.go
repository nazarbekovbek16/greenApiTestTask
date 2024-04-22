package transport

func (s *Server) routes() {
	v1 := s.HTTP.Group("/api/v1")

	v1.POST("/double", s.h.Test.DoubleInteger)
}
