package handler

import (
	"net/http"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/domain"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/event/usecases/events"
	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MakeEventHandler(router gin.IRouter, uc events.UseCase) {
	router.Use(middleware.GinLogger())
	router.POST("/events", CreateEvent(uc))
	router.GET("/events/:id", GetEvent(uc))
	router.POST("/events/slot", OpenEventSlots(uc))
	router.GET("/events/slot/:id", GetEventSlotsById(uc))
	router.GET("/events/:id/slots", GetEventSlotsByEventId(uc))
	// router.POST("/login", getLogin(service))
	// router.Use(middleware.GinLogger())
	// router.POST("/user", createUser(service))
	// router.POST("/login", getLogin(service))
	// pasetoMaker, _ := paseto.NewPasetoMaker()
	// authRoutes := router.Group("/").Use(middleware.AuthMiddleware(pasetoMaker))
	// authRoutes.GET("/user/:id", getUser(service))
}

// Get Event godoc
// @Summary      Create event
// @Description  Create event
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        arg  body domain.CreateEventDto true "Event Info"
// @Success      200  {object}  domain.Event
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /events [Post]
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

// Get Event godoc
// @Summary      get an event
// @Description  get event by Id
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        id   path string  true "Id"
// @Success      200  {object}  domain.Event
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /events/{id} [get]
func GetEvent(usecase events.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		event, errGet := usecase.GetEvent(ctx, id)

		if errGet != nil {
			ctx.JSON(http.StatusInternalServerError, errGet)
		}

		ctx.JSON(http.StatusOK, event)
	})
}

// Get Event godoc
// @Summary      Create slots for an event
// @Description  create slot for an event
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        id   path string  true "Event ID"
// @Param        arg  body domain.CreateEventSlot true "Event Slot Info"
// @Success      200  {object}  domain.EventSlot
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /events/slot [Post]
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

// Get Event godoc
// @Summary      get slots of an event
// @Description  get slot of an event by event Id
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        id   path string  true "ID"
// @Success      200  {object}  []domain.EventSlot
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /events/:id/slots [get]
func GetEventSlotsById(usecase events.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		eventSlot, errGet := usecase.GetEventSlotsById(ctx, id)

		if errGet != nil {
			ctx.JSON(http.StatusInternalServerError, errGet)
		}

		ctx.JSON(http.StatusOK, eventSlot)
	})
}

// Get Event godoc
// @Summary      get slots of an event
// @Description  get slot event by event Id
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        id   path string  true "id"
// @Success      200  {object}  domain.EventSlot
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /events/slot/{id} [get]
func GetEventSlotsByEventId(usecase events.UseCase) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		eventSlots, errGet := usecase.GetEventSlotsByEventId(ctx, id)

		if errGet != nil {
			ctx.JSON(http.StatusInternalServerError, errGet)
		}

		ctx.JSON(http.StatusOK, eventSlots)
	})
}
