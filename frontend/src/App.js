import './App.css';
import React, {useState} from 'react';
import {Login} from "./Login";
import {Register} from "./Register";
import {Mainpage} from "./Mainpage";


function App() {
  const [currentForm, setCurrentForm] = useState('login');

  const toggleForm = (formName) => {
    setCurrentForm(formName);
  }

  return (
    <div className="App">
      {
      currentForm === 'login' ? <Login onFormSwitch={toggleForm}/> : ( currentForm === 'register' ? <Register onFormSwitch={toggleForm}/> : <Mainpage onFormSwitch={toggleForm}/>)
      }
    </div>
  );
}

export default App;