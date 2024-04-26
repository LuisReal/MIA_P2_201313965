import React from 'react'
import { useContext } from "react";
import { UserContext } from "./usercontext";

function Archivo() {

    const {value} = useContext(UserContext)

    return (
        <>
            <div style={{marginLeft:280, border:"1px solid blue"}}>
                <textarea defaultValue={value} style={{width:1000, height:300}}></textarea>
            </div>
        </>
    
    )
}

export default Archivo