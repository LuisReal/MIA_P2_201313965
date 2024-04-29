import React from 'react'
import { useState, useContext, useEffect } from "react";
import { useParams, useNavigate, useLocation } from "react-router-dom";
import { UserContext } from "./usercontext";
import carpetaIMG from "../../assets/carpeta.png";
import archivoIMG from "../../assets/archivo.png";


function Sistema() {

  const location = useLocation();

  //const [currentPath, setCurrentPath] = useState(location.pathname);
  const { disk, particion, archivo} = useParams()
  const [data, setData] = useState([])
  const [ruta, setRuta] = useState(
    {
      path: "/"
    }
  )

  
  const [obj, setObjeto] = useState(
    {
      name: archivo,
      
    }
  )

  if (archivo == "raiz") {
    obj.name = "/"
  }

  //console.log("El valor de archivo params(obj.name) es: ", obj.name)
  
  const navigate = useNavigate()
  const {setValue} = useContext(UserContext) //guarda el valor configurado en setValue en value(para ser usado en otro componente)

  let disco = disk.charAt(0);
  

  useEffect(() => {


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
      //console.log("La respuesta del useEffect es: ", res)
      setData(res)
      
    })


  }, [location]);

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

      var cadena = ""

      if(objIterable.tipo == "archivo"){
        
        for(let i = 0; i < res.length; i++) {
          cadena += res[i].contenido
        }

        setValue(cadena) //useContext

        navigate(`/contenido`)

      }else{
        setObjeto( //useState
          
          {
            name: objIterable.name
          }
        )

        setRuta(
          {
            path: ruta.path+objIterable.name
          }
        )

        navigate(`/disk/${disco}/${particion}/sistema/${objIterable.name}`)
      }
      
    })

  }

      return (
        <>
        {/*<button style={{marginLeft:290,}} onClick={() => navigate(-1)}>go back</button>*/}
        <p style={{marginLeft:290, padding:10, height:50, border: "black 2px solid", borderRadius:"20px"}}>Ruta</p>
        <br/>
        <div style={{position: "relative", marginLeft:280, display: "flex", flexDirection: "row" }}>
           
          {
            data.map((objIterable, index) => {
  
                
                return (
                  
                    <div key={index} style={{
                      //border: "green 1px solid",
                      display: "flex",
                      flexDirection: "column", // Alinea los elementos en columnas
                      alignItems: "center", // Centra verticalmente los elementos
                      maxWidth: "100px",
                      padding:"5px",
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
                    
                  
                  
                  
                  
                )
              
            })
          }
  
        </div>
        
        </>
      
      )
    
  
}

export default Sistema