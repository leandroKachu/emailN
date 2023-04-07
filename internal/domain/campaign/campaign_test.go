package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

// organizacao 3 A
// A1 = organization
// A2 = action
// A3 = assert

var (
	name     = "NewCampaign"
	content  = "body !Hi"
	contacts = []string{"email1@example.com", "email2@example.com"}
	fake     = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
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

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)
	assert.Equal("name is min5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)
	assert.Equal("name is Max20", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)
	assert.Equal("content is min5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)
	assert.Equal("content is Max1028", err.Error())
}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil)
	assert.Equal("contacts is min1", err.Error())
}

func Test_NewCampaign_MustValidateContactsMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})
	assert.Equal("email is invalid", err.Error())
}

func Test_NewCampaign_MustStartWithStatusPending(t *testing.T) {
	assert := assert.New(t)
	status := "Pending"

	campaign, _ := NewCampaign(name, content, contacts)
	assert.Equal(campaign.Status, status)
}
