import React from "react";
import "./App.css"
import Consola from "./Components/consola"
import Disco from "./Components/disco"
import Reportes from "./Components/reportes"
import Partition from "./Components/particion"

import {
  HashRouter,
  Routes,
  Route,
  Link
 } from "react-router-dom";

function App() {

  let consola = "block";
  let sistema = "";
  let reportes ="";

  const pantalla1 = () => {
    console.log("pantalla 1")

    //document.getElementById("b1").style.backgroundColor = "light-blue"

    sistema = document.getElementById("sistema").style.display
    reportes = document.getElementById("reportes").style.display
    particion = document.getElementById("particion").style.display

    
    if (reportes === "block" || sistema ==="block" || particion ==="block"){

      document.getElementById("consola").style.display ="block"
      document.getElementById("reportes").style.display ="none"
      document.getElementById("sistema").style.display ="none"
      document.getElementById("particion").style.display ="none"
      
      
    }
  }

  const pantalla2 = () => {
    console.log("pantalla 2")
    consola = document.getElementById("consola").style.display
    reportes = document.getElementById("reportes").style.display

    if (consola === "block" || reportes ==="block"){

      document.getElementById("consola").style.display ="none"
      document.getElementById("reportes").style.display ="none"
      document.getElementById("sistema").style.display ="block"
      
      
    }
    
  }

  const pantalla3 = () => {
    console.log("pantalla 3")
    
    consola = document.getElementById("consola").style.display
    sistema = document.getElementById("sistema").style.display

    if (consola === "block" || sistema ==="block"){

      document.getElementById("consola").style.display ="none"
      document.getElementById("reportes").style.display ="block"
      document.getElementById("sistema").style.display ="none"
      
      
    }
    
  }

  const particiones = () => {
    console.log("pantalla Particiones")
    
    consola = document.getElementById("consola").style.display
    sistema = document.getElementById("sistema").style.display
    reportes = document.getElementById("reportes").style.display

    if (consola === "block" || sistema ==="block" || reportes ==="block"){

      document.getElementById("consola").style.display ="none"
      document.getElementById("reportes").style.display ="none"
      document.getElementById("sistema").style.display ="none"
      
      document.getElementById("partition").style.display ="block"
    }
    
  }

  

  return (
    
    <>
      <HashRouter>
      
      <div style={{position:"relative"}}>
        <h1 style={{backgroundColor:"rgb(249, 50, 50)", color:"white", textAlign:"center"}} >Sistema de Archivos</h1>
        <div class="d-flex flex-column flex-shrink-0 p-3" style={{width:280, position:"absolute", height:900, backgroundColor:"rgb(255, 255, 224)"}}>
          <ul class="nav nav-pills flex-column mb-auto">
            <li class="nav-item">

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

          
            <Routes>
              <Route path="/" element={<Consola/>}></Route>
            </Routes>
        
            <Routes>
              <Route path="/diskScreen" element={<Disco/>}></Route>
            </Routes>
          
            <Routes>
              <Route path="/disk/:id" element={<Partition/>}></Route>
            </Routes>
          
            <Routes>
              <Route path="/reports" element={<Reportes/>}></Route>
            </Routes>
          
        
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