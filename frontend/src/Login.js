import React, { useRef, useState, useEffect } from "react";

export const Login = (props) => {

  const userRef = useRef();
  const errRef = useRef();

  const [userName, setUserName] = useState('');

  const [password, setPassword] = useState('');

  const [errMsg, setErrMsg] = useState('');
  const [success, setSuccess] = useState(false);
  //const errorMessageStyle = document.getElementById("error-msg");

  useEffect(() => {
    userRef.current.focus();
  }, [])

  useEffect(() => {
    setErrMsg('');
  }, [userName, password])

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
        // console.log(text);
        if (text !== '') {
            setSuccess(true);
        } else {
            setSuccess(false);
            setErrMsg('incorrect username or password');
            errRef.current.focus();
        }
    });
};

  const handleSubmit = async (e) => {
    e.preventDefault();
    /*response(userName, password);
    if (errorMessage) {
      errorMessageStyle.style.opacity = 1;
    }
    else {
      errorMessageStyle.style.opacity = 0;
      props.onFormSwitch('mainpage');
    }*/
    try {
      response(userName, password);
    } catch (err) {
      if (!err?.response) {
        setErrMsg('No Server Response');
      } else {
        setErrMsg('Log in Failed')
      }
      errRef.current.focus();
    }
  }

  return (
    <>
        {success ? (props.onFormSwitch('mainpage')
        ) : (
          <div className="form">
                <h2>ready to grind? log in.</h2>
                <form className="login-form" onSubmit={handleSubmit}>
                <input value={userName} onChange={(e) => setUserName(e.target.value)} placeholder="username" 
                id="username" name="username" ref={userRef} autoComplete="off"/>
                <input value={password} onChange={(e) => setPassword(e.target.value)} type="password" placeholder="password" 
                id="password" name="password"/>
                <button type="submit">log in</button>
                </form>
                <button className='link-button' onClick={() => props.onFormSwitch('register')}>create account</button>
                <p ref={errRef} className={errMsg ? "err-msg" : "offscreen"}>{errMsg}</p>
                </div>
        )}
    </>
)
}

/*
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
*/