package usecase

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/SND1231/go-column/models"
	"github.com/mholt/binding"
	"github.com/volatiletech/sqlboiler/v4/boil"

	_ "github.com/go-sql-driver/mysql"
)

type UserUsecase struct {
	ctx context.Context
	tx  *sql.Tx
}

func NewUserUsecase(ctx context.Context, tx *sql.Tx) *UserUsecase {
	return &UserUsecase{
		ctx: ctx,
		tx:  tx,
	}
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

func (u UserUsecase) Add(input AddUserInput) (AddUserOutput, error) {
	user := models.User{
		Name: input.Name,
		Age:  input.Age,
	}
	err := user.Insert(u.ctx, u.tx, boil.Infer())
	if err != nil {
		return AddUserOutput{}, err
	}
	output := AddUserOutput{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}
	return output, nil
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

func (u UserUsecase) Get(input GetUserInput) (GetUserOutput, error) {
	user, err := models.FindUser(u.ctx, u.tx, input.ID)
	if err != nil {
		return GetUserOutput{}, nil
	}
	output := GetUserOutput{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}
	return output, nil
}
