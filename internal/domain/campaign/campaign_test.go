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
	campaign, _ := NewCampaign(name, content, contacts)
	//A3 assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	//A1
	assert := assert.New(t)
	//A2
	campaign, _ := NewCampaign(name, content, contacts)
	//A3 assert
	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreateOnMustBeNow(t *testing.T) {
	//A1
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)
	//A2
	campaign, _ := NewCampaign(name, content, contacts)
	//A3 assert
	assert.Greater(campaign.CreatedOn, now)
}

// Testing errors handling ()()()()()()()()()()()()()()()()()()()()()()()()()

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)
	assert.Equal("name must not be empty", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)
	assert.Equal("content must not be empty", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})
	assert.Equal("contacts must not be empty", err.Error())
}
