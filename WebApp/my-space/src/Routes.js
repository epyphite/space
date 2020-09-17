import React, { lazy, Suspense, Fragment } from "react";
import { Switch, Route, Redirect, useLocation } from "react-router-dom";
import Grid from "@material-ui/core/Grid";
import { AnimatePresence, motion } from "framer-motion";
import { ThemeProvider } from "@material-ui/styles";
import { PresentationLayout } from "./layout";
import Ephyphite from "images/Epyphite-White.png";
import { HEADER_COLOR } from "bundles/utils/color";
import MuiTheme from "./bundles/theme";

const Home = lazy(() => import("bundles/landing"));

const Routes = () => {
  const location = useLocation();

  const SuspenseLoading = () => {
    return (
      <Fragment>
        <Grid
       
          container
          direction="column"
          alignItems="center"
          justify="center"
          alignContent="center"
          style={{height: '100vh', width: '100%', backgroundColor: HEADER_COLOR}}
        >
          <Grid item>
            <div
              className="d-flex align-items-center flex-column vh-100 justify-content-center text-center py-3"
            >
              <div
                style={{ margin: "auto", width: "50%" }}
                className="d-flex align-items-center flex-column px-4"
              >
                <img
                  src={Ephyphite}
                  alt="Space"
                  style={{ height: 100 }}
                />
              </div>
              <div className="text-muted font-size-xl text-center pt-3">
               
              </div>
            </div>
          </Grid>
        </Grid>
      </Fragment>
    );
  };

  const pageVariants = {
    initial: {
      opacity: 0,
      scale: 0.99,
    },
    in: {
      opacity: 1,
      scale: 1,
    },
    out: {
      opacity: 0,
      scale: 1.01,
    },
  };

  const pageTransition = {
    type: "tween",
    ease: "anticipate",
    duration: 0.4,
  };

  return (
    <ThemeProvider theme={MuiTheme}>
      <AnimatePresence>
        <Suspense fallback={<SuspenseLoading />}>
          <Switch>
            <Redirect exact from="/" to="/home" />
            <Redirect exact from="/Login" to="/home" />
            <Route path={["/home"]}>
              <PresentationLayout>
                <Switch location={location} key={location.pathname}>
                  <motion.div
                    initial="initial"
                    animate="in"
                    exit="out"
                    variants={pageVariants}
                    transition={pageTransition}
                  >
                    <Route path="/home" component={Home} />
                  </motion.div>
                </Switch>
              </PresentationLayout>
            </Route>
          </Switch>
        </Suspense>
      </AnimatePresence>
    </ThemeProvider>
  );
};

export default Routes;
