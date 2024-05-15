package handler

import (
	"net/http"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/domain"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/usecases/events"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MakeEventHandler(router *gin.Engine, uc events.UseCase) {
	router.Use(middleware.GinLogger())
	router.POST("/event", CreateEvent(uc))
	router.GET("/event/:id", GetEvent(uc))
	router.POST("/event/slot", OpenEventSlots(uc))
	router.GET("/event/slot/:id", GetEventSlotsById(uc))
	router.GET("/event/:id/slots", GetEventSlotsByEventId(uc))
	// router.POST("/login", getLogin(service))
	// router.Use(middleware.GinLogger())
	// router.POST("/user", createUser(service))
	// router.POST("/login", getLogin(service))
	// pasetoMaker, _ := paseto.NewPasetoMaker()
	// authRoutes := router.Group("/").Use(middleware.AuthMiddleware(pasetoMaker))
	// authRoutes.GET("/user/:id", getUser(service))
}

func CreateEvent(usecase events.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var req domain.CreateEventDto
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		event, errCreate := usecase.CreateEvent(ctx, req.EventName, req.Note, req.EventOwnerID)

		if errCreate != nil {
			ctx.JSON(http.StatusInternalServerError, errCreate)
		}

		ctx.JSON(http.StatusOK, event)
	})
}

func GetEvent(usecase events.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		event, errCreate := usecase.GetEvent(ctx, id)

		if errCreate != nil {
			ctx.JSON(http.StatusInternalServerError, errCreate)
		}

		ctx.JSON(http.StatusOK, event)
	})
}

func OpenEventSlots(usecase events.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var req domain.CreateEventSlot
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		eventSlot, errCreate := usecase.OpenEventSlots(ctx, &req)

		if errCreate != nil {
			ctx.JSON(http.StatusInternalServerError, errCreate)
		}

		ctx.JSON(http.StatusOK, eventSlot)
	})
}

func GetEventSlotsById(usecase events.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		eventSlot, errCreate := usecase.GetEventSlotsById(ctx, id)

		if errCreate != nil {
			ctx.JSON(http.StatusInternalServerError, errCreate)
		}

		ctx.JSON(http.StatusOK, eventSlot)
	})
}

func GetEventSlotsByEventId(usecase events.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		eventSlots, errCreate := usecase.GetEventSlotsByEventId(ctx, id)

		if errCreate != nil {
			ctx.JSON(http.StatusInternalServerError, errCreate)
		}

		ctx.JSON(http.StatusOK, eventSlots)
	})
}
