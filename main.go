package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct{
	ID int          `json:"id"`
	Completed bool  `json:"completed"`
	Body string     `json:"body"`
}
// Array of todos
var todos=[]Todo{}

func main(){
   fmt.Println("hello go,fiber framwork")
   app:=fiber.New()

   err:=godotenv.Load(".env")
   if err!=nil {
	   log.Fatal("Error loading .env file")
   }
   
   PORT:=os.Getenv("PORT")

   app.Get("/",func(c *fiber.Ctx)error{
       return c.Status(200).JSON(fiber.Map{"msg":"well comw to fiber server"})
   })
   
   app.Get("/api/todos",func(c *fiber.Ctx) error {
	   return c.Status(200).JSON(fiber.Map{"todos":todos})
   })

   app.Post("/api/addtodo",func (c *fiber.Ctx)error {
	  todo:=&Todo{}
	  // validation of data
	  if err:=c.BodyParser(todo); err!=nil{
		return err
	  }
	  if(todo.Body==""){
		return c.Status(400).JSON(fiber.Map{
			"err":"todo length is zero",
		})
	  }
	  todo.ID=len(todos)+1
	  todos = append(todos,*todo)
	  return c.Status(200).JSON(todo)
   })

   // Update todo
   app.Patch("/api/updatetoto/:id",func(c *fiber.Ctx) error {
	  id:=c.Params("id")
	  for i,todo:=range todos {
        if fmt.Sprint(todo.ID)==id {
             todos[i].Completed=!todos[i].Completed
			 return c.Status(200).JSON(fiber.Map{"msg":"Todo updated"})
		}
	  }

	  return c.Status(400).JSON(fiber.Map{
		"err":"id not found",
	  })
   })

   app.Delete("/api/deletetodo/:id",func(c *fiber.Ctx) error {
	   id:=c.Params("id")
	   for i,todo:=range todos{
		  if(fmt.Sprint(todo.ID)==id){
			todos=append(todos[:i],todos[i+1:]...)
			return c.Status(200).JSON(fiber.Map{
				"msg":"Todo deleted",
			})
		  }
	   }
	   return c.Status(200).JSON(fiber.Map{
		"err":"Id not found",
	   })
   })

   log.Fatal(app.Listen(":"+PORT))
}