
import React, { useState,useEffect } from 'react';
import { Grid } from "@material-ui/core";
import { Cookies } from 'react-cookie';
import LoginForm from '../../components/LoginForm';

export default function LoginPage() {
    const cookies = new Cookies();
    const [myCookie,setMyCookie] = useState('')
    useEffect(()=> {
      // Prevent routing on the first load
      if (myCookie !== '') {
        cookies.set('x-wuhan-cookie',`bearer ${myCookie}`)
      }
    },[cookies,myCookie])

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