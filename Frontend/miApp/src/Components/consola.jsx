
import {useState, useContext} from 'react'
import React from 'react'
import { UserContext } from "./usercontext";

function Consola() {

    const {setValue} = useContext(UserContext) //obtiene informacion de value desde otro componente(en particion.jsx)

    const [datos, setDatos] = useState(
        {
            comando: ''
        }
    )

    const [showData, setData] = useState(
        {
            dot: '',
            info: '',
            status: false
        }
    )

    const getDatos = (event) => {
        event.preventDefault();
        setDatos({
            ...datos, //esto es para hacer una copia del contenido anterior con el nuevo(de lo contrario se borrara y sustituira)
            [event.target.name] : event.target.value
        })
        
       //console.log(datos.comando)
    }

    const mostrarDatos = (event) => {
        event.preventDefault();
        
        //console.log(datos.comando)
        
        let obj = {
        
            'Comand': datos.comando
        }

        console.log("fetch de insert: ",datos.comando)
    
        fetch(`http://18.216.113.114:3000/insert`,{
                  
        method : 'POST',
        body: JSON.stringify(obj),
        headers:{
        'Content-Type': 'application/json'   
        }
        }).then(response => response.json()
        ).catch(err =>{
            console.error(err)
        }).then(res =>{
            setValue(res.dot)

            setData(
                {   dot: res.dot,
                    info: res.data,
                    status: res.status
                }
            )
             
        })

        document.getElementById("entrada").value = ""
        
    }

    return (
        <>
            <div style={{position:"relative",  marginLeft:280, border:"1px solid blue", height:500}}>

                <form action="" onSubmit={mostrarDatos}>
                    <p  style={{overflowY: "scroll", whiteSpace: "pre-line", height:450, width:1550, border:"3px solid black", position:"absolute"}}  >{showData.info}</p>

                    <textarea id="entrada" onChange={getDatos} name="comando" style={{overflowY: "scroll", whiteSpace: "pre-line", height:150, width:1480, position:"absolute", marginTop:450}} placeholder="Ingrese comando" ></textarea>
                    <button type="submit" className='btn btn-primary' style={{height:50, position:"absolute", marginLeft:1480, marginTop:450}}>Enviar</button>
                    
                </form>
                
            
            </div>
        </>
    )
}

export default Consola