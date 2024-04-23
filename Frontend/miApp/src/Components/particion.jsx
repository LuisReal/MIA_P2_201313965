import partitionIMG from "../../assets/particion.png";
import { useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import Login from "./login"

export default function Partition() {
  const { id } = useParams()
  const [data, setData] = useState([])
  const navigate = useNavigate()

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
    
    navigate(`/Login/${disco}/${objIterable.name}`)
  }

  return (
    <>

      <div style={{position: "relative", marginLeft:280, border: "red 1px solid", display: "flex", flexDirection: "row" }}>
        
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
                
                <img src={partitionIMG} alt="disk" style={{ width: "100px" }} />
                <p>{objIterable.name}</p>
              </div>
            )
          })
        }

      </div>
    </>
  )
}