import React, { useEffect, useState } from 'react';

const List = (props) => {

const { patients } = props;

if (!patients || patients.length === 0) return <p>Aucun patient trouvé.</p>;
  return (
    <ul>
      {patients.map((patient) => {
        return (
            
          <div>

            <ul>
            <li><span className='lastname'><b>Nom de famille : </b>{patient.lastname}</span></li>
            <li><span className='firstname'><b>Prénom : </b>{patient.firstname}</span></li>
            <li><span className='gender'><b>Genre : </b>{patient.gender}</span></li>
            <li><span className='birthday'><b>Date de naissance : </b>{patient.birthday}</span></li>
            <li><span className='height'><b>Taille : </b>{patient.height}</span></li>
            <li><span className='weight'><b>Poids : </b>{patient.weight}</span></li>
            <li><span className='email'><b>Email : </b>{patient.email}</span></li>
            </ul>
            <br></br>
            </div> 
            );
      })}
    </ul>
  );
};
export default List;