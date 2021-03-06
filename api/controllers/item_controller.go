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

func (s *Server) CreateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Items")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	item := models.Item{}

	item.Prepare()
	err = json.Unmarshal(body, &item)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = item.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	fmt.Printf(item.Title)

	itemCreated, err := item.CreateItem(s.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, item.ID))
	responses.JSON(w, http.StatusCreated, itemCreated)
}

func (s *Server) GetItems(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Get Items")
	item := models.Item{}

	items, err := item.FindAllItem(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, items)
}

func (s *Server) GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	item := models.Item{}
	itemFound, err := item.FindItemByID(s.DB, pid)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, itemFound)
}

func (s *Server) UpdateItem(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	item := models.Item{}

	err = json.Unmarshal(body, &item)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	itemUpdate, err := item.UpdateItem(s.DB)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, itemUpdate)
}
