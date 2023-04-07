package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=20"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1028"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `validate:"required"`
}

const (
	Pending = "Pending"
	Started = "iniciado"
	Done    = "done"
)

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}

	camapaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
	}
	err := internalerrors.ValidateStruct(camapaign)

	if err == nil {
		return camapaign, nil
	}

	return nil, err
}
