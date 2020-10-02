import React from "react";
import {  Box, Button } from "@material-ui/core";

const Advanced = () => {
  return (
    <Box
      style={{
        width: "100%",
        padding: 20,
        zIndex: 9999999,
        backgroundColor: "#FAFAFA",
        position: "-webkit-sticky",
        position: "sticky",
        top: 0,
      }}
      textAlign="center"
    >
      <Button style={{ fontSize: 16 }} variant="outlined">
        Advanced....
      </Button>
    </Box>
  );
};

export default Advanced;
