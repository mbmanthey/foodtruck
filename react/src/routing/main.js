import React, { Component } from 'react';
import LandingPage from '../screen/landing';
import MapPage from '../screen/map';
import { Switch, Route } from 'react-router-dom';


class MainRoute extends Component {
  render() {
    return (
      <Switch>
        <Route exact path='/' component={LandingPage}/>
        <Route path='/map' component={MapPage}/>
      </Switch>
    );
  }
}

export default MainRoute;
