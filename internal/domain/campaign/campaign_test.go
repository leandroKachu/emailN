package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// organizacao 3 A
// A1 = organization
// A2 = action
// A3 = assert

var (
	name     = "NewCampaign"
	content  = "body"
	contacts = []string{"email1@example.com", "email2@example.com"}
)

func Test_NewCampaign(t *testing.T) {
	//A1
	assert := assert.New(t)
	//A2
	campaign := NewCampaign(name, content, contacts)
	//A3 assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	//A1
	assert := assert.New(t)
	//A2
	campaign := NewCampaign(name, content, contacts)
	//A3 assert
	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreateOnMustBeNow(t *testing.T) {
	//A1
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)
	//A2
	campaign := NewCampaign(name, content, contacts)
	//A3 assert
	assert.Greater(campaign.CreatedOn, now)
}
