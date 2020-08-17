import React from 'react'
import User from "./components/User.js"

var users
var clicked

class App extends React.Component {
  state = {
    loading: true,
    user: null
  }

  async loadUsers() {
    var url = "http://localhost:8080/api/load"
    var response = await fetch(url)

    users = await response.json()
    clicked = true;

    this.setState({
      loading: false,
      user: users[0]
    })
  }

  render() {
    return (
      <div className="container">
        <h1 id="title">Content Loader</h1>
        <button onClick={() => this.loadUsers()}>Load</button>
        {
          clicked
          ?
            !this.state.loading && this.state.user
            ? 
              users.map(user => <User username={user["username"]} email={user["email"]} id={user["id"]}/>)
            : 
              <h1>Error</h1>
          :
          ""
        }
      </div>
    )
  }
}

export default App
