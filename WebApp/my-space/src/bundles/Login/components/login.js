import React, { useState } from "react";
import { FormBuilder } from "bundles/utils";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import { postData } from "utils";
import {
  SECONDARY_COLOR,
} from "bundles/utils/color";
import Typography from "@material-ui/core/Typography";

const loginQuestions = [
  {
    label: "Email Address",
    placeholder: "Enter your email",
    type: "email",
    key: "email",
    required: true,
    hideAsterix: true,
    labelDirection: "column",
  },
  {
    label: "Password",
    placeholder: "",
    type: "password",
    key: "password",
    required: true,
    hideAsterix: true,
    labelDirection: "column",
  },
];

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  button: {
    backgroundColor: SECONDARY_COLOR,
    color: "#fff",
    padding: 9,
    marginTop: 15,
    '&:hover': {
        backgroundColor: SECONDARY_COLOR,
    }
  },
  text: {
    color: "#fff",
    fontWeight: 600,
  },
}));

const Login = () => {
  const [loginData, setLoginData] = useState({});
  const [passwordMask, setPasswordMask] = useState(false);
  const classes = useStyles();
  const submitForm = () => {};

  const setFormInitialState = (value) => {
    setLoginData({ ...loginData, ...value });
  };

  const submitLogin = async() => {
    console.log('Hello boy')

    const data = await postData(`${process.env.REACT_APP_URL}/launchapi/api/v1/login `, loginData);

  
  }

  return (
    <div>
      <form onKeyUp={submitForm}>
        {loginQuestions.map((question, index) => (
          <>
            <FormBuilder
              key={index}
              formInput={question}
              setFormState={setFormInitialState}
              formState={loginData}
              setPasswordMask={setPasswordMask}
              passwordMask={passwordMask}
            />
          </>
        ))}
        <Button onClick={submitLogin} fullWidth variant="contained" className={classes.button}>
          <Typography className={classes.text}> Login </Typography>
        </Button>
        <Button
          variant="text"
          size="small"
          disableRipple={true}
          style={{ padding: 0, paddingTop: 10, color: "#942B32", fontSize: 13 }}
          disableFocusRipple={true}
        >
          Forgot your password?
        </Button>
      </form>
    </div>
  );
};
export default Login;
