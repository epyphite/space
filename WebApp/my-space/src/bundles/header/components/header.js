import React from "react";
import AppBar from "@material-ui/core/AppBar";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import { Route, withRouter } from 'react-router-dom';
import Grid from "@material-ui/core/Grid";
import Toolbar from "@material-ui/core/Toolbar";
import whiteLogo from "images/Epyphite-White.png";
import TwitterIcon from "@material-ui/icons/Twitter";
import IconButton from "@material-ui/core/IconButton";
import Paper from "@material-ui/core/Paper";
import Button from "@material-ui/core/Button";
import {
  LIGHT_HEADER_COLOR,
  HEADER_COLOR,
  SECONDARY_COLOR,
} from "bundles/utils/color";
import { Login } from "bundles/Login";
import { SignUp } from "bundles/SignUp";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    [theme.breakpoints.up("sm")]: {
      paddingLeft: 54,
      paddingRight: 54,
    },
    [theme.breakpoints.down("xs")]: {
      paddingLeft: 20,
      paddingRight: 20,
    },
    paddingTop: 23,
    paddingBottom: 60,
    boxShadow: "0px 5px 5px rgba(0, 0, 0, 0.05)",
    border: "1px solid rgba(219, 220, 224, 0.5)",
    borderRadius: 5,
  },
  title: {
    flexGrow: 1,
  },
  appBar: {
    paddingLeft: 50,
    paddingRight: 50,
    backgroundColor: LIGHT_HEADER_COLOR,
  },
  twitter: {
    color: "#fff",
  },
  buttonColor: {
    color: "#fff",
    "&:hover": {
      color: SECONDARY_COLOR,
    },
  },
  content: {
    backgroundColor: HEADER_COLOR,
    height: "15%",
  },
  footer: {
    position: "fixed",
    left: 0,
    bottom: 0,
    backgroundColor: HEADER_COLOR,
    height: "15%",
  },
}));

const ButtonText = ({ classes, children, ...props }) => (
  <Button className={classes.buttonColor} {...props} color="inherit">
    {children}
  </Button>
);

const Menu = ({ classes, history }) => {
  return (
    <Grid>
      <ButtonText classes={classes}  onClick={() => history.push('/home')}  color="inherit">
        Home
      </ButtonText>
      <ButtonText classes={classes} onClick={() => history.push('/home/login')} color="inherit">
        Login
      </ButtonText>
      <ButtonText classes={classes}  onClick={() => history.push('/home/signup')}  color="inherit">
        Sign Up
      </ButtonText>
      <IconButton>
        <TwitterIcon className={classes.twitter} />
      </IconButton>
    </Grid>
  );
};

const Header = ({match, history}) => {
  const classes = useStyles();

  return (
    <React.Fragment>
      <div style={{ height: '100vh' }}>
        <div style={{ backgroundColor: HEADER_COLOR, height: 20 }}></div>
        <AppBar className={classes.appBar} position="sticky">
          <Toolbar>
            <img alt="brand-logo" width={113} height={32} src={whiteLogo} />
            <Grid className={classes.title}></Grid>
            <Menu classes={classes} history={history} />
          </Toolbar>
        </AppBar>
        <Grid
          container
          justify="center"
          alignItems="center"
          alignContent="center"
          className={classes.content}
        >
          <Typography variant="h1" style={{ color: "#fff", fontWeight: 600 }}>
            {" "}
            Launch Orbital Solutions
          </Typography>
        </Grid>

        <Grid container justify="center" alignContent="center" alignItems="center" style={{ marginTop: 50 }}>
          <Grid item xs={12} sm={8} md={6} lg={4}>
            <Paper className={classes.paper} elevation={2} >
               <Route path={`/home/login`} component={Login}/>
               <Route path={`/home/signup`} component={SignUp}/>
            </Paper>
          </Grid>
        </Grid>

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
                <Typography
                  variant="h3"
                  style={{ color: "#fff", fontWeight: 600 }}
                >
                  {" "}
                  Few Words About Epyphite
                </Typography>
              </Grid>
              <Grid item>
                <Typography style={{ color: "#fff", fontWeight: 500 }}>
                  Our goal is to provide 'just in time' intelligence, data and
                  quality software solutions to the aerospace sector, driving
                  innovation, disruption and growth, pushing boundaries of deep
                  technologies and integrating them for a common and social
                  goal.
                </Typography>
              </Grid>
            </Grid>
          </Grid>
          <Grid item md={4}>
            <Grid container direction="column" spacing={2}>
              <Grid item>
                <Typography
                  variant="h3"
                  style={{ color: "#fff", fontWeight: 600 }}
                >
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
      </div>
    </React.Fragment>
  );
};

export default withRouter(Header);
