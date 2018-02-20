import React from 'react'
import {Link} from 'react-router-dom'

import "./header.css"

const Header = () => (
  <div className="header">
    <div className="header-section">
      <div className="header-item header-logo">
        <div className="">T</div>
      </div>
      <Link className="header-item header-button is-site-header-item-selected" to="/">Home</Link>
    </div>
    <div className="header-section">
      <Link className="header-item header-button" to="/settings">Settings</Link>
      <Link className="header-item header-button" to="/logout">Logout</Link>
    </div>
  </div>
)

export default Header
