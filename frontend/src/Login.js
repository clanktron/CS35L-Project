import React, { useState } from "react";

export const Login = (props) => {
    const [userName, setUserName] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        console.log("tbd");
    }

    return (
        <div className="form">
            <h2>ready to grind? log in.</h2>
            <form className="login-form" onSubmit={handleSubmit}>
            <input value={userName} onChange={(e) => setUserName(e.target.value)} placeholder="username" id="username" name="username" />
            <input value={password} onChange={(e) => setPassword(e.target.value)} type="password" placeholder="password" id="password" name="password" />
            <button type="submit">log in</button>
            </form>
            <button className='link-button' onClick={() => props.onFormSwitch('register')}>create account</button>
        </div>
    )
}