import React from "react";
import { Header } from "bundles/Header";

const Landing = ({ match, ...props }) => (
  <div>
    <Header {...props} />
  </div>
);

export default Landing;
