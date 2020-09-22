import React, { Fragment } from "react";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import { TitleText } from "bundles/utils";

const Text = ({ text }) => {
  return (
    <Grid item xs={6}>
      <Typography variant="h5" >{text}</Typography>
    </Grid>
  );
};

const payload = [
  {
    title: "Payload Mass(calculated)",
    value: "300kg",
  },
  {
    title: "Payload to life mass off ratio",
    value: "300kg",
  },
  {
    title: "Payload Change",
    value: "10%",
  },
  {
    title: "Payload Mass(calculated)",
    value: "300kg",
  },
];

const Result = () => {
  return (
    <Grid container justify="center" direction="column">
      <Grid item>
        <TitleText text={"Result"} />
      </Grid>
      <Grid item></Grid>
      <Grid item>
        <Grid container direction="row" spacing={1}>
          {payload.map((item, index) => (
            <Fragment key={index}>
              <Text text={item.title} />
              <Text text={item.value} />
            </Fragment>
          ))}
        </Grid>
      </Grid>
    </Grid>
  );
};

export default Result;
