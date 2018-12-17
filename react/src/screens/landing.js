import React, { Component } from 'react';
import { Link, Switch } from 'react-router-dom';
class LandingPage extends Component {

  render() {
    return (
      <Switch>
        <div>
          <Link to="/list"><button>List</button></Link>
          <Link to='/map'><button>Map</button></Link>
        </div>
      </Switch>
        
    )
  }
}

export default LandingPage;
