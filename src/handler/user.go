package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mholt/binding"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
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

// ユーザーの作成APIのリクエスト
type AddUserInput struct {
	Name string
	Age  int
}

// リクエストのマッピング。ポインターレシーバーにすること
func (input *AddUserInput) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&input.Name: "name",
		&input.Age:  "age",
	}
}

// ユーザーの作成APIのレスポンス
type AddUserOutput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ユーザーの作成API
func (h *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	var err error
	var berr binding.Errors
	var response *AddUserOutput
	var res []byte

	// request -> AddUserInput型に変換
	var request AddUserInput
	berr = binding.Bind(r, &request)
	if berr != nil {
		log.Println(berr)
		setErrorResponse(w, http.StatusInternalServerError)
		return
	}

	// responseの作成
	response = &AddUserOutput{
		ID:   1,
		Name: request.Name,
		Age:  request.Age,
	}

	// レスポンスをjsonに変換
	res, err = json.Marshal(response)
	if err != nil {
		log.Println(err)
		setErrorResponse(w, http.StatusInternalServerError)
		return
	}
	setSuccessResponse(w, res)
}

// ユーザーの詳細取得APIのリクエスト
type GetUserInput struct {
	ID int
}

// リクエストのマッピング。ポインターレシーバーにすること
func (input *GetUserInput) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&input.ID: "id",
	}
}

// ユーザーの詳細取得APIのレスポンス
type GetUserOutput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ユーザーの詳細取得API
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	var err error
	var berr binding.Errors
	var response *GetUserOutput
	var res []byte

	// request -> GetUserInput型に変換
	var request GetUserInput
	berr = binding.Bind(r, &request)
	if berr != nil {
		log.Println(berr)
		setErrorResponse(w, http.StatusInternalServerError)
		return
	}

	// responseの作成
	response = &GetUserOutput{
		ID:   request.ID,
		Name: "Jony",
		Age:  45,
	}

	// レスポンスをjsonに変換
	res, err = json.Marshal(response)
	if err != nil {
		log.Println(err)
		setErrorResponse(w, http.StatusInternalServerError)
		return
	}
	setSuccessResponse(w, res)
}
