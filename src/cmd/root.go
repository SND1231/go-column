/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/SND1231/go-column/router"
	"github.com/SND1231/go-column/setting"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dbSetting setting.DB

var rootCmd = &cobra.Command{
	Use: "go-column",
	Run: func(cmd *cobra.Command, args []string) {
		// configの中身を出力
		log.Printf("configの中身:{type: %s, host: %s, port: %d, user: %s, pass: %s, name: %s}",
			dbSetting.Type, dbSetting.Host, dbSetting.Port, dbSetting.User, dbSetting.Password, dbSetting.Name)

		// サーバーの設定
		r := router.Get(dbSetting)
		srv := &http.Server{
			Addr:    ":3020",
			Handler: r,
		}

		// サーバーの起動
		srv.ListenAndServe()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 初期化処理
	// configの設定
	rootCmd.Flags().StringP("configName", "n", "default.toml", "config file name")

	// Runを実行するたびに、initConfigを呼び出す。その後に、Runの処理が動き出す。
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	configName, _ := rootCmd.Flags().GetString("configName")
	viper.SetConfigFile(configName)

	// 設定ファイルを読み込む
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// 設定ファイルの内容を構造体に設定
	if err := viper.Unmarshal(&dbSetting); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
