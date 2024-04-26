import React from 'react'
import { useState, useContext } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { UserContext } from "./usercontext";
import carpetaIMG from "../../assets/carpeta.png";
import archivoIMG from "../../assets/archivo.png";

function Sistema() {

  const { disk } = useParams()
  const [data, setData] = useState([])
  
  const navigate = useNavigate()
  const {setValue} = useContext(UserContext) //guarda el valor configurado en setValue en value(para ser usado en otro componente)

  let disco = disk.charAt(0);
  

  useState(() => {

    let obj = {
        
      'name': "Files"

    }

    fetch(`http://localhost:3000/archivo`,{
              
    method : 'POST',
    body: JSON.stringify(obj),
    headers:{
    'Content-Type': 'application/json'   
    }
    }).then(response => response.json()
    ).catch(err =>{
        console.error(err)
    }).then(res =>{
      
      setData(res)
      
    })


  }, [])

  const onClick = (objIterable) => {

    //setValue(objIterable.id) // guarda el valor del id
    
    //navigate(`/Login/${disco}/${objIterable.name}`)
  }

  return (
    

      <div style={{position: "relative", marginLeft:280, border: "red 1px solid", display: "flex", flexDirection: "row" }}>
         <p>Ruta</p>
          <br/>
        {
          data.map((objIterable, index) => {
            return (
            <>
              <div key={index} style={{
                border: "green 1px solid",
                display: "flex",
                flexDirection: "column", // Alinea los elementos en columnas
                alignItems: "center", // Centra verticalmente los elementos
                maxWidth: "100px",
              }}
                onClick={() => onClick(objIterable)}
              >
                {objIterable.tipo == "carpeta" &&
                <>
                  <img src={carpetaIMG} alt="carpeta" style={{ width: "100px" }} /> 
                  <p>{objIterable.name}</p>
                </> 
                
                }
                

                {objIterable.tipo == "archivo" && 
                <>
                  <img src={archivoIMG} alt="archivo" style={{ width: "100px" }} /> 
                  <p>{objIterable.name}</p>
                </> 
                
                }
                
                
              </div>

            </>
            )
          })
        }

      </div>

    
  )

}

export default Sistema