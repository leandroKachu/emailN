package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// organizacao 3 A
// A1 = organization
// A2 = action
// A3 = assert
func TestNewCampaign(t *testing.T) {
	//A1
	assert := assert.New(t)
	name := "NewCampaign"
	content := "body"
	contacts := []string{"email1@example.com", "email2@example.com"}
	//A2
	campaign := NewCampaign(name, content, contacts)
	//A3 assert
	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
}
