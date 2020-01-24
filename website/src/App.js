import React,{useState,useEffect} from 'react';
import { Grid, Typography, Button, Card, Avatar, CardContent } from '@material-ui/core'
import axios from 'axios'
import './App.css';

function App() {
  const [count,setCount] = useState(0)
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
    axios.get("/api/get")
      .then((res) => {
        setCount(res.data.count)
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
      <Grid item xs={12}>
        <Typography variant="h6" style={{marginTop:'2vh',textAlign:'center'}}> THE TEAM</Typography>
      </Grid>
      <Grid container spacing={2} direction="row" alignItems="center" justify="center">
        <Grid item xs={6} md={6}>
          <Card>
            <CardContent>
              <Grid item xs={6}>
                <Avatar alt="Deon Lim" src="/deon.jpg"/>
              </Grid>
              <Grid item xs={6}>
              <Typography variant="body1"> Deon Lim </Typography>
              </Grid>
            </CardContent>
          </Card>
        </Grid>
        <Grid item xs={6} md={6}>
          <Card>
            <CardContent>
              <Grid item xs={6}>
                <Avatar alt="Vince Toh" src="/vince.jpg"/>
              </Grid>
              <Grid item xs={6}>
              <Typography variant="body1"> Vince Toh</Typography>
              </Grid>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Grid>
  );
}

export default App;