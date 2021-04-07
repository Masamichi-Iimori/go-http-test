package injector

import (
	"github.com/go-http-test/handler"
	"github.com/go-http-test/infrastructure"
	"github.com/go-http-test/repository"
	"github.com/go-http-test/usecase"
)

func InjectDB() infrastructure.SqlHandler {
	sqlhandler := infrastructure.NewSQLHandler()
	return *sqlhandler
}

/*
TodoRepository(interface)に実装であるSqlHandler(struct)を渡し生成する。
*/
func InjectTodoRepository() repository.TodoRepository {
	sqlHandler := InjectDB()
	return infrastructure.NewTodoRepository(sqlHandler)
}

func InjectTodoUsecase() usecase.TodoUsecase {
	TodoRepo := InjectTodoRepository()
	return usecase.NewTodoUsecase(TodoRepo)
}

func InjectTodoHandler() handler.TodoHandler {
	return handler.NewTodoHandler(InjectTodoUsecase())
}
