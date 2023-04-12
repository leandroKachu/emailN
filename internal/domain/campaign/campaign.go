package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	ID         string `validate:"required" gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignID string ` gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50"`
	Name      string    `validate:"min=5,max=20" gorm:"size:20"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1028" gorm:"size:1028"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `validate:"required" gorm:"size:20"`
}

const (
	Pending  = "Pending"
	Started  = "iniciado"
	Done     = "done"
	Canceled = "canceled"
)

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
		contacts[index].ID = xid.New().String()

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
