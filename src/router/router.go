package router

import (
	"log"
	"net/http"

	"github.com/SND1231/go-column/handler"
	"github.com/SND1231/go-column/setting"
	"github.com/go-chi/chi"
)

func Get(dbSetting setting.DB) *chi.Mux {
	r := chi.NewRouter()

	// ハンドラーの初期化
	userHandler := handler.NewUserHandler(dbSetting)
	printDBsetting := makePrintDBSetting(dbSetting)

	r.Use(printDBsetting)
	// httpルーティング
	r.Route("/user", func(r chi.Router) {
		r.Post("/add", userHandler.Add)
		r.Get("/detail", userHandler.Get)
	})
	return r
}

func makePrintDBSetting(dbSetting setting.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return printDBSetting(next, dbSetting)
	}
}

func printDBSetting(next http.Handler, dbSetting setting.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("type:%s, host:%s", dbSetting.Type, dbSetting.Host)
		next.ServeHTTP(w, r)
	})
}
