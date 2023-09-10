package services

import (
	"context"
	"github.com/google/uuid"
	"grpc/internal/pb"
)

type ContactService struct {
	pb.UnimplementedContactServiceServer
	contacts []pb.Contact
}

func NewContactService() *ContactService {
	return &ContactService{}
}

func (c *ContactService) Create(ctx context.Context, input *pb.CreateContactRequest) (*pb.Contact, error) {
	c.contacts = append(c.contacts, pb.Contact{
		Uuid:        uuid.New().String(),
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
	})

	return &pb.Contact{
		Uuid:        uuid.New().String(),
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
	}, nil
}
