package router

import (
	"github.com/SND1231/go-column/handler"
	"github.com/SND1231/go-column/setting"
	"github.com/go-chi/chi"
)

func Get(dbSetting setting.DB) *chi.Mux {
	r := chi.NewRouter()

	// ハンドラーの初期化
	userHandler := handler.NewUserHandler(dbSetting)

	// httpルーティング
	r.Route("/user", func(r chi.Router) {
		r.Post("/add", userHandler.Add)
		r.Get("/detail", userHandler.Get)
	})
	return r
}
