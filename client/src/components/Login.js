import React, { useState } from 'react';
import './Login.css';
import PropTypes from 'prop-types';

export default function Login({ setToken }) {
    const [username, setUserName] = useState();
    const [password, setPassword] = useState();

    const handleSubmit = async e => {
        e.preventDefault();
        const token = await loginUser({
          username,
          password
        });
        setToken(token);
      }

    async function loginUser(credentials) {
        return fetch('http://localhost:10000/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          credentials: 'include',
          body: JSON.stringify(credentials)
        })
          .then(data => data.json())
       }

  return(
    <div className="login-wrapper">
      <h1>Veuillez vous connecter.</h1>
      <form onSubmit={handleSubmit}>
        <label>
          <p>Nom d'utilisateur</p>
          <input type="text" onChange={e => setUserName(e.target.value)}/>
        </label>
        <label>
          <p>Mot de passe</p>
          <input type="password" onChange={e => setPassword(e.target.value)}/>
        </label>
        <div>
          <button type="submit">Submit</button>
        </div>
      </form>
    </div>
  )
}

Login.propTypes = {
    setToken: PropTypes.func.isRequired
  }
