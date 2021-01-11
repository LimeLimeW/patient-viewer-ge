import React, { useState } from 'react';
import './Login.css';
import PropTypes from 'prop-types';

export default function AddPatient({ setPatient }) {
    const [firstname, setFirstname] = useState();
    const [lastname, setLastname] = useState();
    const [birthday, setBirthday] = useState();
    const [height, setHeight] = useState();
    const [weight, setWeight] = useState();
    const [email, setEmail] = useState();
    const [gender, setGender] = useState();

    const handleSubmit = async e => {
        e.preventDefault();
        const data = await addPatient({
            firstname,
            lastname,
            birthday,
            height,
            weight,
            email,
            gender
        });
        setPatient(data);
      }

    async function addPatient(patient) {
        return fetch('http://localhost:10000/patient', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(patient)
        }).then(data => console.log(data))
          .then(data => data.json())
       }

  return(
    <div className="login-wrapper">
      <h1>Ajout de patient :</h1>
      <form onSubmit={handleSubmit}>
        <label>
          <p>Prénom</p>
          <input type="text" onChange={e => setFirstname(e.target.value)}/>
        </label>
        <label>
        <p>Nom de famille</p>
          <input type="text" onChange={e => setLastname(e.target.value)}/>
        </label>
        <label>
          <p>Date de naissance</p>
          <input type="text" onChange={e => setBirthday(e.target.value)}/>
        </label>
        <label>
          <p>Taille</p>
          <input type="text" onChange={e => setHeight(e.target.value)}/>
        </label>
        <label>
        <p>Poids</p>
          <input type="text" onChange={e => setWeight(e.target.value)}/>
        </label>
        <label>
        <p>Email</p>
          <input type="text" onChange={e => setEmail(e.target.value)}/>
        </label>
        <label>
        <p>Genre</p>
          <input type="text" onChange={e => setGender(e.target.value)}/>
        </label>
        <div>
          <button type="submit">Ajouter</button>
        </div>
      </form>
    </div>
  )
}

