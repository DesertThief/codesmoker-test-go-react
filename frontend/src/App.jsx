import { useState, useEffect } from 'react'
import './App.css'

const API_URL = '/api/todos'

function App() {
  const [todos, setTodos] = useState([])
  const [newTodo, setNewTodo] = useState('')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  // Fetch todos on mount
  useEffect(() => {
    fetchTodos()
  }, [])

  const fetchTodos = async () => {
    try {
      setLoading(true)
      const response = await fetch(API_URL)
      if (!response.ok) throw new Error('Failed to fetch todos')
      const data = await response.json()
      setTodos(data)
      setError(null)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const addTodo = async (e) => {
    e.preventDefault()
    if (!newTodo.trim()) return

    try {
      const response = await fetch(API_URL, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name: newTodo, isComplete: false }),
      })
      if (!response.ok) throw new Error('Failed to add todo')
      const todo = await response.json()
      setTodos([...todos, todo])
      setNewTodo('')
    } catch (err) {
      setError(err.message)
    }
  }

  const toggleTodo = async (id, isComplete) => {
    try {
      const response = await fetch(`${API_URL}/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ isComplete: !isComplete }),
      })
      if (!response.ok) throw new Error('Failed to update todo')
      const updated = await response.json()
      setTodos(todos.map(t => t.id === id ? updated : t))
    } catch (err) {
      setError(err.message)
    }
  }

  const deleteTodo = async (id) => {
    try {
      const response = await fetch(`${API_URL}/${id}`, {
        method: 'DELETE',
      })
      if (!response.ok) throw new Error('Failed to delete todo')
      setTodos(todos.filter(t => t.id !== id))
    } catch (err) {
      setError(err.message)
    }
  }

  return (
    <div className="app">
      <header>
        <h1>✅ Go + React Todo App</h1>
        <p className="subtitle">Full-stack with Go Gin API + React Frontend</p>
      </header>

      {error && <div className="error">{error}</div>}

      <form onSubmit={addTodo} className="add-form">
        <input
          type="text"
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
          placeholder="Add a new todo..."
          className="todo-input"
        />
        <button type="submit" className="add-btn">Add</button>
      </form>

      {loading ? (
        <p className="loading">Loading...</p>
      ) : (
        <ul className="todo-list">
          {todos.map(todo => (
            <li key={todo.id} className={`todo-item ${todo.isComplete ? 'completed' : ''}`}>
              <input
                type="checkbox"
                checked={todo.isComplete}
                onChange={() => toggleTodo(todo.id, todo.isComplete)}
              />
              <span className="todo-name">{todo.name}</span>
              <button
                onClick={() => deleteTodo(todo.id)}
                className="delete-btn"
              >
                ×
              </button>
            </li>
          ))}
        </ul>
      )}

      {todos.length === 0 && !loading && (
        <p className="empty">No todos yet. Add one above!</p>
      )}
    </div>
  )
}

export default App
