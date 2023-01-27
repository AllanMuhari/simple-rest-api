package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var tod = []todo{
	{ID: "1", Item: "Coffee", Completed: false},
	{ID: "2", Item: "Juice", Completed: false},
	{ID: "3", Item: "Tea", Completed: false},
}

func getTodoById(id string) (*todo, error) {
	for i, t := range tod {
		if t.ID == id {
			return &tod[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {

	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {

		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func getTod(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, tod)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	tod = append(tod, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}
func main() {

	router := gin.Default()
	router.GET("/tod", getTod)
	router.GET("/tod/:id", getTod)
	router.POST("/tod", addTodo)
	router.Run("localhost:8080")
}
