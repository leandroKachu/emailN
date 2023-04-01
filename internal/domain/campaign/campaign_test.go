package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	name := "NewCampaign"
	content := "body"
	contacts := []string{"email1@example.com", "email2@example.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
}
