package app

func (app *application) routes() {
	app.router.HandleFunc("/api/v1/articles", app.handleGetAricles()).Methods("GET")
	//app.router.HandleFunc("/api/v1/articles", app.handlePostAricle).Methods("POST")
}
