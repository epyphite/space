import React from "react";
import Grid from "@material-ui/core/Grid";
import { Calculator, Rocket, Result } from "bundles/Dashboard/components/modules";

const Header = () => {

  return (
    <Grid container justify="flex-start" style={{backgroundColor: 'rgb(21, 62, 82)', color: '#fff'}} spacing={4}>
      <Grid item md={4} xs={12}>
        <Calculator />
      </Grid>
      <Grid item md={3} xs={12}>
        <Rocket />
      </Grid>
      {/* <Grid item md={2} xs={12}>
       
      </Grid> */}
      <Grid item md={5} xs={12}>
        <Result />
      </Grid>
    </Grid>
  );
};

export default Header;
