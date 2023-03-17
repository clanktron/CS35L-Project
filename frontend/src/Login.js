import React, { useState } from "react";

export const Login = (props) => {
  const [userName, setUserName] = useState('');
  const [password, setPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const errRef = useRef();

  const login = async (userName, password) => {
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
    .then((data) => console.log(data))
    .catch(err => { 
      if (!err?.response) {
        setErrMsg('No Server Response');
      } else if (err.response?.status === 400) {
        setErrMsg('Missing Username or Password');
      } else if (err.response?.status === 401) {
        setErrMsg('Unauthorized');
      } else {
        setErrMsg('Login Failed');
      }
      errRef.current.focus();
    })
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    login(userName, password);
  };

  return (
      <div className="form">
          <h2>ready to grind? log in.</h2>
          { ( error &&
          <h3 className="error">"log in failed</h3> )
          || ( props.onFormSwitch('mainpage') ) }
          <form className="login-form" onSubmit={handleSubmit}>
          <input value={userName} onChange={(e) => setUserName(e.target.value)} placeholder="username" id="username" name="username" required/>
          <input value={password} onChange={(e) => setPassword(e.target.value)} type="password" placeholder="password" id="password" name="password" required/>
          <button type="submit" onSubmit={handleSubmit}>log in</button>
          </form>
          <button className='link-button' onClick={() => props.onFormSwitch('register')}>create account</button>
      </div>
  )
}

/*
<error className="error" onSubmit={error ? <error value={error} /> : props.onFormSwitch('mainpage') }/>

          {onSubmit={ if (error) { <error value={error} /> } else: { props:onFormSwitch('mainpage') }}}
<button type="submit" onClick={() => props.onFormSwitch('mainpage')}>log in</button>

          .then(handleErrors)
    .then((data) => {
      console.log(data)
      props.onFormSwitch('mainpage')
    }, reason => {
      console.error(reason);
      setErrorMessage('login failed')
    });
      */