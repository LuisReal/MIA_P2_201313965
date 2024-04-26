import React, { useState } from "react";
import "./App.css"
import Consola from "./Components/consola"
import Disco from "./Components/disco"
import Reportes from "./Components/reportes"
import Partition from "./Components/particion"
import Login from "./Components/login"
import Sistema from "./Components/sistema"
import Archivo from "./Components/archivo"
import { UserContext } from "./Components/usercontext"

import {
  HashRouter,
  Routes,
  Route,
  Link
 } from "react-router-dom";

function App() {
  
  const [value, setValue] = useState("")

  const logout = (e) => {
    e.preventDefault()

    let obj = {
        
      'Comand': "logout"
    }

    fetch(`http://localhost:3000/insert`,{
              
    method : 'POST',
    body: JSON.stringify(obj),
    headers:{
    'Content-Type': 'application/json'   
    }
    }).then(response => response.json()
    ).catch(err =>{
        console.error(err)
    }).then(res =>{
        
    })

  }

  return (
    
    <>
      <HashRouter>
      
      <div style={{position:"relative"}}>
        <h1 style={{backgroundColor:"rgb(249, 50, 50)", color:"white", textAlign:"center"}} >Sistema de Archivos</h1>
        <div className="d-flex flex-column flex-shrink-0 p-3" style={{width:280, position:"absolute", marginTop:-10, height:900, backgroundColor:"rgb(93, 173, 226)"}}>
          <ul className="nav nav-pills flex-column mb-auto" style={{color:"white", width:250 }}>
            <li className="nav-item">

            <Link to="/" id="enlace" style={{color: "white", textDecoration:"none", fontSize:"20px"}} aria-current="page">Pantalla1</Link>
            {/*<button onClick={pantalla1} id="b1" style={{width:250}} class="nav-link link-dark" aria-current="page">Pantalla1 </button>*/}
            </li>
            <hr/>
            <li>
            <Link to="/diskScreen" id="enlace" style={{color: "white", textDecoration:"none", fontSize:"20px"}} >Pantalla2</Link>
              {/*<button onClick={pantalla2} id="b2" style={{width:250}} class="nav-link link-dark">Pantalla2</button>*/}
            </li>
            <hr/>
            <li>
            <Link to="/reports" id="enlace" style={{color: "white", textDecoration:"none", fontSize:"20px"}} >Pantalla3</Link>
              {/*<button onClick={pantalla3} id="b3" style={{width:250}} class="nav-link link-dark">Pantalla3</button>*/}
            
            </li>
            
          </ul>

          <div>
              <button onClick={logout} id="btn-logout" style={{ width:100, marginLeft:10, marginBottom:50}} className="nav-link link-dark" aria-current="page">Logout </button>
          </div>
          

        </div>

        

        <UserContext.Provider value={{ value, setValue }}>{/*configura un value para ser usado en los siguientes componentes*/ }
            <Routes>
              <Route path="/" element={<Consola/>}></Route>
              <Route path="/diskScreen" element={<Disco/>}></Route>
              <Route path="/disk/:id" element={<Partition/>}></Route>
              <Route path="/Login/:disk/:particion" element={<Login/>}></Route>
              <Route path="/disk/:disk/:particion/sistema/:archivo" element={<Sistema/>}></Route>
              <Route path="/contenido" element={<Archivo/>}></Route>
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