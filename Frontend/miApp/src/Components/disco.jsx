import disco from "../../assets/disco.png";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function DiskScreen() {
  const [data, setData] = useState([]) 
  const navigate = useNavigate()
  
  // execute the fetch command only once and when the component is loaded


  useState(() => {

    fetch(`http://localhost:3000/discos`,{
              
    method : 'GET',
    mode: "cors",
    headers:{
    'Content-Type': 'application/json'   
    }
    }).then(response => response.json()
    ).catch(err =>{
        console.error(err)
    }).then(diskData =>{
       setData(diskData)
    })

  }, [])

  const onClick = (obj) => {
    //e.preventDefault()
    
    navigate(`/disk/${obj}`)
    
  }

  return (
    <>
      
      
      <div style={{position: "relative",  marginLeft:280, border:"red 1px solid",display: "flex", flexDirection: "row"}}>

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
              onClick={() => onClick(objIterable.nombre)}
              >
                <img src={disco} alt="disk" style={{width: "100px"}} />
                <p>{objIterable.nombre}</p>
              </div>
            )
          })
        }
      
      </div>
    
     
    </>
   )
 }