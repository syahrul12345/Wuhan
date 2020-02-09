
import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";
import Home from './pages/Home'
import Create from './pages/Create'
import Login from './pages/Login'
import Play from './pages/Play';

export default function App() {
  return(
    <Router>
      <Switch>
        <Route exact path="/">
          <Home />
        </Route>
        <Route exact path="/create">
          <Create />
        </Route>
        <Route exact path="/login">
          <Login />
        </Route>
        <Route exact path="/play">
          <Play />
        </Route>
      </Switch>
    </Router>
  )
}
