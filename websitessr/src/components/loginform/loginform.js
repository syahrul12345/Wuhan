import { useState } from 'react'
import { Grid,TextField, Button } from '@material-ui/core'
import axios from 'axios'

function LoginForm(props) {
    const [user,setUser] = useState({
        email:'',
        password:''
    })
    const handleChange = (input) => event => {
        setUser({
            ...user,[input]:event.target.value
        })
    }
    const login = () => {
        // const setCookies  = props.cookieHandler;
        axios.post('/api/login',user)
            .then(res => {
                const token = res.data.token
                props.cookieHandler(token)
            })
            .catch(err => {
                console.log(err)
            })
    }
    return(
        <>
            <Grid item xs={12} md={12}>
                <TextField 
                onChange={handleChange('email')}
                style={{marginBottom:'1vh',minWidth:'80vw'}} 
                variant="outlined" 
                label="Email" />
            </Grid>
            <Grid item xs={12} md={12}>
                <TextField 
                onChange={handleChange('password')}
                style={{marginBottom:'1vh',minWidth:'80vw'}} 
                variant="outlined"  
                label="Password" />
            </Grid>
            <Grid item xs={12}>
                <Button variant="outlined" onClick={() => login()}> Login </Button>
            </Grid>
        </>
    )    
}
export default LoginForm;