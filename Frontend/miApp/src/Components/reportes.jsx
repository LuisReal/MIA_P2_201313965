import React from 'react'
import { useState } from "react";
import svgIMG from "../../assets/svg.png";
import {Graphviz} from "graphviz-react"

function Reportes() {
  
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
    
    navigate(`/reporte/info`)
  }

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
                  {/*<p>{objIterable.name}</p>*/}
                  
                </div>

              
              )
            })
          }
        </div>
    </>
  )
}

export default Reportes