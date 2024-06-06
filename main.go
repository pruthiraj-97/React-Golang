package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct{
	ID primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	Completed bool           `json:"completed"`
	Body string              `json:"body"`
}

var collection *mongo.Collection

func main(){
	app:=fiber.New()
	// load the .env file
	err:=godotenv.Load(".env")
	if err!=nil {
      log.Fatal("Error loading .env file")
	}
	MONGODB_URI:=os.Getenv("MONGODB_URI")
	ClientOptions:=options.Client().ApplyURI(MONGODB_URI)
	client,err:=mongo.Connect(context.Background(),ClientOptions)
	if err!= nil {
		log.Fatal(err)
	}
    err=client.Ping(context.Background(),nil)
	if err!=nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB ATLAS")
	collection=client.Database("golangTodo").Collection("Todo")

    app.Get("/api/gettodos",getTodos)
	app.Post("/api/addtodo",postTodo)
	app.Patch("/api/updatetodo/:id",updateTodo)
	app.Delete("/api/deletetodo/:id",deleteTodo)

    PORT:=os.Getenv("PORT")
	if PORT=="" {
		PORT="4000"
	}
	log.Fatal(app.Listen(":"+PORT))
}

func getTodos(c *fiber.Ctx)error{
	var todos []Todo
    cursor,err:=collection.Find(context.Background(),bson.M{})  // bson.M{} for the filteration of data
	if err!=nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()){
		var todo Todo
		if err:=cursor.Decode(&todo);err!=nil{
			return err
		}
		todos=append(todos, todo)
	}
	return c.Status(200).JSON(fiber.Map{
		"msg":todos,
	})
}

func postTodo(c *fiber.Ctx)error{
	todo:=new(Todo)
	err:=c.BodyParser(todo)
	if err!=nil {
		return err
	}
	if todo.Body==""{
		return c.Status(400).JSON(fiber.Map{
			"msg":"invalid body",
		})
	}
	insertResult,err:=collection.InsertOne(context.Background(),todo)
	if err!=nil {
		return err
	}
	todo.ID=insertResult.InsertedID.(primitive.ObjectID)
	return c.Status(200).JSON(fiber.Map{
		"msg":"Todo crrated",
		"todo":todo,
	})
}

func updateTodo(c *fiber.Ctx)error{
	id:=c.Params("id")
	objectId,err:=primitive.ObjectIDFromHex(id)
	if err!=nil {
		return c.Status(400).JSON(fiber.Map{
			"msg":"invalid id",
		})
	}
	filter:=bson.M{"_id":objectId}
	update:=bson.M{"$set":bson.M{"completed":true}}
	_,err=collection.UpdateOne(context.Background(),filter,update)
	if err!=nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{
		"msg":"Todo updated",
	})
}

func deleteTodo(c *fiber.Ctx)error{
	id:=c.Params("id")
	objectId,err:=primitive.ObjectIDFromHex(id)
	if err!=nil {
       return c.Status(400).JSON(fiber.Map{
		"err":"id not found",
	   })
	}
	filter:=bson.M{"_id":objectId}
	_,err=collection.DeleteOne(context.Background(),filter)
	if err!=nil {
       return err
	}
	return c.Status(200).JSON(fiber.Map{
		"msg":"Todo deleted",
	})
}