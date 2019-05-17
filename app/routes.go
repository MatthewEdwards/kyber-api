package app

func (app *application) routes() {
	app.router.HandleFunc("/api/v1/articles", app.handleGetArticles()).Methods("GET")
	app.router.HandleFunc("/api/v1/articles", app.handlePostArticle()).Methods("POST")
}
