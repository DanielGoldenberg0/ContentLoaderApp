import React from 'react'

class User extends React.Component {
  render() {
    return (
      <div className="user-box">
        <div style={{marginBottom: 10}}>
        <h2>{this.props.username}</h2>
        </div>
        <h3>{this.props.email}</h3>
        <h3>{this.props.id}</h3>
      </div>
    )
  }
}

export default User
