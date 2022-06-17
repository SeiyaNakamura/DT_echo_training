package main

import "app/src/controller"

func main() {
	controller.Init()

	//articleテーブルのデータを出力
	//var article []models.Article
	//db := models.DatabaseConnection()
	//articles := db.Find(&article)
	//fmt.Println(articles)

	//DB接続テスト
	//_, err := models.DatabaseConnection()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("接続成功しました")
}
