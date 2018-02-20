import React from 'react'
import {Link} from 'react-router-dom'

import "./header.css"

const Header = () => (
  <div className="header">
    <div className="header-section">
      <div className="header-item header-logo">
        <div className="">T</div>
      </div>
      <div className="header-item header-button is-site-header-item-selected">
        <Link to="/">Home</Link>
      </div>
    </div>
    <div className="header-section">
      <div className="header-item header-button">Settings</div>
      <div className="header-item header-button">Log out</div>
    </div>
  </div>
)

export default Header
