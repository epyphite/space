import React from "react";
import Header from "./components/header";
import Dashbord from "./components/center";
import Grid from "@material-ui/core/Grid";

const Rocket = () => {
  return (
    <Grid container direction="column" spacing={2}>
      <Grid xs item>
        <Header />
      </Grid>
      <Grid xs item>
        <Dashbord />
      </Grid>
    </Grid>
  );
};

export default Rocket;
