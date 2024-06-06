import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'
import TodoSection from './components/TodoSection.jsx'
ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <TodoSection />
  </React.StrictMode>,
)
