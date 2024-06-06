import React from "react";
import { useState,useEffect } from "react";
import { MdDelete,MdClose } from "react-icons/md";
import { FaCheck } from 'react-icons/fa';
import { backend_url } from "./backend_url";
import './Todo.css'
function Todo({body,completed,id,deleteTodo}){
    const [isCompleted,SetIncompleted]=useState(completed)
    async function updateTodo(id){
        const response=await fetch(`${backend_url}/api/updatetodo/${id}`,{
            method:'PATCH',
            headers:{
                'Content-Type':'application/json'
            }
        })
        const data=await response.json()
        console.log(data)
        if(data.status==200){
            SetIncompleted(!isCompleted)
        }
    }
    return (
        <div>
            <div className="Todo-div">
                <div>{body}</div>{
                    isCompleted?
                    <FaCheck className="right-icon" onClick={()=>updateTodo(id)} />
                    :<MdClose className="right-icon" onClick={()=>updateTodo(id)}/>
                }
                <MdDelete className="delete-icon" onClick={()=>deleteTodo(id)}/>
            </div>
        </div>
    )
}

export default Todo