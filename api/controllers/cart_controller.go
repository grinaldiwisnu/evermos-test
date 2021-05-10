package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-awesomeapi/api/formaterror"
	"go-awesomeapi/api/models"
	"go-awesomeapi/api/responses"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (s *Server) AddToCart(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	cart := models.CartItem{}

	cart.Prepare()
	err = json.Unmarshal(body, &cart)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = cart.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	cartCreated, err := cart.AddToCart(s.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, cart.ID))
	responses.JSON(w, http.StatusCreated, cartCreated)
}

func (s *Server) GetCartByUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	cart := models.CartItem{}
	itemFound, err := cart.FindAllCartByUID(s.DB, pid)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, itemFound)
}
