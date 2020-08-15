import React from 'react'
import User from "./components/User.js"

class App extends React.Component {
  state = {
    loading: true,
    user: null
  }

  async loadUsers() {
    var url = "http://localhost:8080/api/load"
    var response = await fetch(url)
    var users = await response.json()
    this.setState({
      user: users[0]
    })
  }

  render() {
    return (
      <div className="container">
        <h1 id="title">Content Loader</h1>
        <button onClick={() => this.loadUsers()}>Load</button>
        {this.state.loading || !this.state.user 
        ? "" 
        : <User username={this.state.user.username} email={this.state.user.email} id={this.state.user.id}/>}
      </div>
    )
  }
}

export default App
