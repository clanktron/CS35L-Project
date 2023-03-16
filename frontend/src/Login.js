import React, { useState } from "react";

export const Login = (props) => {
  const [userName, setUserName] = useState('');
  const [password, setPassword] = useState('');

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
    .then((data) => console.log(data));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    login(userName, password);
  };

  return (
      <div className="form">
          <h2>ready to grind? log in.</h2>
          <form className="login-form" onSubmit={handleSubmit}>
          <input value={userName} onChange={(e) => setUserName(e.target.value)} placeholder="username" id="username" name="username"/>
          <input value={password} onChange={(e) => setPassword(e.target.value)} type="password" placeholder="password" id="password" name="password" />
          <button type="submit" onClick={() => props.onFormSwitch('mainpage')}>log in</button>
          </form>
          <button className='link-button' onClick={() => props.onFormSwitch('register')}>create account</button>
      </div>
  )
}

/*
const { setAuth } = useContext(AuthContext);
    const userRef = useRef();
    const errRef = useRef();
    const [userName, setUserName] = useState('');
    const [password, setPassword] = useState('');
    const [errMsg, setErrMsg] = useState("");
    const [success, setSuccess] = useState(false);
    useEffect(() => {
      userRef.current.focus();
    }, []);
    useEffect(() => {
      setErrMsg("");
    }, [userName, password]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
          const response = await axios.post(
            'http://localhost:4000/login',
            JSON.stringify({ userName, password }),
            {
              headers: {
                cookie: 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzgxOTEwMTMsInVzZXJpZCI6Ijg0NTc5NzgyODg4MDU5Njk5MyJ9.ib23dJ-KYCps9Bw589bA-ExjkouH1EmPfUCR5_KKiw8',
                "Content-Type": "application/json"
            },
            data: {Username: 'admin', Password: 'admin'},
              withCredentials: true,
            }
          );
          const accessToken = response?.data?.accessToken;
          const roles = response?.data?.roles;
          setAuth({ userName, password, roles, accessToken });
          setUserName("");
          setPassword("");
          setSuccess(true);
        } catch (err) {
          if (!err?.response) {
            setErrMsg("No Server Response");
          } else if (err.response?.status === 400) {
            setErrMsg("Missing Username or Password");
          } else if (err.response?.status === 401) {
            setErrMsg("Unauthorized");
          } else {
            setErrMsg("Login Failed");
          }
          errRef.current.focus();
        }
      };
      */