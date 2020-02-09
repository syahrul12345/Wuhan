import React,{useState,useEffect} from 'react';
import { Grid, Typography, Button, Card, Avatar, CardContent } from '@material-ui/core'
import axios from 'axios'
import './Home.css';

function Home() {
  const [count,setCount] = useState(0)
  const [deathCount,setDeathCount] = useState(0)
  const incrementPrayer = () => {
    const payload = {
      message:"increment!"
    }
    axios.post("/api/save",payload)
      .then((res) => {
        setCount(res.data.count)
      }).catch((err) => {
        console.log(err)
      })
  }
  useEffect(() => {
    // ALso returns the death count...
    axios.get("/api/get")
      .then((res) => {
        setCount(res.data.count)
        setDeathCount(res.data.death)
      }).catch((err) => {
        console.log(err)
        console.log("Couldnt connect to database")
      })
  })
  return (
    <Grid
    container
    direction="column"
    justify="center"
    alignItems="center"
    alignContent="center"
    style={{minHeight:'93vh'}}>
      <Grid item xs={12}>
        <Typography variant="h1" style={{textAlign:'center'}}> {deathCount} <img style={{maxHeight:'160',maxWidth:'160'}} src="/skull.png"/>  </Typography>
        <Typography variant="body1" style={{textAlign:'center'}}>{deathCount} have already died.... the virus is coming for <strong>YOU.</strong></Typography>
        
      </Grid>
      <Grid item xs={12}>
      <Button href="/create" style={{textAlign:'center'}} > Create account</Button>
      {/* <AdSense.Google
        client='ca-pub-9373441186970265'
        slot='7806394673'
        style={{ display: 'block' }}
        layout='in-article'
        format='fluid'
      /> */}
      </Grid>
      <Grid item xs={12}>
        <Typography variant="h1"> {count} <img src="/prayer.png"/>  </Typography>
      </Grid>
      <Grid item xs={12}>
        <Button 
        variant="contained"
        colour="primary"
        onClick={incrementPrayer}
        style={{
          marginTop:'2vh',
          marginBottom:'2vh'}}> <Typography variant="body1">1 Click 1 Prayer </Typography> </Button>
      </Grid>
      <Grid item xs={12}>
        <Typography variant="body1" style={{textAlign:'center'}}> Do not underestimate the power of positive thought! </Typography>
        <Typography variant="body1" style={{textAlign:'center'}}> Send a wireless prayer to the kids! </Typography>
        <Typography variant="subtitle1" style={{textAlign:'center'}}> Seriously, dont eat bats </Typography>
      </Grid>
    </Grid>
  );
}

export default Home;
