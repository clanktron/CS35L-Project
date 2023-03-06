import React, {useState} from 'react';

export const Register = (props) => {

    const [userName, setUserName] = useState('');
    const [password, setPassword] = useState('');
    const [cPassword, setCPassword] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        if (password !== cPassword) {
            console.log("password not the same");
        } else {
            console.log("success");
        }
    }

    
    return (
        <div className='form'>
            <form className='register-form' onSubmit={handleSubmit}>
                <label for="username">Username</label>
                <input value={userName} onChange={(e) => setUserName(e.target.value)} placeholder="Username" id="username" name="username" />
                <label for="password">Password</label>
                <input value={password} onChange={(e) => setPassword(e.target.value)} placeholder="Password" id="password" name="password" />
                <label for="cPassword">Confirm Password</label>
                <input value={cPassword} onChange={(e) => setCPassword(e.target.value)} placeholder="Confirm Password" id="cPassword" name="cPassword" />
                <button type='submit'>Create</button>
            </form>
            <button className='link-button' onClick={() => props.onFormSwitch('login')}>Back to Login</button>        
        </div>
    )
}