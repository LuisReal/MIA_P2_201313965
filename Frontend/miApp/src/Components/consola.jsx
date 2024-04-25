
import {useState} from 'react'
import React from 'react'

function Consola() {

    const [datos, setDatos] = useState(
        {
            comando: ''
        }
    )

    const [showData, setData] = useState(
        {
            info: ''
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
        
        let obj = {
        
            'Comand': datos.comando
        }
    
        fetch(`http://localhost:3000/insert`,{
                  
        method : 'POST',
        body: JSON.stringify(obj),
        headers:{
        'Content-Type': 'application/json'   
        }
        }).then(response => response.json()
        ).catch(err =>{
            console.error(err)
        }).then(res =>{
           
           setData(
            {info: res.data}
        )
             
        })
        
    }


   

    return (
        <>
            <div style={{position:"relative",  marginLeft:280, border:"1px solid blue", height:500}}>

                <form action="" onSubmit={mostrarDatos}>
                    <p style={{whiteSpace: "pre-line", height:450, width:1550, border:"3px solid black", position:"absolute"}}  >{showData.info}</p>

                    <input type="text" onChange={getDatos} name="comando" style={{height:50, width:1480, position:"absolute", marginTop:450}} placeholder="Ingrese comando" />
                    <button type="submit" className='btn btn-primary' style={{height:50, position:"absolute", marginLeft:1480, marginTop:450}}>Enviar</button>
                    
                </form>
                
            
            </div>
        </>
    )
}

export default Consola