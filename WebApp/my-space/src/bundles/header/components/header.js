import React, { useEffect, useState } from "react";
import AppBar from "@material-ui/core/AppBar";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import { Route, withRouter } from "react-router-dom";
import Box from "@material-ui/core/Box";
import "react-perfect-scrollbar/dist/css/styles.css";
import PerfectScrollbar from "react-perfect-scrollbar";
import Grid from "@material-ui/core/Grid";
import Toolbar from "@material-ui/core/Toolbar";
import whiteLogo from "images/Epyphite-White.png";
import TwitterIcon from "@material-ui/icons/Twitter";
import IconButton from "@material-ui/core/IconButton";
import Paper from "@material-ui/core/Paper";
import Button from "@material-ui/core/Button";
import DoneIcon from "@material-ui/icons/Done";
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
    height: "180px",
  },
  completeImg: {
    marginTop: 50,
    color: SECONDARY_COLOR,
    size: 30,
  },
  completeHeader: {
    fontSize: 30,
    fontWeight: 500,
    marginTop: 40,
    marginBottom: 40,
  },
  iconBorder: {
    border: 1,
    borderColor: SECONDARY_COLOR,
  },
  circle: {
    border: `2px solid ${SECONDARY_COLOR}`,
    backgroundColor: "#FFFFFF",
    height: "110px",
    borderRadius: "50%",
    width: "110px",
  },
  login: {
    marginTop: 100, marginBottom: 30
  },
  homepage: {
    marginTop: 100, marginBottom: 30
  },
  signup: {
    marginTop: 30, marginBottom: 90
  }
}));

const ButtonText = ({ classes, children, ...props }) => (
  <Button className={classes.buttonColor} {...props} color="inherit">
    {children}
  </Button>
);

const HomePage = ({ classes }) => {
  return (
    <Box
      alignItems="center"
      display="flex"
      justifyContent="center"
      flexDirection="column"
    >
      <>
        <Grid
          container
          direction="column"
          justify="center"
          alignContent="center"
          alignItems="center"
          className={classes.circle}
        >
          <Grid item>
            <DoneIcon
              className={classes.completeImg}
              style={{ fontSize: "100px", paddingBottom: 48 }}
            />
          </Grid>
        </Grid>

        <Typography align="center" className={classes.completeHeader}>
          Thank you for registering on Launch Orbital Solutions.
        </Typography>
      </>
    </Box>
  );
};

const Menu = ({ classes, history, setPath }) => {
  return (
    <Grid>
      <ButtonText classes={classes} onClick={() => setPath("homepage")} color="inherit">
        Home
      </ButtonText>
      <ButtonText
        classes={classes}
        onClick={() => setPath("login")}
        color="inherit"
      >
        Login
      </ButtonText>
      <ButtonText
        classes={classes}
        onClick={() => setPath("signup")}
        color="inherit"
      >
        Sign Up
      </ButtonText>
      <IconButton>
        <TwitterIcon className={classes.twitter} />
      </IconButton>
    </Grid>
  );
};

const Header = ({ match, history }) => {
  const classes = useStyles();
  const [path, setPath] = useState("signup");

  const renderView = () => {
    if (path == "login") return <Login />;

    if (path == "signup") return <SignUp />;

    return <HomePage classes={classes} />;
  };

  return (
    <React.Fragment>
      <div>
        <div style={{ backgroundColor: HEADER_COLOR, height: 20 }}></div>
        <AppBar className={classes.appBar} position="sticky">
          <Toolbar>
            <img alt="brand-logo" width={113} height={32} src={whiteLogo} />
            <Grid className={classes.title}></Grid>
            <Menu classes={classes} history={history} setPath={setPath} />
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

        <Grid
          container
          justify="center"
          alignContent="center"
          alignItems="center"
          className={classes[path]}
        >
          <Grid item xs={12} sm={8} md={6} lg={4}>
           
              <Paper className={classes.paper} elevation={2}>
                {renderView()}
              </Paper>
    
          </Grid>
        </Grid>
      </div>
    </React.Fragment>
  );
};

export default withRouter(Header);
