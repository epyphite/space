import React from "react";
import { makeStyles, withStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import LinearProgress from "@material-ui/core/LinearProgress";
import {
  ERROR_COLOR,
  WARNING_COLOR,
  SECONDARY_COLOR,
} from "bundles/utils/color";

const useStyles = makeStyles((theme) => ({
  text: {
      marginTop: 7,
  },
  weak: {
    borderRadius: 5,
    backgroundColor: ERROR_COLOR,
  },
  fair: {
    borderRadius: 5,
    backgroundColor: WARNING_COLOR,
  },
  ok: {
    borderRadius: 5,
    backgroundColor: "#1a90ff",
  },
  strong: {
    borderRadius: 5,
    backgroundColor: SECONDARY_COLOR,
  },
}));

const BorderLinearProgress = withStyles((theme) => ({
  root: {
    height: 10,
    borderRadius: 5,
  },
  colorPrimary: {
    backgroundColor:
      theme.palette.grey[theme.palette.type === "light" ? 200 : 700],
  },
}))(LinearProgress);

const PasswordMeter = ({ value = 0, type = "weak", text = "" }) => {
  const classes = useStyles();

  const showPasswordBar = () => {
    if (value > 0) {
      return (
        <>
          <BorderLinearProgress
            variant="determinate"
            classes={{ bar: classes[type] }}
            value={value}
          />
          <Typography className={classes.text}> {text} </Typography>
        </>
      );
    }
  };

  return <>{showPasswordBar()}</>;
};

export default PasswordMeter;
