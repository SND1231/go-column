package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/SND1231/go-column/db"
	"github.com/SND1231/go-column/setting"
	"github.com/SND1231/go-column/usecase"
	"github.com/mholt/binding"
)

type UserHandler struct {
	dbSetting setting.DB
}

func NewUserHandler(dbSetting setting.DB) *UserHandler {
	return &UserHandler{
		dbSetting: dbSetting,
	}
}

// 失敗時のレスポンスの設定
func setErrorResponse(w http.ResponseWriter, status int) {
	// レスポンスのヘッダ設定
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// 引数のステータス設定
	w.WriteHeader(status)
}

// 成功時のレスポンスの設定
func setSuccessResponse(w http.ResponseWriter, res []byte) {
	// レスポンスの内容をjsonに変換
	w.Write(res)
}

// ユーザーの作成API
func (h *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	var err error
	var berr binding.Errors
	var response usecase.AddUserOutput
	var res []byte

	// request -> AddUserInput型に変換
	var request usecase.AddUserInput
	berr = binding.Bind(r, &request)
	if berr != nil {
		log.Println(berr)
		setErrorResponse(w, http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	// dbのコネクション設定
	conn, err := db.GetDBconnection(h.dbSetting)
	if err != nil {
		log.Println(err)
		return
	}
	// コネクションのクローズ
	defer func() {
		conn.Close()
	}()

	// トランザクションの開始
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// トランザクションのコミット or ロールバック
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// usecaseの初期化
	userUsecase := usecase.NewUserUsecase(ctx, tx)
	// ユーザー作成実施
	response, err = userUsecase.Add(request)

	// レスポンスをjsonに変換
	res, err = json.Marshal(response)
	if err != nil {
		log.Println(err)
		setErrorResponse(w, http.StatusInternalServerError)
		return
	}
	setSuccessResponse(w, res)
}

// ユーザーの詳細取得API
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	var err error
	var berr binding.Errors
	var response usecase.GetUserOutput
	var res []byte

	// request -> GetUserInput型に変換
	var request usecase.GetUserInput
	berr = binding.Bind(r, &request)
	if berr != nil {
		log.Println(berr)
		setErrorResponse(w, http.StatusInternalServerError)
		return
	}

	ctx := context.Background()

	// dbのコネクション設定
	conn, err := db.GetDBconnection(h.dbSetting)
	if err != nil {
		log.Println(err)
		return
	}
	// コネクションのクローズ
	defer func() {
		conn.Close()
	}()

	// トランザクションの開始
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// トランザクションのコミット or ロールバック
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// usecaseの初期化
	userUsecase := usecase.NewUserUsecase(ctx, tx)
	response, err = userUsecase.Get(request)

	// レスポンスをjsonに変換
	res, err = json.Marshal(response)
	if err != nil {
		log.Println(err)
		setErrorResponse(w, http.StatusInternalServerError)
		return
	}
	setSuccessResponse(w, res)
}
