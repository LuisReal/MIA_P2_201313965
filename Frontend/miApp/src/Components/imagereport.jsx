import React, { useContext } from 'react';
import { Graphviz } from "graphviz-react";

import { GrafoContext } from "./usercontext";

function ImageReport() {

    const {grafo} = useContext(GrafoContext)

    const style = {
        
        width: 1500 || "100%",
        height: 800 || "100%"
      };

    return (
        <div className="graph" style={{marginLeft:250, marginTop:50}}>
            {/*<Graphviz dot={grafo}/>*/}
            <Graphviz
              dot={grafo}
              options={{
                useWorker: false,
                ...style,
                zoom: true
                //...props
              }}
            />,
        </div>
    )
}

export default ImageReport