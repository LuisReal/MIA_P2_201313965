
import { useState, useContext } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { UserContext } from "./usercontext";

export default function Login() {
  const [data, setData] = useState(
   
        {
            info: '',
            status: false
        }
  ) 
  const navigate = useNavigate()

  const { disk, particion } = useParams()
  const {value} = useContext(UserContext) //obtiene informacion de value desde otro componente(en particion.jsx)
  
  console.log("La info del id es: ", value)
  let username;
  let password;

  const validarUsuario = (e) => {
    e.preventDefault();
    
    username = e.target.username.value
    password = e.target.password.value


    let obj = {
        
        'Comand': "login -user="+ username + " -pass="+ password + " -id="+value

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
        {   info: res.data,
            status: res.status
        }

       )
    })

    if (data.status) {
        navigate(`/disk/${disk}/${particion}/sucess`)
    }else{
        alert("Usuario o contrasena invalidos");
    }
   
    
    
}

  return (
    <>
      
      
      <div style={{position: "relative",  marginLeft:280}}>

            <div className="sidenav">
                    
                    <h2>Sign in</h2>
                    <img className="image-bird" src="../../assets/bird.png" width="50" height="50" />
                    <img className="image-logo" src="../../assets/logo.jpg" width="100" height="100"/>  

                
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