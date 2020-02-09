import { useState,useEffect } from 'react';
import { Grid } from "@material-ui/core";
import LoginForm from '../src/components/loginform'
import { Cookies } from 'react-cookie';
import Router from 'next/router'
export default function Login() {
    const cookies = new Cookies();
    const [myCookie,setMyCookie] = useState('')
    useEffect(()=> {
      // Prevent routing on the first load
      if (myCookie != '') {
        cookies.set('token',myCookie)
        Router.push('/play');
      }
    },[myCookie])

    return(
      <Grid
      container
      spacing={0}
      direction="column"
      alignItems="center"
      justify="center"
      style={{ minHeight: '100vh' }}
      >
          <LoginForm cookieHandler={setMyCookie}/>
      </Grid>   
    )
}