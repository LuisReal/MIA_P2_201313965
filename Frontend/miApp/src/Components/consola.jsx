
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
    }


    return (
        <>
            <div style={{width:280, position:"absolute"}}>

                <form action="" onSubmit={mostrarDatos}>
                    <p name="consola" style={{height:450, width:1580, position:"absolute"}} ></p>
                    <input type="text" onChange={getDatos} name="comando" style={{height:50, width:1480, position:"absolute", marginTop:450}} placeholder="Ingrese comando" />
                    <button type="submit" className='btn btn-primary' style={{height:50, position:"absolute", marginLeft:1480, marginTop:450}}>Enviar</button>
                    
                </form>
                
            
            </div>
        </>
    )
}

export default Consola