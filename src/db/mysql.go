package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SND1231/go-column/setting"
)

func GetDBconnection(dbSetting setting.DB) (*sql.DB, error) {
	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbSetting.User,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.Port,
		dbSetting.Name,
	)
	dataSource = dataSource + "&loc=Asia%2FTokyo"
	db, err := sql.Open(dbSetting.Type, dataSource)
	return db, err
}

func NewUnitTestDB() *UnitTestDB {
	conn, err := GetDBconnectionForTest()
	if err != nil {
		panic(err)
	}
	return &UnitTestDB{
		conn: conn,
	}
}

func GetDBconnectionForTest() (*sql.DB, error) {
	dbSetting := setting.DB{
		Type:     "mysql",
		User:     "root",
		Password: "test",
		Host:     "mysql",
		Port:     3306,
		Name:     "unit_test",
	}
	return GetDBconnection(dbSetting)
}

type UnitTestDB struct {
	conn *sql.DB
}

// 他のテストに影響がないように、テスト終了後にロールバックをしてテスト前と同じ状態にする。
func (db *UnitTestDB) InTx(exec func(context.Context, *sql.Tx)) {
	if db.conn == nil {
		panic("connection is nil")
	}

	ctx := context.Background()
	tx, err := db.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		panic(err)
	}
	// ロールバックしてテスト中のデータ作成や更新は他のテストに影響ないようにしている。
	defer tx.Rollback()

	exec(ctx, tx)

}
