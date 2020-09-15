import React from "react";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import { TitleText } from "bundles/utils";

const Text = ({ text }) => {
  return (
    <Grid item>
      <Typography>{text}</Typography>
    </Grid>
  );
};

const Rocket = () => {
  return (
    <Grid container direction="column">
      <Grid item>
        <TitleText text={"Select Rocket"} />
      </Grid>
      <Grid item>
        <Typography></Typography>
      </Grid>
      <Grid item>
        <Grid container direction="row" spacing={2}>
          <Text text={"Space X"} />
        </Grid>
      </Grid>
    </Grid>
  );
};

export default Rocket;
