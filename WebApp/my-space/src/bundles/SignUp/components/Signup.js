import React, { useState } from "react";
import { FormBuilder } from "bundles/utils";
import { makeStyles } from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import { SECONDARY_COLOR } from "bundles/utils/color";
import Typography from "@material-ui/core/Typography";

const signUpQuestions = [
  {
    label: "First Name",
    placeholder: "Enter first name",
    type: "text",
    key: "firstName",
    required: true,
    hideAsterix: true,
    labelDirection: "column",
  },
  {
    label: "Last Name",
    placeholder: "Enter your lastname",
    type: "text",
    key: "lastName",
    required: true,
    hideAsterix: true,
    labelDirection: "column",
  },
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
    label: "Phone Number",
    placeholder: "Enter your phone number",
    type: "phoneNumber",
    key: "phoneNumber",
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
  {
    label: "Re Enter Password",
    placeholder: "",
    type: "password",
    key: "repassword",
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
  },
  text: {
    color: "#fff",
    fontWeight: 600,
  },
}));

const SignUp = () => {
  const [signUpData, setSignUp] = useState({});
  const [passwordMask, setPasswordMask] = useState(false);
  const classes = useStyles();

  const setFormInitialState = (value) => {
    setSignUp({ ...signUpData, ...value });
  };

  return <div> Sign Up</div>;
};

export default SignUp;
