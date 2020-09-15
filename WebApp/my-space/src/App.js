import React, { Component } from "react";
import { BrowserRouter } from "react-router-dom";
import ReduxProvider from "./bundles/client/components/Redux/Provider";
import CssBaseline from "@material-ui/core/CssBaseline";
import Routes from "./Routes";
import ScrollToTop from "./utils/ScrollToTop";

class App extends Component {
  render() {
    return (
      <ReduxProvider>
        <BrowserRouter>
          <CssBaseline />
          <ScrollToTop>
            <Routes />
          </ScrollToTop>
        </BrowserRouter>
      </ReduxProvider>
    );
  }
}

export default App;
