package controllers

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/v1")
	{
		// Books routes
		v1.POST("/books", s.PostBook)
		v1.GET("/books", s.GetBooks)
		v1.GET("/books/:id", s.GetBookById)
		v1.PUT("/books/:id", s.UpdateBook)
		v1.DELETE("/books/:id", s.DeleteBook)

		// Genres routes
		v1.GET("/genres", s.GetGenres)

		// Authors routes
		v1.GET("/authors", s.GetAuthors)
	}
}
