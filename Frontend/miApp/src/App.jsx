import React from "react";

import Consola from "./Components/consola"
import Sistema from "./Components/sistema"
import Reportes from "./Components/reportes"

import {
  BrowserRouter as Router,
  Routes,
  Route
 } from "react-router-dom";

function App() {

  return (
    
    <>
      <div style={{position:"relative"}}>
        <h1 style={{backgroundColor:"rgb(249, 50, 50)", color:"white", textAlign:"center"}} >Sistema de Archivos</h1>
        <div class="d-flex flex-column flex-shrink-0 p-3" style={{width:280, position:"absolute", height:900, backgroundColor:"rgb(255, 255, 224)"}}>
          <ul class="nav nav-pills flex-column mb-auto">
            <li class="nav-item">
              <a href="#" class="nav-link active" aria-current="page">
                <svg class="bi me-2" width="16" height="16"><use xlink:href="#home"/></svg>
                Pantalla1
              </a>
            </li>
            <hr/>
            <li>
              <a href="#" class="nav-link link-dark">
                <svg class="bi me-2" width="16" height="16"><use xlink:href="#speedometer2"/></svg>
                Pantalla2
              </a>
            </li>
            <hr/>
            <li>
              <a href="#" class="nav-link link-dark">
                <svg class="bi me-2" width="16" height="16"><use xlink:href="#table"/></svg>
                Pantalla3
              </a>
            </li>
          </ul>

        </div>

        <div style={{position:"relative", marginLeft:280, border:"1px solid blue", height:500}}>
        <Router>
          <Routes>
            <Route path="/consola" element={<Consola/>}></Route>
          </Routes>
        </Router>
        </div>
      </div>
      
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