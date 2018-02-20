import React from 'react'
import Header from 'header/header'
import Footer from 'footer/footer'

import 'app.css'

// Comprises of the outer shell of the application - Header, Footer, and containing body
const App = (props) => (
  <div className="with-sticky-footer">
    <Header/>
    {props.children}
    <Footer/>
  </div>
)

export default App
