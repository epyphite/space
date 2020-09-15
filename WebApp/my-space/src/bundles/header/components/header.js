import React from "react";
import Grid from "@material-ui/core/Grid";
import { Calculator, Rocket } from "bundles/header/components/modules";

const Header = () => {
  return (
    <Grid container>
      <Grid item md={3} xs={12}>
        <Calculator />
      </Grid>
      <Grid item md={3} xs={12}>
        <Rocket />
      </Grid>
      <Grid item md={3} xs={12}></Grid>
    </Grid>
  );
};

export default Header;
