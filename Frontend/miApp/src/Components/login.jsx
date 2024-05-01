
import { useState, useContext } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { UserContext } from "./usercontext";
import birdIMG from "../../assets/bird.png";
import logoIMG from "../../assets/logo.jpg";
import fondoIMG from "../../assets/fondo.jpg"; // importantes hacer esto para que las imagenes sean importadas con npm run build
import laptopIMG from "../../assets/laptop.jpg"; // importantes hacer esto para que las imagenes sean importadas con npm run build
import passwordIMG from "../../assets/password.png"; // importantes hacer esto para que las imagenes sean importadas con npm run build
import userIMG from "../../assets/user.png"; // importantes hacer esto para que las imagenes sean importadas con npm run build


export default function Login() {
  const [data, setData] = useState(
   
        {
            info: '',
            status: false
        }
  )

  var estado = false
  
  const navigate = useNavigate()

  const { disk, particion } = useParams()
  const {value} = useContext(UserContext) //obtiene informacion de value desde otro componente(en particion.jsx)
  
  //console.log("La info del id es: ", value)
  let username;
  let password;

  const validarUsuario = async(e) => {
    e.preventDefault();
    
    username = e.target.username.value
    password = e.target.password.value


    let obj = {
        
        'Comand': "login -user="+ username + " -pass="+ password + " -id="+value

    }

    const respuesta = await fetch(`http://localhost:3000/insert`,{
              
    method : 'POST',
    body: JSON.stringify(obj),
    headers:{
    'Content-Type': 'application/json'   
    }
    }).then(response => response.json()
    ).catch(err =>{
        console.error(err)
    }).then(res =>{
        //console.log("el valor data.status en la respuesta es: ",res.status)
       setData(
        {   info: res.data,
            status: res.status
        }

       )

       estado = res.status

       //console.log("el valor de variable estado es: ",estado)

        if (estado) {
            navigate(`/disk/${disk}/${particion}/sistema/raiz`)
        }else{
            alert("No se pudo iniciar sesion");
        }
    })

    
   
    
    
}

  return (
    <>
      
      
      <div style={{position: "relative",  marginLeft:280}}>

            <div className="sidenav">
                    
                    <h2>Sign in</h2>
                    <img className="image-bird" src={birdIMG} width="50" height="50" />
                    <img className="image-logo" src={logoIMG} width="100" height="100"/>  

                
            </div>
            <div className="main">
                    
                    <form onSubmit={validarUsuario}>

                        <div className="banner">

                            
                        </div>

                        <div className="form-group">
                            
                            <input type="text" className="form-user" placeholder="User Name" name="username" />
                        </div>
                        <div className="form-group">
                            
                            <input type="password" className="form-password" placeholder="Password" name="password" />
                        </div>

                        <div className="botones">
                        
                            <button type="submit" className="btnL">Login</button>
                        
                        
                        </div>

                    </form>

                
            
            </div>
      
      </div>
    
     
    </>
   )
 }