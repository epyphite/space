import React from "react";
import Grid from "@material-ui/core/Grid";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import { HEADER_COLOR } from "bundles/utils/color";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  footer: {
   // position: "sticky",
     left: 0,
    bottom: 2,
    backgroundColor: HEADER_COLOR,
    //height: "180px",
  },
}));

const Footer = () => {
  const classes = useStyles();
  return (
    <Grid
      container
      justify="center"
      alignItems="center"
      alignContent="center"
      spacing={8}
      className={classes.footer}
    >
      <Grid item md={3}></Grid>
      <Grid item md={4}>
        <Grid container direction="column" spacing={2}>
          <Grid item>
            <Typography variant="h3" style={{ color: "#fff", fontWeight: 600 }}>
              {" "}
              Few Words About Epyphite
            </Typography>
          </Grid>
          <Grid item>
            <Typography style={{ color: "#fff", fontWeight: 500 }}>
              Our goal is to provide 'just in time' intelligence, data and
              quality software solutions to the aerospace sector, driving
              innovation, disruption and growth, pushing boundaries of deep
              technologies and integrating them for a common and social goal.
            </Typography>
          </Grid>
        </Grid>
      </Grid>
      <Grid item md={4}>
        <Grid container direction="column" spacing={2}>
          <Grid item>
            <Typography variant="h3" style={{ color: "#fff", fontWeight: 600 }}>
              {" "}
              Links
            </Typography>
          </Grid>
          <Grid item>
            <Typography style={{ color: "#fff", fontWeight: 500 }}>
              Terms & Conditions
            </Typography>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
};

export default Footer;
