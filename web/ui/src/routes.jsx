import React from 'react'
import {
  BrowserRouter as Router,
  Route
} from 'react-router-dom'
import App from './app'

const Routes = () => (
  <Router>
      <Route path="/" component={App}/>
  </Router>
)

export default Routes
