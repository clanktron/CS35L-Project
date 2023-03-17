import React, { useState } from "react";

export const Login = (props) => {
  const [userName, setUserName] = useState('');
  const [password, setPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState();
  const errorMessageStyle = document.getElementById("error-msg");

  const response = async (userName, password) => {
    await fetch('http://localhost:4000/login', {
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-type': 'application/json; charset=UTF-8',
      },
      body: JSON.stringify({
        username: userName,
        password: password,
      }),
    })
    .then((response) => response.text())
    .then((text) => {
      if (text === '') {
        setErrorMessage(true);
      }
      else {
        setErrorMessage(false);
      }
    })
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    response(userName, password);
    if (errorMessage) {
      errorMessageStyle.style.opacity = 1;
    }
    else {
      errorMessageStyle.style.opacity = 0;
      props.onFormSwitch('mainpage');
    }
  };

  return (
      <div className="form">
          <h2>ready to grind? log in.</h2>
          <form className="login-form" onSubmit={handleSubmit}>
          <input value={userName} onChange={(e) => setUserName(e.target.value)} placeholder="username" id="username" name="username" required/>
          <input value={password} onChange={(e) => setPassword(e.target.value)} type="password" placeholder="password" id="password" name="password" required/>
          <button type="submit" onSubmit={handleSubmit}>log in</button>
          </form>
          <button className='link-button' onClick={() => props.onFormSwitch('register')}>create account</button>
          <p id="error-msg">incorrect username or password</p>
      </div>
  )
}