

import React from 'react'

function Consola() {


  return (
    <>
        <div style={{width:280, position:"absolute"}}>
          <input type="text" style={{height:450, width:1580, position:"absolute"}} disabled />
          <input type="text" style={{height:50, width:1500, position:"absolute", marginTop:450}} placeholder="Ingrese comando" />
          <button type="submit" className='btn btn-primary' style={{height:50, position:"absolute", marginLeft:1500, marginTop:450}}>Enviar</button>
        </div>
    </>
  )
}

export default Consola