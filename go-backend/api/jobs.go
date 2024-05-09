package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/EmilioCliff/event-scheduler/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createJobRequest struct {
	Description   string    `json:"description" binding:"required"`
	ClientName    string    `json:"client_name" binding:"required"`
	ClientAddress string    `json:"client_address" binding:"required"`
	ClientEmail   string    `json:"client_email" binding:"required"`
	Price         int64     `json:"price" binding:"required"`
	StartDate     time.Time `json:"start_date" binding:"required"`
	EndDate       time.Time `json:"end_date" binding:"required"`
}

func (server *Server) createJob(ctx *gin.Context) {
	var req createJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	job, err := server.store.CreateJob(ctx, db.CreateJobParams{
		Description:   req.Description,
		ClientName:    req.ClientName,
		ClientAddress: req.ClientAddress,
		ClientEmail:   req.ClientEmail,
		Price:         req.Price,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, job)
}

type getJobRequestURI struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getJob(ctx *gin.Context) {
	var uri getJobRequestURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	job, err := server.store.GetJob(ctx, uri.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, job)
}

func (server *Server) listJob(ctx *gin.Context) {
	jobs, err := server.store.ListJobs(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, jobs)
}

type updateJobRequestURI struct {
	ID int64 `uri:"id"`
}

type updateJobRequest struct {
	Description   string    `json:"description"`
	ClientName    string    `json:"client_name"`
	ClientAddress string    `json:"client_address"`
	ClientEmail   string    `json:"client_email"`
	Price         int64     `json:"price"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
}

func (server *Server) updateJob(ctx *gin.Context) {
	var uri updateJobRequestURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req updateJobRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// TODO: get the job then check if the value if empty if empty update to existing one

	job, err := server.store.UpdateJob(ctx, db.UpdateJobParams{
		Description:   req.Description,
		ClientName:    req.ClientName,
		ClientAddress: req.ClientAddress,
		ClientEmail:   req.ClientEmail,
		Price:         req.Price,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		ID:            uri.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, job)
}

type deleteJobRequestURI struct {
	ID int64 `uri:"id"`
}

func (server *Server) deleteJob(ctx *gin.Context) {
	var uri deleteJobRequestURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := server.store.DeleteJob(ctx, uri.ID); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "deletion of job successful"})
}
