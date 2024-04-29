import React from 'react'
import { useState, useContext } from "react";
import svgIMG from "../../assets/svg.png";
import { GrafoContext } from "./usercontext";
import { useNavigate } from "react-router-dom";

function Reportes() {
  const navigate = useNavigate()
  const {setGrafo} = useContext(GrafoContext) 

  const [data, setData] = useState([]) 
  
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
       setData(reports)
    })

  }, [])

  const onClick = (objIterable) => {

    setGrafo(objIterable.contenido)//guarda el grafo(dot string)
    
    
    navigate(`/reporte/${objIterable.nombre}`)
  }

  if(data != null){
    return (
      <>  
          <div style={{width:1500, position:"relative", marginLeft:280, border:"1px solid blue", display: "flex", flexDirection: "row"}}>
            {
              data.map((objIterable, index) => {
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