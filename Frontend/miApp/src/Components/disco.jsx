import disco from "../../assets/disco.png";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

export default function DiskScreen() {
  const [data, setData] = useState([]) 
  const navigate = useNavigate()
  
  // execute the fetch command only once and when the component is loaded
  useState(() => {
    // fetch('http://localhost:3001/api/commands')
    //   .then(response => response.json())
    //   .then(rawData => {console.log(rawData);  setData(rawData.rutas);})
    const diskData = {
      "rutas":["A.dsk","B.dsk","C.dsk","D.dsk","A.dsk","B.dsk","C.dsk","D.dsk","A.dsk","B.dsk","C.dsk","D.dsk"]
    }
    setData(diskData.rutas)

  }, [])

  const onClick = (obj) => {
    //e.preventDefault()
    console.log("click",obj)
    navigate(`/disk/${obj}`)
    {particiones}
  }

  return (
    <>
      
      <p>Pantalla Discos</p>
      <div style={{position: "relative",  marginLeft:280, height:500, border:"red 1px solid",display: "flex", flexDirection: "row"}}>
      
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
                <img src={disco} alt="disk" style={{width: "100px"}} />
                <p1>{objIterable}</p1>
              </div>
            )
          })
        }
      
      </div>
    </>
   )
 }