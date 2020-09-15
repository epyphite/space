import React, { lazy, Suspense, Fragment } from "react";
import { Switch, Route, Redirect, useLocation } from "react-router-dom";
import { AnimatePresence, motion } from "framer-motion";
import { ThemeProvider } from "@material-ui/styles";
import { PresentationLayout } from "./layout";
import MuiTheme from "./bundles/theme";


const Home = lazy(() => import('bundles/landing'));

const Routes = () => {
  const location = useLocation();

  const SuspenseLoading = () => {
    return (
      <Fragment>
        <div className="d-flex align-items-center flex-column vh-100 justify-content-center text-center py-3">
          <div className="d-flex align-items-center flex-column px-4"></div>
          <div className="text-muted font-size-xl text-center pt-3">
            Loading ...
          </div>
        </div>
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
