import {
  Routes,
  Route,
} from "react-router-dom";
import './App.css';
import Sidenav from './Components/Sidenav';
import Todos from "./Pages/Todos";
import Home from "./Pages/Home";
import Settings from "./Pages/Settings";
import Archive from "./Pages/Archive";

function App() {
  return (
    <div className="App">
      <Sidenav/>
      <main>
      <Routes>
        <Route  path="/" element={<Home />}/>
        <Route path="/todos" element={<Todos />} />
        <Route path="/archive" element={<Archive />}/>
        <Route path="/settings" element={<Settings />} />
      </Routes>
      </main>
     
    </div>
  );
}

export default App;
