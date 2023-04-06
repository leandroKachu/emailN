package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignGet))

	http.ListenAndServe(":3000", r)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })

	//	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(b)
	//b, _ := json.Marshal(obj)
	// isso que esta comentado e sem o framework do render do proprio go-chi-render
	// obj := map[string]string{"name": "Leandro Leonardo"}
	// agora com o framework do render
	// render.JSON(w, r, obj)

	//	})

	// r.Post("/createProduct", func(w http.ResponseWriter, r *http.Request) {
	// 	var product product

	// 	render.DecodeJSON(r.Body, &product)
	// 	product.ID = 10
	// 	render.JSON(w, r, product)
	// })
}
