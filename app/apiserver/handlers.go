package apiserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/openmind13/http-api-chat/app/model"
)

// POST
// http://localhost:9000/users/add
func (s *server) handleAddUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	type request struct {
		Username string `json:"username"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	// create user model
	user := &model.User{
		Username: req.Username,
	}
	user.CreatedAt = time.Now()

	// store user in database
	id, err := s.store.AddUser(user)
	if err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	user.ID = id
	s.respondJSON(w, r, http.StatusCreated, user.ID)
}

// GET
// http:/localhost:9000/users/get
func (s *server) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	users, err := s.store.GetAllUsers()
	if err != nil {
		s.error(w, r, http.StatusNotFound, err)
		return
	}

	s.respondJSON(w, r, http.StatusOK, users)
}

// POST
// http://localhost:9000/chats/add
func (s *server) handleAddChat(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name  string   `json:"name"`
		Users []string `json:"users"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
	}

	chat := &model.Chat{
		Name:  req.Name,
		Users: req.Users,
	}

	if err := s.store.AddUsersIntoChat(chat); err != nil {
		s.error(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	s.respondJSON(w, r, http.StatusCreated, chat.ID)
}

// POST
// http://localhost:9000/messages/add
func (s *server) handleAddMessage(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Chat   int    `json:"chat"`
		Author int    `json:"author"`
		Text   string `json:"text"`
	}

	req := &request{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}

	message := &model.Message{
		Chat:   req.Chat,
		Author: req.Author,
		Text:   req.Text,
	}

	if err := s.store.AddMessageIntoChat(message); err != nil {
		// handle error
		s.error(w, r, http.StatusUnprocessableEntity, err)
	}

	s.respondJSON(w, r, http.StatusInternalServerError, message.ID)
}

// POST
// http://localhost:9000/chats/get
func (s *server) handleGetUserChats(w http.ResponseWriter, r *http.Request) {
	type request struct {
		User string `json:"user"`
	}

	s.respondJSON(w, r, http.StatusInternalServerError, nil)
}

// POST
// http://localhost:9000/messages/get
func (s *server) handleGetChatMessages(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Chat int `json:"chat"`
	}

	s.respondJSON(w, r, http.StatusInternalServerError, nil)
}
