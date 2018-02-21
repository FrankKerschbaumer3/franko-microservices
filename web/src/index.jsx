import React from 'react'
import ReactDOM from 'react-dom'
import Routes from './routes'

import "normalize.css"

ReactDOM.render(<Routes />, document.getElementById('app'))

// Will be removed in PROD bundling
if (module.hot) {
  module.hot.accept()
}