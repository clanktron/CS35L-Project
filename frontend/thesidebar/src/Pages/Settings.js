import React, { useState, useEffect } from 'react'

export default function Settings() {
  const [username, setUsername] = useState([])
  const [password, setPassword] = useState([])

  const login = async (username, password) => {                 //fetch data by login credentials
    await fetch('http://localhost:4000/login', {
       method: 'POST',
       mode: 'cors',
       credentials: 'include',
       headers: {
          'Content-type': 'application/json; charset=UTF-8',
       },
       body: JSON.stringify({
          username: username,
          password: password,
       }),
    })
       .then((data) => console.log(data));
 };
 
 const handleSubmit = (e) => {
    e.preventDefault();
    login(username, password);
 }; 
  
  useEffect(() => {
    fetch('http://localhost:4000') //fetch data from API, need help on this
      .then(response => response.json()) // parse into json file
      .then(response => console.log(response))
      .then(data => password(data))
      .catch(error => console.error(error)) //If there is an error, console.error function, which will log the error message to the console.
  }, [])

  return (
    <ul>
      {username.map(item => <li key={item.id}>{item.name}</li>)}     
    </ul>
  ) //unordered list (<ul>) with each item represented by a list item
}
//list output