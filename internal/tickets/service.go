package tickets

import (
	"context"
	"desafio-go-web-freymanlozano/internal/domain"
)

type service struct {
	r Repository
}

type Service interface {
	GetTotalTickets(ctx context.Context, dest string) (ticket []domain.Ticket, err error)
	AverageDestination(ctx context.Context, dest string) (float64, error)
}

func NewService(repo Repository) Service {
    return &service{
		r: repo,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, dest string) (
	ticket []domain.Ticket, err error) {
    return s.r.GetTicketByDestination(ctx, dest)
}

func (s *service) AverageDestination(ctx context.Context, dest string) (float64, error) {
	ticketsByDest, err := s.r.GetTicketByDestination(ctx, dest)
	if err!= nil {
        return 0, err
    }
	totalTickets, err := s.r.GetAll(ctx)
	if err!= nil {
        return 0, err
    }
	avg := float64(len(ticketsByDest)) / float64(len(totalTickets))
	return avg, nil
}