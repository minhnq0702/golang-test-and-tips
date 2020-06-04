package main

import (
	"MYGO/studying/go-swagger/gen/models"
	"MYGO/studying/go-swagger/gen/restapi"
	"MYGO/studying/go-swagger/gen/restapi/operations"
	"MYGO/studying/go-swagger/gen/restapi/operations/todos"
	myModel "MYGO/studying/go-swagger/models"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/go-openapi/loads"

	"github.com/danhtran94/copi"
)

var portFlag = flag.Int("port", 3000, "Port to run service")

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err.Error())
	}
	api := operations.NewTodolistAPI(swaggerSpec)

	value := "Todo 1"
	lstItem := []*myModel.QItem{
		&myModel.QItem{
			Completed:   true,
			Description: value,
			ID:          1,
		},
		&myModel.QItem{
			Completed:   false,
			Description: value,
			ID:          2,
		},
	}

	api.TodosFindTodoHandler = todos.FindTodoHandlerFunc(
		func(params todos.FindTodoParams) middleware.Responder {
			res := []*models.Item{}

			err = copi.Dup(&lstItem, &res)
			fmt.Println("=========>", res, lstItem)
			if err != nil {
				log.Fatalln(err.Error())
			}

			return todos.NewFindTodoOK().WithPayload(res)
		})
	api.TodosAddTodoHandler = todos.AddTodoHandlerFunc(
		func(params todos.AddTodoParams) middleware.Responder {
			params.Body.ID = 3
			var qitem *myModel.QItem
			err := copi.Dup(&params.Body, &qitem)

			if err != nil {
				todos.NewAddTodoDefault(http.StatusInternalServerError)
			}

			lstItem = append(lstItem, qitem)
			fmt.Println("sau khi add", lstItem, qitem)

			// res = append(res, params.Body)
			return todos.NewAddTodoDefault(todos.AddTodoCreatedCode)
		})

	// server := restapi.NewServer(api)

	flag.Parse()

	err = api.Validate()
	if err != nil {
		log.Fatalln(err.Error())
	}

	http.ListenAndServe("localhost:8000", api.Serve(nil))
}
