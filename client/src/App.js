import React, { useEffect, useState } from 'react';
import './App.css';
import withListLoading from './components/withListLoading';
import Login from './components/Login';
import AddPatient from './components/AddPatient'
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Cookies from "js-cookie"
import List from './components/List';
import Addpatient from './components/AddPatient';

function setToken(userToken) {
  Cookies.set('token', userToken);
}

function getToken() {
  const userToken = Cookies.get('token');
  return userToken
}


function App() {
  const token = getToken();
  const ListLoading = withListLoading(List);
  const [appState, setAppState] = useState({
    loading: false,
    patients: null,
  });



  useEffect(() => {
    setAppState({ loading: true });
    const apiUrl = `http://localhost:10000/patients`;
    fetch(apiUrl)
      .then((res) => res.json())
      .then((patients) => {
        setAppState({ loading: false, patients: patients });
        console.log(patients)
      });
  }, [setAppState]);

    
  if(!token) {
    return <Login setToken={setToken} />
  }

  return (
    <div className='App'>
      <div className='container'>
        <h1>Liste des patients : </h1>
      </div>
      <div className='patients'>
        <ListLoading isLoading={appState.loading} patients={appState.patients} />
        <Switch>
          <Route path="/add">
            <AddPatient />
          </Route>
        </Switch>
      </div>
    </div>
  );
}

export default App;