import partitionIMG from "../../assets/particion.png";
import { useState, useContext } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { UserContext } from "./usercontext";


export default function Partition() {
  const { id } = useParams()
  const [data, setData] = useState([])
  
  const navigate = useNavigate()
  const {setValue} = useContext(UserContext) //guarda el valor configurado en setValue en value(para ser usado en otro componente)

  let disco = id.charAt(0);
  

  useState(() => {

    fetch(`http://localhost:3000/disk/${disco}`,{
              
    method : 'GET',
    mode: "cors",
    headers:{
    'Content-Type': 'application/json'   
    }
    }).then(response => response.json()
        ).catch(err =>{
            console.error(err)
        }).then(partitionData =>{
            setData(partitionData.particiones)
        })


  }, [])

  const onClick = (objIterable) => {

    setValue(objIterable.id) // guarda el valor del id
    
    navigate(`/Login/${disco}/${objIterable.name}`)
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
                
                <img src={partitionIMG} alt="disk" style={{ width: "100px" }} />
                <p>{objIterable.name}</p>
                
              </div>

            </>
            )
          })
        }

      </div>
      
      
    
  )
}