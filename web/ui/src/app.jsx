import React from 'react'
import {Switch, Route} from 'react-router-dom'
import Header from 'header/header'
import Footer from 'footer/footer'
import LoginForm from "./login/login";

import 'app.css'

// Comprises of the outer shell of the application - Header, Footer, and containing body
const App = () => (
  <div className="with-sticky-footer">
    <Header/>
    <Switch>
      <Route exact path='/' component={LoginForm}/>
    </Switch>
    <Footer/>
  </div>
)

export default App
