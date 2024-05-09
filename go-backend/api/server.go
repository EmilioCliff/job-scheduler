package api

import (
	db "github.com/EmilioCliff/event-scheduler/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Queries
	router *gin.Engine
}

func NewServer(store *db.Queries) *Server {
	server := &Server{
		store: store,
	}

	server.setRoutes()
	return server
}

func (server *Server) setRoutes() {
	r := gin.Default()

	r.POST("/jobs", server.createJob)
	r.GET("/jobs/:id", server.getJob)
	r.GET("/jobs", server.listJob)
	r.PUT("/jobs/:id", server.updateJob)
	r.DELETE("/jobs/:id", server.deleteJob)

	server.router = r
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
