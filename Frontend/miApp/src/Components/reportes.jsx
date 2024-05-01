import React from 'react'
import { useState, useContext } from "react";
import svgIMG from "../../assets/svg.png";
import { GrafoContext } from "./usercontext";
import { useNavigate } from "react-router-dom";

function Reportes() {
  const navigate = useNavigate()
  const {setGrafo} = useContext(GrafoContext) 

  const [datos, setDatos] = useState([]) 
  
  useState(() => {

    fetch(`http://localhost:3000/getDot`,{
              
    method : 'GET',
    mode: "cors",
    headers:{
    'Content-Type': 'application/json'   
    }
    }).then(response => response.json()
    ).catch(err =>{
        console.error(err)
    }).then(reports =>{
       setDatos(reports)
    })

  }, [])

  const onClick = (objIterable) => {

    setGrafo(objIterable.contenido)//guarda el grafo(dot string)
    
    
    navigate(`/reporte/${objIterable.nombre}`)//llama al componente ImageReport
  }

  if(datos != null){
    return (
      <>  
          <div style={{width:1500, position:"relative", marginLeft:280, border:"1px solid blue", display: "flex", flexDirection: "row"}}>
            {
              datos.map((objIterable, index) => {
                return (
                
                  <div key={index} style={{
                    border: "green 1px solid",
                    display: "flex",
                    flexDirection: "column", // Alinea los elementos en columnas
                    alignItems: "center", // Centra verticalmente los elementos
                    maxWidth: "100px",
                  }}
                    onClick={() => onClick(objIterable)}
                  >
                    
                    <img src={svgIMG} alt="disk" style={{ width: "100px" }} />
                    <p>{objIterable.nombre}</p>
                    
                  </div>
  
                
                )
              })
            }
          </div>
      </>
    )
  }else{
    return(
      <p id="no-reportes">Todavia no hay reportes creados</p>
    )
    
  }
  
}

export default Reportes