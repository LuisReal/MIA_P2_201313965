
import {useState} from 'react'
import React from 'react'

function Consola() {

    const [datos, setDatos] = useState(
        {
            comando: ''
        }
    )

    const getDatos = (event) => {
        event.preventDefault();
        setDatos({
            ...datos,
            [event.target.name] : event.target.value
        })
        
       //console.log(datos.comando)
    }

    const mostrarDatos = (event) => {
        event.preventDefault();
        console.log(datos.comando)

        /*
        fetch('http://localhost:3000/disk')
        .then(response => response.json())
        .then(partitionData => {setData(partitionData.particiones);})*/
    }


    return (
        <>
            <div style={{position:"relative",  marginLeft:280, border:"1px solid blue", height:500}}>

                <form action="" onSubmit={mostrarDatos}>
                    <p name="consola" style={{height:450, width:1580, position:"absolute"}} >{datos.comando}</p>
                    <input type="text" onChange={getDatos} name="comando" style={{height:50, width:1480, position:"absolute", marginTop:450}} placeholder="Ingrese comando" />
                    <button type="submit" className='btn btn-primary' style={{height:50, position:"absolute", marginLeft:1480, marginTop:450}}>Enviar</button>
                    
                </form>
                
            
            </div>
        </>
    )
}

export default Consola