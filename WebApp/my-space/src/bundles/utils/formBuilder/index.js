import React, { Fragment, memo } from "react";
import { makeStyles } from "@material-ui/core/styles";
import {
  Grid,
  Typography,
  OutlinedInput,
  FormControlLabel,
  Radio,
  RadioGroup,
  TextField,
  FormControl,
  FormHelperText,
} from "@material-ui/core";
import clsx from "clsx";
// import IconButton from "@material-ui/core/IconButton";
import MenuItem from "@material-ui/core/MenuItem";
// import FilterListIcon from "@material-ui/icons/FilterList";
import KeyboardArrowDown from "@material-ui/icons/KeyboardArrowDown";
const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

var nameRegex = /^[a-zA-Z\-]+$/;

const RequiredIndicator = memo(() => {
  return <span style={{ color: "#D74A4C" }}>*</span>;
});

const TransformButtonIcon = ({ ...props }) => {
  return <KeyboardArrowDown {...props} />;
};

const useStyles = makeStyles((theme) => {
  return {
    subInputText: {
      fontSize: 11,
      paddingTop: 2,
      color: theme.palette.grey[500],
    },
    subInputRedText: {
      fontSize: 11,
      paddingTop: 2,
      color: theme.palette.error.main,
    },
    button: {
      "&:hover": {
        backgroundColor: "transparent",
      },
    },
    errorText: {
      color: theme.palette.error.main,
    },
    labelText: {
      fontSize: 13,
      paddingTop: 10,
      marginBottom: 5,
      color: "rgba(0, 0, 0, 0.45)",
    },
    blackText: {
      fontSize: 13,
      paddingTop: 10,
      marginBottom: 5,
    },
    iconButtonn: {
      border: "1px solid #ced4da",
      marginLeft: 4,
      height: 41,
      borderRadius: 4,
      borderTopRightRadius: 0,
      borderRight: "0",
      borderBottomRightRadius: 0,
      backgroundColor: "#f5f5f5",
    },
    iconText: {
      fontSize: 16,
      fontWeight: 500,
      // paddingTop: 10,
      marginBottom: 5,
    },
    verificationFields: {
      textAlign: "center",
    },
  };
});

const emailValidator = (data) => {
  return {
    valid: re.test(String(data).toLowerCase()),
    label: "Please Enter a valid email",
  };
};
const validateText = (data) => {
  const isValidCharacters = nameRegex.test(String(data).toLowerCase());
  const isValidLength = data?.trim()?.length > 2;

  if (!isValidCharacters) {
    return {
      valid: false,
      label:
        "Please enter only valid characters. Valid characters include A-Z, a-z and -",
    };
  }

  if (!isValidLength) {
    return {
      valid: false,
      label: "Length is too short.",
    };
  }

  return { valid: true, label: "" };
};

const validator = (data) => {
  return {
    email: emailValidator(data),
    text: validateText(data),
  };
};

const validateFields = (input, data) => {
  return input.validator ? input.validator(data) : validator(data)[input.type];
};

export const capitalizeFirstWord = (word) =>
  word && word.charAt(0).toUpperCase() + word.slice(1).toLowerCase();

const InputTextComp = (
  classes,
  input,
  setFormState = () => "",
  formState = {},
  setFormStateValidation = () => ""
) => {
  let rows = null;
  const multiline = input.type === "textArea" ? true : false;
  if (multiline) rows = 5;

  let nonValid = false;
  if (
    formState[input.key] !== undefined &&
    formState[input.key] === "" &&
    input.required
  ) {
    nonValid = true;
  }

  let customValidator = {};
  if (
    !nonValid &&
    formState[input.key] !== undefined &&
    input.customValidation
  ) {
    customValidator = validateFields(input, formState[input.key]); // input.validator ? input.validator(formState[input.key]) : validator(formState[input.key])[input.type]
    if (!customValidator.valid) {
      nonValid = true;
    }
  }

  let defaultValue = input.defaultValue;
  if (input.type === "number") {
    defaultValue = parseInt(input.defaultValue);
  }

  if (input.ref) {
    return (
      <Fragment>
        <OutlinedInput
          fullWidth
          placeholder={input.placeholder || ""}
          inputProps={{ style: input.style || {}, ...input.inputProps }}
          defaultValue={defaultValue || ""}
          disabled={input.disabled}
          error={nonValid}
          ref={input.ref}
          multiline={multiline}
          type={input.type}
          onChange={(e) =>
            input.type === "number"
              ? setFormState({ [input.key]: parseInt(e.target.value) })
              : input.capitalize
              ? setFormState({
                  [input.key]: capitalizeFirstWord(e.target.value),
                })
              : setFormState({ [input.key]: e.target.value })
          }
          onBlur={input.onBlur}
          rows={rows}
        />
        {nonValid || input.subTitle ? (
          <Typography
            className={clsx(
              { [classes.errorText]: nonValid },
              classes.subInputText
            )}
          >
            {nonValid
              ? customValidator.label
                ? customValidator.label
                : `Please enter ${input.label}`
              : input.subTitle}
          </Typography>
        ) : null}
      </Fragment>
    );
  }

  return (
    <Fragment>
      <OutlinedInput
        fullWidth
        placeholder={input.placeholder || ""}
        inputProps={{ style: input.style || {}, ...input.inputProps }}
        defaultValue={defaultValue || ""}
        value={formState[input.key] || ""}
        error={nonValid}
        disabled={input.disabled}
        ref={input.ref}
        multiline={multiline}
        type={input.type}
        onChange={(e) => {
          input.type === "number"
            ? setFormState({ [input.key]: parseInt(e.target.value) })
            : input.capitalize
            ? setFormState({ [input.key]: capitalizeFirstWord(e.target.value) })
            : setFormState({ [input.key]: e.target.value });

          if (input.customValidation) {
            setFormStateValidation({
              [input.key]: validateFields(input, e.target.value)?.valid,
            });
          }
        }}
        onBlur={input.onBlur}
        rows={rows}
      />
      {nonValid || input.subTitle ? (
        <Typography
          className={clsx(
            { [classes.errorText]: nonValid },
            input.subTitleColor ? classes.subInputRedText : classes.subInputText
          )}
        >
          {nonValid
            ? customValidator.label
              ? customValidator.label
              : `Please enter ${input.label}`
            : input.subTitle}
        </Typography>
      ) : null}
    </Fragment>
  );
};

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

const generateRadioType = (
  input,
  _classes,
  setFormState,
  formState,
  nonValid
) => {
  return (
    <FormControl error={nonValid}>
      <RadioGroup
        style={{ display: "flex", flexDirection: "row" }}
        name={input.label}
        onChange={() => ""}
      >
        {input.fields.map((field, i) => {
          return (
            <FormControlLabel
              style={{ marginBottom: 0 }}
              key={`${input.label}${i}`}
              value={field}
              control={
                <Radio
                  disabled={input.disabled}
                  onChange={(e) =>
                    setFormState({ [input.key]: e.target.value })
                  }
                  checked={field === formState[input.key]}
                  color="secondary"
                />
              }
              label={field}
            />
          );
        })}
      </RadioGroup>
      {/* {nonValid || input.subTitle ? (
          <FormHelperText style={{ marginTop: 0 }}>
            <Typography
              className={clsx(
                { [classes.errorText]: nonValid },
                classes.subInputText
              )}>
              {nonValid ? `please specify ${input.label}` : input.subTitle}
            </Typography>
          </FormHelperText>
        ) : null} */}
    </FormControl>
  );
};

const remapField = (input) =>
  input.fields.map((field) => {
    if (typeof field !== "object") {
      return { label: field, value: field };
    }

    return field;
  });

const generateSelectType = (
  input,
  classes,
  setFormState,
  formState,
  nonValid
) => {
  const remap = input.fields && remapField(input);

  return (
    <Grid container item xs={12}>
      {" "}
      {SelectFieldComp(
        classes,
        input,
        remap,
        setFormState,
        formState,
        nonValid
      )}
    </Grid>
  );
};

const SelectFieldComp = (
  classes,
  input,
  mappedValue,
  setFormState = () => "",
  formState = {},
  nonValid,
  keyValue
) => {
  const value = mappedValue || (input.fields && remapField(input));
  let enteredValue = "";
  let addedValue = "";
  // let nonValid = false;
  if (keyValue) {
    addedValue = `-${keyValue}`;
  }
  if (formState && formState[`${input.key}${addedValue}`]) {
    enteredValue = formState[`${input.key}${addedValue}`];
  }

  // if (
  //   formState &&
  //   formState[`${input.key}${addedValue}`]?.length < 1 &&
  //   input.required
  // ) {
  //   nonValid = true;
  // }

  let multiple = null;

  return (
    <Fragment>
      {input.multiple ? (
        multiple
      ) : (
        <TextField
          select
          fullWidth
          onChange={(e) =>
            setFormState({ [`${input.key}${addedValue}`]: e.target.value })
          }
          defaultValue={input.defaultValue}
          value={enteredValue || input.defaultValue}
          error={nonValid}
          placeholder={"DD"}
          disabled={input.disabled}
        //  iconcomponent={TransformButtonIcon}
          variant="outlined"
        >
          {value.map((option) => (
            <MenuItem
              key={`${option.value}--${input.label}`}
              value={option.value}
            >
              {option.label}
            </MenuItem>
          ))}
        </TextField>
      )}
      {/* {nonValid || input.subTitle ? (
          <Typography
            className={clsx(
              { [classes.errorText]: nonValid },
              classes.subInputText
            )}>
            {nonValid ? `Please enter ${input.key}` : input.subTitle}
          </Typography>
        ) : null} */}
    </Fragment>
  );
};

const SelectTransform = ({ input, setFormState, formState }) => {
  const classes = useStyles();
  let nonValid = false;
  if (
    formState[input.key] !== undefined &&
    formState[input.key] === "" &&
    input.required
  ) {
    nonValid = true;
  }

  const selectType = {
    select: generateSelectType(
      input,
      classes,
      setFormState,
      formState,
      nonValid
    ),
    radio: generateRadioType(input, classes, setFormState, formState, nonValid),
  };

  return (
    <Grid
      container
     // style={{ marginBottom: 12 }}
      direction={input.labelDirection}
    >
      <Grid item xs={!input.labelDirection && 12}>
        <Grid
          item
          container
          direction="row"
          xs={12}
          spacing={0}
          justify={!input.labelDirection ? "flex-end" : "flex-start"}
        >
          {selectType[input.type]}
        </Grid>
      </Grid>
      {nonValid || input.subTitle ? (
        <FormHelperText style={{ marginTop: 0 }}>
          <Typography
            className={clsx(
              { [classes.errorText]: nonValid },
              classes.subInputText
            )}
          >
            {nonValid ? `Please enter ${input.label}` : input.subTitle}
          </Typography>
        </FormHelperText>
      ) : null}
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

const FormBuilder = ({
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

export default FormBuilder;
