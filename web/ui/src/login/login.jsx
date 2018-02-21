import React from 'react'

import "./login.css"

const LoginForm = () => {
  return (
    <div className="form-container">
      <form className="form-signin">
        <h1 className="h3 mb-3 font-weight-normal">Please sign in</h1>
        <label htmlFor="inputEmail" className="sr-only">Email address</label>
        <input type="email" id="inputEmail" className="form-control" placeholder="Email address" required autoFocus />
        <label htmlFor="inputPassword" className="sr-only">Password</label>
        <input type="password" id="inputPassword" className="form-control" placeholder="Password" required />
        <hr/>
        <button className="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
        <button className="btn btn-lg btn-primary btn-block" type="submit">Register</button>
      </form>
    </div>
  )
}

export default LoginForm
