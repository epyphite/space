import React, { useState } from "react";
import { FormBuilder, PasswordMeter, isValidState } from "bundles/utils";
import { makeStyles } from "@material-ui/core/styles";
import CircularProgress from '@material-ui/core/CircularProgress';
import Button from "@material-ui/core/Button";
import { SECONDARY_COLOR } from "bundles/utils/color";
import { postData } from "utils";
import Typography from "@material-ui/core/Typography";

const signUpQuestions = [
  {
    label: "First Name",
    placeholder: "Enter first name",
    type: "text",
    key: "firstName",
    customValidation: true,
    required: true,
    labelDirection: "column",
  },
  {
    label: "Last Name",
    placeholder: "Enter your lastname",
    customValidation: true,
    type: "text",
    key: "lastName",
    required: true,
    labelDirection: "column",
  },
  {
    label: "Email Address",
    placeholder: "Enter your email",
    customValidation: true,
    type: "email",
    key: "email",
    required: true,
    labelDirection: "column",
  },
  {
    label: "Phone Number",
    placeholder: "Enter your phone number",
    customValidation: true,
    type: "phone",
    key: "phoneNumber",
    required: true,
    labelDirection: "column",
  },
  {
    label: "Password",
    placeholder: "",
    type: "password",
    key: "password",
    required: true,
    labelDirection: "column",
  },
  {
    label: "Re Enter Password",
    placeholder: "",
    type: "password",
    errorMessage: "",
    key: "repassword",
    required: true,
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
    "&:hover": {
      backgroundColor: SECONDARY_COLOR,
    },
  },
  text: {
    color: "#fff",
    fontWeight: 600,
  },
}));

const SignUp = () => {
  const [formState, setFormState] = useState({});
  const [formStateValidation, setFormStateValidation] = useState({});
  const [loading, setLoading] = useState(false);
  const [passwordMask, setPasswordMask] = useState(false);
  const [passwordMatch, setPasswordMatch] = useState(true);
  const [passwordStrength, setPasswordStrength] = useState({});
  const classes = useStyles();

  const submitForm = async() => {
    setLoading(true)
    if (isValidState(signUpQuestions, setFormState, formState)) {
      console.log("Valid");

      console.log(formState)
      const data = await postData(`${process.env.REACT_APP_URL}/register`, formState);

      debugger
      
    } else {
      console.log("Not Valid");
    }

    setLoading(false);
  };

  const handleVerificationChange = (value) => {
    setFormStateValidation({ ...formStateValidation, ...value });
  };

  const setFormInitialState = (value) => {
    setFormState({ ...formState, ...value });

    if (value.password) {
      setPasswordStrength({ ...passwordChecker(value.password) });
      if (formState["repassword"])
        validatePasswordMatch(formState["repassword"]);
    }

    if (value.repassword) validatePasswordMatch(value.repassword);
  };

  const validatePasswordMatch = (currentPassword) => {
    const userPassword = formState["password"];

    if (currentPassword !== userPassword) {
      setPasswordMatch(false);
      return;
    }
    setPasswordMatch(true);
  };

  const passwordChecker = (value) => {
    const result = window?.zxcvbn(value);

    if (result.score == 4) {
      return { value: 100, type: "strong", text: "Very Strong" };
    }

    if (result.score == 3) {
      return { value: 75, type: "ok", text: "Good Password" };
    }

    if (result.score == 2) {
      return {
        value: 50,
        type: "fair",
        text: `${result.feedback.warning} ${result.feedback.suggestions}`,
      };
    }

    return {
      value: 25,
      type: "weak",
      text: result.feedback.warning || result.feedback.suggestions,
    };
  };

  return (
    <div>
      <form onSubmit={submitForm}>
        {signUpQuestions.map((question, index) => (
          <>
            <FormBuilder
              key={index}
              formInput={
                question.key == "repassword" && !passwordMatch
                  ? { ...question, errorMessage: "Password Does not Match" }
                  : question
              }
              setFormState={setFormInitialState}
              formState={formState}
              setPasswordMask={setPasswordMask}
              setFormStateValidation={handleVerificationChange}
              passwordMask={passwordMask}
            />
          </>
        ))}
        <PasswordMeter {...passwordStrength} />
        <Button
          fullWidth
          variant="contained"
          startIcon={
            loading ? (
              <CircularProgress size={15} disableShrink />
            ) : null
          }
          onClick={submitForm}
          className={classes.button}
        >
          <Typography className={classes.text}> Sign Up </Typography>
        </Button>
        <Typography style={{ textAlign: "center", marginTop: 5 }}>
          By Signing up you agree to our terms and condidtions
        </Typography>
      </form>
    </div>
  );
};

export default SignUp;
