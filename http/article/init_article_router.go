package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	domain "ximlr/go-lang-web/domain"
	repos "ximlr/go-lang-web/repositories"

	"github.com/julienschmidt/httprouter"
)

func Init(http *httprouter.Router) {

	http.GET("/a/:id", getArticleById)
	http.GET("/a", getArticles)
	http.PUT("/a", createArticle)
	http.PATCH("/a", updateArticle)
}

func getArticleById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	param := p.ByName("id")

	i, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, "bad income data", http.StatusBadRequest)
		fmt.Println("ID=" + param + "; incorrect params; HTTP 400 Bad Request")
		return
	}

	if i < 0 {
		http.Error(w, "bad income data", http.StatusBadRequest)
		fmt.Println("ID=" + param + "; incorrect params; HTTP 400 Bad Request")
		return
	}

	repo := repos.ArticleRepository{}

	result, err := repo.GetById(i)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

	fmt.Println("ID=" + param + "; HTTP 200 OK")
}

func getArticles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	repo := repos.ArticleRepository{}

	result, _ := repo.Get()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)

	fmt.Println("HTTP 200 OK")
}

func createArticle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var model *domain.ArticleModel

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, "bad income data", http.StatusBadRequest)
		return
	}

	repo := repos.ArticleRepository{}

	err = repo.Create(model)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateArticle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var model *domain.ArticleModel

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, "bad income data", http.StatusBadRequest)
		return
	}

	repo := repos.ArticleRepository{}
	err = repo.Update(model)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
