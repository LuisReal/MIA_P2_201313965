import React, { useState } from "react";
import "./App.css"
import Consola from "./Components/consola"
import Disco from "./Components/disco"
import Reportes from "./Components/reportes"
import Partition from "./Components/particion"
import Login from "./Components/login"
import { UserContext } from "./Components/usercontext"

import {
  HashRouter,
  Routes,
  Route,
  Link
 } from "react-router-dom";

function App() {
  
  const [value, setValue] = useState("")

  return (
    
    <>
      <HashRouter>
      
      <div style={{position:"relative"}}>
        <h1 style={{backgroundColor:"rgb(249, 50, 50)", color:"white", textAlign:"center"}} >Sistema de Archivos</h1>
        <div className="d-flex flex-column flex-shrink-0 p-3" style={{width:280, position:"absolute", marginTop:-10, height:900, backgroundColor:"rgb(255, 255, 224)"}}>
          <ul className="nav nav-pills flex-column mb-auto">
            <li className="nav-item">

            <Link to="/" className="nav-link link-dark" aria-current="page">Pantalla1</Link>
            {/*<button onClick={pantalla1} id="b1" style={{width:250}} class="nav-link link-dark" aria-current="page">Pantalla1 </button>*/}
            </li>
            <hr/>
            <li>
            <Link to="/diskScreen" className="nav-link link-dark">Pantalla2</Link>
              {/*<button onClick={pantalla2} id="b2" style={{width:250}} class="nav-link link-dark">Pantalla2</button>*/}
            </li>
            <hr/>
            <li>
            <Link to="/reports" className="nav-link link-dark">Pantalla3</Link>
              {/*<button onClick={pantalla3} id="b3" style={{width:250}} class="nav-link link-dark">Pantalla3</button>*/}
            
            </li>
          </ul>

        </div>

        <UserContext.Provider value={{ value, setValue }}>
            <Routes>
              <Route path="/" element={<Consola/>}></Route>
              <Route path="/diskScreen" element={<Disco/>}></Route>
              <Route path="/disk/:id" element={<Partition/>}></Route>
              <Route path="/Login/:disk/:particion" element={<Login/>}></Route>
              <Route path="/reports" element={<Reportes/>}></Route>
            </Routes>
          
        </UserContext.Provider>
      </div>

      </HashRouter>
    </>
    
  )
}

export default App
/*

<Switch>
          <Route path="/consola"  exact > 
            <Consola/>
          </Route>
           
         
          <Route path="/sistema" >
            <Sistema/>
          </Route>
           
          
          <Route path="/reportes">
            <Reportes/>
          </Route>
            
        
        </Switch>
  
*/