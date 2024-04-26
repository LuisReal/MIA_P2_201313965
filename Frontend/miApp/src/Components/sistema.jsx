import React from 'react'
import { useState, useContext } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { UserContext } from "./usercontext";
import carpetaIMG from "../../assets/carpeta.png";
import archivoIMG from "../../assets/archivo.png";

function Sistema() {

  
  const { disk, particion  } = useParams()
  const [data, setData] = useState([])


  const [obj, setObjeto] = useState(
    {
      name: '/',
      
    }
  )
  
  const navigate = useNavigate()
  const {setValue} = useContext(UserContext) //guarda el valor configurado en setValue en value(para ser usado en otro componente)

  let disco = disk.charAt(0);
  

  useState(() => {

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

    //alert(objIterable.name)
    const obj_ = {
      'name': objIterable.name
    }

    fetch(`http://localhost:3000/archivo`,{
              
    method : 'POST',
    body: JSON.stringify(obj_),
    headers:{
    'Content-Type': 'application/json'   
    }
    }).then(response => response.json()
    ).catch(err =>{
        console.error(err)
    }).then(res =>{
      
      setData(res)
      
    })
    //setValue(objIterable.id) // guarda el valor del id
    //console.log("obj.name es: ", obj.name)
    navigate(`/disk/${disco}/${particion}/sistema/${objIterable.name}`)

  
  }

  return (
    

      <div style={{position: "relative", marginLeft:280, border: "red 1px solid", display: "flex", flexDirection: "row" }}>
         
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
               {/*} {objIterable.tipo == "carpeta" && (objIterable.name != "." || objIterable.name != "..") &&
               
                
                }*/}

                {(() => {
                  if (objIterable.tipo == "carpeta" && (objIterable.name != "."&& objIterable.name != "..")){
                      return (
                        <>
                          <img src={carpetaIMG} alt="carpeta" style={{ width: "100px" }} /> 
                          <p>{objIterable.name}</p>
                        </> 
                      )
                  }
                  
                  return null;
                })()}
                
                {(() => {
                  if (objIterable.tipo == "archivo" && (objIterable.name != "." && objIterable.name != "..")){
                      return (
                        <>
                          <img src={archivoIMG} alt="archivo" style={{ width: "100px" }} /> 
                          <p>{objIterable.name}</p>
                        </> 
                      )
                  }
                  
                  return null;
                })()}

                {/*{objIterable.tipo == "archivo" && (objIterable.name != "." || objIterable.name != "..") &&
                <>
                  <img src={archivoIMG} alt="archivo" style={{ width: "100px" }} /> 
                  <p>{objIterable.name}</p>
                </> 
                
                }*/}
                
                
              </div>

            </>
            )
          })
        }

      </div>

    
  )

}

export default Sistema