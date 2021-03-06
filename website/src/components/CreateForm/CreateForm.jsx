import React, { useState } from 'react'
import { Grid,TextField, Button } from '@material-ui/core'
import axios from 'axios'
import { useHistory } from 'react-router-dom'
function CreateForm(props) {
    const history = useHistory()
    const [user,setUser] = useState({
        email:'',
        password:'',
        confirmPassword: ''
    })
    const handleChange = (input) => event => {
        setUser({
            ...user,[input]:event.target.value
        })
    }
    const createAccount = () => {
        // const setCookies  = props.cookieHandler;
        axios.post('/api/createAccount',user)
            .then(res => {
                const token = res.data.token
                props.cookieHandler(token)
                history.push('/play')
            })
            .catch(err => {
                console.log(err.data)
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
            <Grid item xs={12} md={12}>
                <TextField 
                onChange={handleChange('confirmPassword')}
                style={{marginBottom:'1vh',minWidth:'80vw'}} 
                variant="outlined" 
                label="Confirm Password" />
            </Grid>
            <Grid item xs={12}>
                <Button variant="outlined" onClick={() => createAccount()}> Create Account </Button>
            </Grid>
        </>
    )    
}
export default CreateForm;