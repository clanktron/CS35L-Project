import React, { useRef, useState, useEffect, useContext } from "react";
import AuthContext from "./context/AuthProvider";
import axios from "./api/axios"

export const Login = (props) => {
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

    return (
        <div className="form">
            <h2>ready to grind? log in.</h2>
            <div> {!success && <>ref={errRef} className={errMsg ? "errmsg" : "offscreen"} aria-live="assertive" {errMsg}</>}</div>
            <form className="login-form" onSubmit={handleSubmit}>
            <input value={userName} onChange={(e) => setUserName(e.target.value)} placeholder="username" id="username" name="username" ref={userRef}/>
            <input value={password} onChange={(e) => setPassword(e.target.value)} type="password" placeholder="password" id="password" name="password" />
            <button type="submit" onClick={() => props.onFormSwitch('mainpage')}>log in</button>
            </form>
            <button className='link-button' onClick={() => props.onFormSwitch('register')}>create account</button>
        </div>
    )
}