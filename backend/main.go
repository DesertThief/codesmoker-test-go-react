package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Todo represents a todo item
type Todo struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	IsComplete bool   `json:"isComplete"`
}

var (
	todos  = []Todo{
		{ID: 1, Name: "Learn Go", IsComplete: false},
		{ID: 2, Name: "Build API with Gin", IsComplete: false},
		{ID: 3, Name: "Connect React frontend", IsComplete: false},
	}
	nextID = 4
	mu     sync.Mutex
)

func main() {
	r := gin.Default()

	// CORS configuration for React frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// API routes
	api := r.Group("/api")
	{
		api.GET("/todos", getTodos)
		api.GET("/todos/:id", getTodo)
		api.POST("/todos", createTodo)
		api.PUT("/todos/:id", updateTodo)
		api.DELETE("/todos/:id", deleteTodo)
	}

	// Root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Go + React Todo API",
			"endpoints": gin.H{
				"GET /api/todos":     "Get all todos",
				"GET /api/todos/:id": "Get todo by ID",
				"POST /api/todos":    "Create a todo",
				"PUT /api/todos/:id": "Update a todo",
				"DELETE /api/todos/:id": "Delete a todo",
			},
		})
	})

	r.Run(":8080")
}

func getTodos(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	c.JSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func createTodo(c *gin.Context) {
	var input struct {
		Name       string `json:"name" binding:"required"`
		IsComplete bool   `json:"isComplete"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	todo := Todo{
		ID:         nextID,
		Name:       input.Name,
		IsComplete: input.IsComplete,
	}
	nextID++
	todos = append(todos, todo)

	c.JSON(http.StatusCreated, todo)
}

func updateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input struct {
		Name       string `json:"name"`
		IsComplete *bool  `json:"isComplete"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, todo := range todos {
		if todo.ID == id {
			if input.Name != "" {
				todos[i].Name = input.Name
			}
			if input.IsComplete != nil {
				todos[i].IsComplete = *input.IsComplete
			}
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
