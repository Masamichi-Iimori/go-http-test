package injector

import (
	"github.com/masamichhhhi/go-clean/domain/repository"
	"github.com/masamichhhhi/go-clean/handler"
	"github.com/masamichhhhi/go-clean/infrastructure"
	"github.com/masamichhhhi/go-clean/usecase"
)

func InjectDB() infrastructure.GormHandler {
	sqlhandler := infrastructure.NewGormHandler()
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
