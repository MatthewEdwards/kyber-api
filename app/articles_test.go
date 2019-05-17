package app

import (
	"bytes"
	"encoding/json"
	"kyber-api/datastore"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

type MockApplication struct {
	db MockDatastore
}

type MockDatastore struct {
	Session mock.Mock
}

func (c *MockDatastore) GetArticles() (articles []datastore.Article) {
	articles = []datastore.Article{
		datastore.Article{Title: "Example article", URL: "www.example.com", Source: "Local"},
		datastore.Article{Title: "Example article 2", URL: "www.example.com", Source: "Local"},
	}

	return articles
}

func (c *MockDatastore) AddArticle(article datastore.Article) (err error) {
	return nil
}

func TestHandleGetAricles(t *testing.T) {
	app := &application{
		router: mux.NewRouter().StrictSlash(true),
		db:     &MockDatastore{},
	}

	req, err := http.NewRequest("GET", "/api/v1/articles", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(app.handleGetArticles())
	handler.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
	}
}

func TestHandlePostAricle(t *testing.T) {
	app := &application{
		router: mux.NewRouter().StrictSlash(true),
		db:     &MockDatastore{},
	}

	article := datastore.Article{
		Title:  "Utility equipment sparked massive California wildfire, investigators say",
		URL:    "https://arstechnica.com/tech-policy/2019/05/fire-officials-say-utilitys-power-lines-were-responsible-for-cas-deadliest-fire/",
		Source: "Ars Technica",
	}

	s, _ := json.Marshal(article)
	b := bytes.NewBuffer(s)

	req, err := http.NewRequest("POST", "/api/v1/articles", b)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(app.handlePostArticle())
	handler.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
	}
}
