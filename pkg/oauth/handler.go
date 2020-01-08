package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"net/http"
)

type Handler struct {
	Server *server.Server
}

func (o *Handler) Credentials(w http.ResponseWriter, r *http.Request) {
	clientId := uuid.New().String()[:8]
	clientSecret := uuid.New().String()[:8]
	err := ClientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: "http://localhost",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"client_id": clientId, "client_secret": clientSecret})
}

func (o *Handler) Protected(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, I'm protected"))
}

func (o *Handler) Token(w http.ResponseWriter, r *http.Request) {
	srv := o.Server
	if err := srv.HandleTokenRequest(w, r); err != nil {
		fmt.Printf("handle token reques: %v\n", err.Error())
	}
}
