import React from 'react';
import './App.css';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import  {PageUserInfo} from './pages/UserInfo'
import  {PageCreateAccount} from  './pages/CreateAccount'
import { UserStore } from './stores/user';
import { UserInfoAPI } from './api/user';
import { IAPIConfig } from './api';
import { IUser } from './stores';
import { CreateAccountStore } from './stores/account';
import { AccountCreateAPI } from './api/account';

const user: IUser = {
  id: 1,
}

// api
const config: IAPIConfig = {baseURL: 'http://127.0.0.1:8000'}
const userInfoAPI = new UserInfoAPI(config)
const accountCreateAPI = new AccountCreateAPI(config)
// --

// stores
const userStore = new UserStore(userInfoAPI)
const accountStore = new CreateAccountStore(accountCreateAPI)
// --

const routes = [
  {
    route: '/userinfo',
    label: 'User Info',
    page: <PageUserInfo store={userStore} user={user} />,
  },
  {
    route: '/createaccount',
    label: 'Create Account',
    page: <PageCreateAccount store={accountStore} user={user} />,
  },
]

function App() {
  return (
    <div className="App">
      <header className="App-header">
      <Router>
          <div>
            <ul>
              {routes.map((item, index) => 
                <li key={index}>
                  <Link to={item.route}>{item.label}</Link>
              </li>
              )}
            </ul>

            <hr />

            {/*
              A <Switch> looks through all its children <Route>
              elements and renders the first one whose path
              matches the current URL. Use a <Switch> any time
              you have multiple routes, but you want only one
              of them to render at a time
            */}
            <Switch>
              {routes.map((item, index) => 
                <Route key={index} exact path={item.route}>
                  {item.page}
                </Route>
              )}
            </Switch>
          </div>
        </Router>
      </header>
    </div>
  );
}

export default App;
