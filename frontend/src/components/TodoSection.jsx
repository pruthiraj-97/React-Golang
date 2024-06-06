import React from "react";
import { useState,useEffect } from "react";
import { backend_url } from "./backend_url";
import Todo from "./Todo";
import './TodoSection.css'
function TodoSection(){
    const [Todos,setTodos]=useState([])
    const [newTodo,setnewTodo]=useState('')
    useEffect(()=>{
       (async ()=>{
          const response=await fetch(`${backend_url}/api/gettodos`,{
            method:'GET',
            headers:{
              'Content-Type':'application/json'
            }
          })
          const data=await response.json()
          console.log(data)
          setTodos(data.todos)
       })()
    },[])
    async function addNewTodo(e){
        e.preventDefault()
        const response=await fetch(`${backend_url}/api/addtodo`,{
            method:'POST',
            headers:{
                'Content-Type':'application/json'
            },
            body:JSON.stringify({
                body:newTodo
            })
        })
        const data=await response.json()
        console.log(data)
        setnewTodo('')
        if(data.status==200){
            setTodos([...Todos,data.todo])
        }
    }
    async function deleteTodo(id){
       const response=await fetch(`${backend_url}/api/deletetodo/${id}`,{
           method:'DELETE',
           headers:{
               'Content-Type':'application/json'
           }
       })
       const data=await response.json()
       console.log(data)
       if(data.status==200){
           setTodos(Todos.filter((todo)=>todo.id!=id))
       }
    }
    return (
        <div className="Todosection">
            <div className="Todo-input-section">
                <input type="text" onChange={(e)=>setnewTodo(e.target.value)} />
                <button onClick={addNewTodo}>submit</button>
            </div>
            <div className="AllTodos-div">
            {
                Todos.length>0&&
                Todos.map((todo)=>{
                    return (
                        <Todo key={todo.id} body={todo.body}
                          deleteTodo={deleteTodo}
                        completed={todo.completed} id={todo.id}/>
                    )
                })
            }
            </div>
        </div>
    )
}

export default TodoSection