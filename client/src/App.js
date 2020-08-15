import React from 'react'
import User from "./components/User.js"

const App = () => {
  return (
    <div className="container">
      <h1 id="title">Content Loader</h1>
      <button>Load</button>
      <User/>
    </div>
  )
}

export default App
