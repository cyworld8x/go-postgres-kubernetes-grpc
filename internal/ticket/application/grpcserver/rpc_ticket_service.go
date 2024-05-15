package grpcserver

import (
	"context"

	"github.com/cyworld8x/go-postgres-kubernetes-grpc/pkg/pb"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateTickets(ctx context.Context, req *pb.CreateTicketsRequest) (*pb.CreateTicketsResponse, error) {
	eventSlotId, errInput := uuid.Parse(req.EventSlotId)
	if errInput != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse event slot id: %v", errInput)
	}

	newIssued, err := server.uc.CreateTicketsByEventSlot(ctx, eventSlotId, float64(req.Price), req.Capacity)

	if err != nil {
		log.Error().Err(err).Msg("cannot create tickets")
		return nil, status.Errorf(codes.Internal, "cannot create tickets: %v", err)
	}

	total, err := server.uc.GetTotalTicketByEventSlot(ctx, eventSlotId)
	if err != nil {
		log.Error().Err(err).Msg("cannot get total tickets")
	}

	return &pb.CreateTicketsResponse{
		Issued:      newIssued,
		Total:       total,
		EventSlotId: req.EventSlotId,
	}, nil
}
func (server *Server) SellTicket(ctx context.Context, req *pb.SellTicketRequest) (*pb.SellTicketResponse, error) {
	eventSlotId, errInput := uuid.Parse(req.EventSlotId)

	if errInput != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse event slot id: %v", errInput)
	}

	buyerId, errInput := uuid.Parse(req.BuyerId)
	if errInput != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse buyer id: %v", errInput)
	}

	ticket, err := server.uc.SellTicket(ctx, eventSlotId, buyerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot sell ticket: %v", err)
	}

	return &pb.SellTicketResponse{
		Code:        ticket.Code,
		EventSlotId: ticket.EventSlotID.String(),
		Status:      ticket.Status,
		Issued:      ticket.Issued.String(),
		BuyerId:     buyerId.String(),
	}, nil
}
func (server *Server) CheckIn(context.Context, *pb.CheckinTicketRequest) (*pb.CheckinTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIn not implemented")
}
