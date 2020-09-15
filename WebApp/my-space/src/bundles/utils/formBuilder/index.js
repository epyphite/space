import React from "react";
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';

const TextTransform = ({
  input,
  setFormState,
  formState,
  setFormStateValidation,
  setPasswordMask,
  passwordMask,
}) => {
  const classes = useStyles();

  return (
    <Grid
      container
      style={{ marginBottom: 12 }}
      direction={input.labelDirection}
    >
      <Grid item xs={!input.labelDirection && 4}>
        <Typography className={classes.labelText}>
          {input.label}{" "}
          {input.required && !input.hideAsterix ? <RequiredIndicator /> : null}
        </Typography>
      </Grid>
      <Grid item xs={!input.labelDirection && 8} className={classes.container}>
        {InputTextComp(
          classes,
          input,
          setFormState,
          formState,
          setFormStateValidation,
          setPasswordMask,
          passwordMask
        )}
      </Grid>
    </Grid>
  );
};

const renderType = (
  input,
  setFormState,
  formState,
  setFormStateValidation,
  setPasswordMask,
  passwordMask
) => {
  switch (input.type) {
    case "text" || "password" || "email" || "number":
      return (
        <TextTransform
          setFormState={setFormState}
          input={input}
          formState={formState}
          setFormStateValidation={setFormStateValidation}
          setPasswordMask={setPasswordMask}
          passwordMask={passwordMask}
        />
      );

    case "select":
      return (
        <SelectTransform
          setFormState={setFormState}
          input={input}
          formState={formState}
        />
      );

    default:
      return (
        <TextTransform
          setFormState={setFormState}
          input={input}
          formState={formState}
          setFormStateValidation={setFormStateValidation}
          setPasswordMask={setPasswordMask}
          passwordMask={passwordMask}
        />
      );
  }
};

export const FormBuilder = ({
  formInput,
  setFormState,
  formState,
  setFormStateValidation,
  setPasswordMask,
  passwordMask,
}) => {
  return (
    <Fragment>
      {renderType(
        formInput,
        setFormState,
        formState,
        setFormStateValidation,
        setPasswordMask,
        passwordMask
      )}
    </Fragment>
  );
};
