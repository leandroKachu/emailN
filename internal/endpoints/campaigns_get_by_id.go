package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	// campaigns, err := h.CampaignService.Repository.Get()

	id := chi.URLParam(r, "id")

	campagin, err := h.CampaignService.GetBy(id)

	if err == nil && campagin == nil {
		return nil, http.StatusNotFound, err
	}
	return campagin, 200, err
}
