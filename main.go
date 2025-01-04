package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        int    `json:"_id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello world")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MONGODB ATLAS")

	collection = client.Database("todo_db").Collection("todos")

	app := fiber.New()
	port := os.Getenv(("PORT"))
	if port == "" {
		port = "4000"
	}

	// Get all Todos
	app.Get("/api/todos", getTodos)

	// Create a Todo
	// app.Post("/api/todos", createTodo)

	// Update a Todo
	// app.Patch("/api/todos/:id", updateTodo)

	// Delete a Todo
	// app.Delete("/api/todos/:id", deleteTodo)

	// todos := []Todo{}

	// app.Post("/api/todos", func(c *fiber.Ctx) error {
	// 	todo := &Todo{} // {id:0, completed:false, body:""}

	// 	if err := c.BodyParser(todo); err != nil {
	// 		return err
	// 	}

	// 	if todo.Body == "" {
	// 		return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
	// 	}

	// 	todo.ID = len(todos) + 1
	// 	todos = append(todos, *todo)

	// 	// var x int = 5 // 0x00001
	// 	// var p *int = &x	// 0x00001

	// 	// fmt.Println(p)	// 0x00001
	// 	// fmt.Println(*p) // 5

	// 	return c.Status(201).JSON(todo)
	// })

	// app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")

	// 	for i, todo := range todos {
	// 		if fmt.Sprint(todo.ID) == id {
	// 			todos[i].Completed = true
	// 			return c.Status(200).JSON(todos[i])
	// 		}
	// 	}

	// 	return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	// })

	// app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")

	// 	for i, todo := range todos {
	// 		if fmt.Sprint(todo.ID) == id {
	// 			todos = append(todos[:i], todos[i+1:]...)
	// 			return c.Status(200).JSON(fiber.Map{"success": true})
	// 		}
	// 	}

	// 	return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	// })

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)
}
