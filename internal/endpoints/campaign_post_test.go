package endpoints

import (
	"bytes"
	"emailn/internal/domain/campaign/contract"
	internalMock "emailn/internal/test/mock"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_campaign_post_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "Leandro teste",
		Content: "bodyiee",
		Emails:  []string{"Leandrogaia@hotmail.com", "i like a new campaign"},
	}
	service := new(internalMock.CampaignServiceMock)

	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {

		if request.Name == body.Name && request.Content == body.Content {
			return true
		} else {
			return false
		}

	})).Return("23dxaf", nil)

	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)

}
