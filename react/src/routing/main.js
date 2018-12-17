import React, { Component } from 'react';
import LandingPage from '../screens/landing';
import MapPage from '../screens/map';
import ListPage from '../screens/list'
import { Switch, Route } from 'react-router-dom';


class MainRoute extends Component {
  render() {
    return (
      <Switch>
        <Route exact path='/' component={LandingPage}/>
        <Route path='/map' component={MapPage}/>
        <Route path='/list' component={ListPage}/>
      </Switch>
    );
  }
}

export default MainRoute;
