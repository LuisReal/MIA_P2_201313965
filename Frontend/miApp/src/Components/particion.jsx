import partitionIMG from "../../assets/particion.png";
import { useState } from "react";
import { Link, useParams, useNavigate } from "react-router-dom";

export default function Partition() {
  const { id } = useParams()
  const [data, setData] = useState([])
  const navigate = useNavigate()

  // execute the fetch command only once and when the component is loaded
  useState(() => {
    // fetch('http://localhost:3001/api/commands') id
    //   .then(response => response.json())
    //   .then(rawData => {console.log(rawData);  setData(rawData.rutas);})
    const partitionData = {
      "particiones": ["Part1", "Part2", "Part3", "Part4", "Part5",]
    }
    setData(partitionData.particiones)

  }, [])

  const onClick = (objIterable) => {
    console.log("click", objIterable)
    navigate(`/login/${id}/${objIterable}`)
  }

  return (
    <>
      <p>Partition {id}</p>
      <br />
      <Link to="/">Home</Link>
      <br />
      <br />
      <br />
      <br />

      <div style={{position: "relative", marginLeft:280, height:500, border: "red 1px solid", display: "flex", flexDirection: "row" }}>

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
                <p1>{objIterable}</p1>
              </div>
            )
          })
        }

      </div>
    </>
  )
}