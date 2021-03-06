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

// check env
if (!process.env.REACT_APP_API_BASE_URL) {
  throw new Error("REACT_APP_API_BASE_URL not specified")
}

// api
const config: IAPIConfig = {
  baseURL: process.env.REACT_APP_API_BASE_URL
}
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
      <Router>
          <div className="content">
              {routes.map((item, index) => 
                <div>
                  <Link to={item.route}>{item.label}</Link>
                </div>
              )}

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
    </div>
  );
}

export default App;
