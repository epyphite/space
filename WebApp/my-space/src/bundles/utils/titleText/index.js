import React from "react";
import Typography from "@material-ui/core/Typography";

const TitleText = ({ text, ...props }) => (
  <Typography style={{ fontWeight: 600, fontSize: 26 }} {...props} >{text}</Typography>
);

export default TitleText;
