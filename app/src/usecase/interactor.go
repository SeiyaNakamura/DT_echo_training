package usecase

import (
	"app/src/entities"
	"app/src/infrastructure/sqlhandler"
)

type Interactor struct {
	Repository Repository
}

// アプリケーション固有のビジネスルール
// このファイルでは取得したデータを組み合わせたりしてユースケースを実現する

func (i Interactor) GetAllArticle() (article []entities.Article, err error) {
	return i.Repository.GetAllArticle()
}

func (i Interactor) GetUndeletedTodos(sqlhandler *sqlhandler.SqlHandler) (todo []entities.Todo, err error) {
	return i.Repository.GetUndeletedTodos(sqlhandler)
}

func (i Interactor) GetTodo(sqlhandler *sqlhandler.SqlHandler, todoId uint) (todo *entities.Todo, err error) {
	return i.Repository.GetTodo(sqlhandler, todoId)
}

func (i Interactor) InsertNewTodo(sqlhandler *sqlhandler.SqlHandler, title string, content string) {
	i.Repository.InsertNewTodo(sqlhandler, title, content)
}

func (i Interactor) UpdateTodo(sqlhandler *sqlhandler.SqlHandler, todoId uint, title string, content string) {
	i.Repository.UpdateTodo(sqlhandler, todoId, title, content)
}

func (i Interactor) DeleteTodo(sqlhandler *sqlhandler.SqlHandler, todoId uint) {
	i.Repository.DeleteTodo(sqlhandler, todoId)
}
