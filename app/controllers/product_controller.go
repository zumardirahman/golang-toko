package controllers

import (
	"net/http"

	"github.com/unrolled/render"
	"github.com/zumardirahman/golang-toko/app/models"
)

func (server *Server) Products(w http.ResponseWriter, r *http.Request) {

	render := render.New(render.Options{
		Layout: "layout",
	})

	productModel := models.Product{}
	products, err := productModel.GetProducts(server.DB)

	if err != nil {
		return
	}

	//module ini akan merender yang ada didalam module templates
	_ = render.HTML(w, http.StatusOK, "products", map[string]interface{}{
		"products": products,
	})
}
