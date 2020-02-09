
import React, { useState,useEffect } from 'react';
import { Grid } from "@material-ui/core";
import CreateForm from '../../components/CreateForm'
import { Cookies } from 'react-cookie';

export default function CreatePage() {
    const cookies = new Cookies();
    const [myCookie,setMyCookie] = useState('')
    useEffect(()=> {
      // Prevent routing on the first load
      if (myCookie != '') {
        cookies.set('token',myCookie)
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
          <CreateForm cookieHandler={setMyCookie}/>
      </Grid>   
    )
}