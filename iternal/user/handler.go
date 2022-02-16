package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"userGoServ/iternal/apperror"
	"userGoServ/iternal/handlers"
	"userGoServ/pgk/logging"
)

const (
	usersURL = "/users"
	userURL  = "/user"
)

type handler struct {
	logger     *logging.Logger
	repository Repository
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUser))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.PutUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {

	all, err := h.repository.FindAll(context.TODO())
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)

	return nil
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) error {
	user, err := h.repository.FindOne(context.TODO(), r.FormValue("username"))
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	allBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)

	return nil
}

func (h *handler) PutUser(w http.ResponseWriter, r *http.Request) error {
	var DaysInWeek []DaysInWeek

	json.Unmarshal([]byte(r.FormValue("DaysInWeek")), &DaysInWeek)

	usr := User{
		Username:        r.FormValue("Username"),
		Email:           r.FormValue("Email"),
		Password:        r.FormValue("Password"),
		Level:           r.FormValue("Level"),
		DaysInRow:       r.FormValue("DaysInRow"),
		DaysInWeek:      DaysInWeek,
		DoesSendPushUps: false,
		Theme:           r.FormValue("Theme"),
		Language:        r.FormValue("Language"),
		Image:           r.FormValue("Image"),
	}
	err := h.repository.Update(context.TODO(), usr)
	if err != nil {
		w.WriteHeader(400)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here is a string...."))

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	err := h.repository.Delete(context.TODO(), r.FormValue("bodu"))
	if err != nil {
		fmt.Println(err)

		w.WriteHeader(400)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here is a string...."))

	return nil
}
