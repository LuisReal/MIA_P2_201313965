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
      <h1 style={{marginLeft:280}} >Sistema de Archivos</h1>
      <div class="d-flex flex-column flex-shrink-0 p-3 bg-light" style={{width:280, marginTop:-57}}>
        <ul class="nav nav-pills flex-column mb-auto">
          <li class="nav-item">
            <a href="#" class="nav-link active" aria-current="page">
              <svg class="bi me-2" width="16" height="16"><use xlink:href="#home"/></svg>
              Pantalla1
            </a>
          </li>
          <li>
            <a href="#" class="nav-link link-dark">
              <svg class="bi me-2" width="16" height="16"><use xlink:href="#speedometer2"/></svg>
              Pantalla2
            </a>
          </li>
          <li>
            <a href="#" class="nav-link link-dark">
              <svg class="bi me-2" width="16" height="16"><use xlink:href="#table"/></svg>
              Pantalla3
            </a>
          </li>
        </ul>

      </div>

      <Router>
        <Routes>
          <Route path="/consola" element={<Consola/>}></Route>
        </Routes>
      </Router>
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