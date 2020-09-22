import React from "react";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import { TitleText } from 'bundles/utils'

const DividerText = ({ text }) => {
  return (
    <Grid item>
      <Typography>{text}</Typography>
      <Divider orientation="vertical" flexItem />
    </Grid>
  );
};

const Text = ({ text }) => {
    return (
      <Grid item>
        <Typography>{text}</Typography>
      </Grid>
    );
  };

const Calculator = () => {
  return (
    <Grid container justify="center"  direction="column"> 
      <Grid item>
        <TitleText text={'Launcher Calculator'}/>
      </Grid>
      <Grid item  >
        <Typography style={{paddingTop: 10}}>
          An open source orbital launch vehicle payload calculator by Epyphite
        </Typography>
      </Grid>
      <Grid item>
        <Grid container direction="row" spacing={2}>
          <DividerText text={"About"} />
          <DividerText text={"MIT License"} />
          <Text text={'GitHub'} />
        </Grid>
      </Grid>
    </Grid>
  );
};

export default Calculator;
