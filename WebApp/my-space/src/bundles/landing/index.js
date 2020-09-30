import React from "react";
import { Header } from "bundles/header";

const Landing = ({ match, ...props }) => (
    <div style={{ height: "100vh" }}>
        <Header {...props} />
      
    </div>
  
);

export default Landing;
