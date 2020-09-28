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
  Slider,
  FormHelperText,
} from "@material-ui/core";
import InputAdornment from "@material-ui/core/InputAdornment";
import clsx from "clsx";
// import IconButton from "@material-ui/core/IconButton";
import MenuItem from "@material-ui/core/MenuItem";
import MuiPhoneNumber from "material-ui-phone-number";
import Checkbox from "@material-ui/core/Checkbox";
// import FilterListIcon from "@material-ui/icons/FilterList";
import { parsePhoneNumberFromString } from "libphonenumber-js";
import KeyboardArrowDown from "@material-ui/icons/KeyboardArrowDown";
import IconButton from "@material-ui/core/IconButton";
import Visibility from "@material-ui/icons/Visibility";
import { ERROR_COLOR } from "bundles/utils/color";
import VisibilityOff from "@material-ui/icons/VisibilityOff";
const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

var nameRegex = /^[a-zA-Z\-]+$/;

const RequiredIndicator = memo(() => {
  return <span style={{ color: ERROR_COLOR }}>*</span>;
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
    subText: {
      fontSize: 15,
      fontWeight: 400,
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
  setFormStateValidation = () => "",
  setPasswordMask = () => "",
  passwordMask = false
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

  if (input.type == "password") {
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
          type={passwordMask ? "text" : input.type}
          onChange={(e) => {
            input.type === "number"
              ? setFormState({ [input.key]: parseInt(e.target.value) })
              : input.capitalize
              ? setFormState({
                  [input.key]: capitalizeFirstWord(e.target.value),
                })
              : setFormState({ [input.key]: e.target.value });

            if (input.customValidation) {
              setFormStateValidation({
                [input.key]: validateFields(input, e.target.value)?.valid,
              });
            }
          }}
          onBlur={input.onBlur}
          rows={rows}
          endAdornment={
            <InputAdornment position="end">
              <IconButton
                disableRipple={true}
                className={classes.button}
                aria-label="toggle password visibility"
                onClick={(e) => setPasswordMask(!passwordMask)}
                onMouseDown={(e) => e.preventDefault()}
              >
                {passwordMask ? <Visibility /> : <VisibilityOff />}
              </IconButton>
            </InputAdornment>
          }
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
        ) : (
          input.errorMessage && (
            <Typography className={clsx({ [classes.errorText]: true })}>
              {" "}
              {input.errorMessage}
            </Typography>
          )
        )}
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
        value={formState[input.key] ||  defaultValue ||  ""}
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

const TextTransformDouble = ({
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
      // style={{ marginBottom: 12 }}
      spacing={1}
      direction={input.labelDirection}
    >
      <Grid item xs={!input.labelDirection && 3}>
        <Typography className={classes.labelText}>
          {input.label}{" "}
          {input.required && !input.hideAsterix ? <RequiredIndicator /> : null}
        </Typography>
      </Grid>
      <Grid item xs={!input.labelDirection && 3} className={classes.container}>
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
      <Grid item xs={!input.labelDirection && 3}>
        <Typography className={classes.labelText}>
          {input.label}{" "}
          {input.required && !input.hideAsterix ? <RequiredIndicator /> : null}
        </Typography>
      </Grid>
      <Grid item xs={!input.labelDirection && 3} className={classes.container}>
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

const CheckBoxComp = (
  classes,
  input,
  setFormState = () => "",
  formState = {}
) => {
  return (
    <Grid item xs>
      <FormControlLabel
        label={
          <Typography style={{fontSize: 12}} className={classes.selectColor}>
            {" "}
            {input.label}
          </Typography>
        }
        control={
          <Checkbox
            checked={formState[input.key] ? true : false}
            onChange={(e) => setFormState({ [input.key]: e.target.checked })}
          />
        }
      />
    </Grid>
  );
};

const CheckBoxTransform = ({ input, setFormState, formState }) => {
  const classes = useStyles();

  return (
    <Fragment>
      <Grid item className={classes.container}>
        {CheckBoxComp(classes, input, setFormState, formState)}
      </Grid>
    </Fragment>
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

const validateNumber = (data, countryCode) => {
  const phoneNumber = parsePhoneNumberFromString(data || "123", countryCode);
  return phoneNumber
    ? phoneNumber.isValid() &&
        phoneNumber.country.toLowerCase() == countryCode.toLowerCase()
    : false;
};

const PhoneNumber = ({
  input,
  setFormState,
  formState,
  setFormStateValidation,
}) => {
  const classes = useStyles();

  input.defaultValue =
    formState[input.key] && !input.defaultValue
      ? formState[input.key]
      : input.defaultValue;

  const countryCode =
    formState[input.key]?.country?.countryCode?.toUpperCase() || "ng";

  const valid = validateNumber(
    formState[input.key]?.value || formState[input.key],
    countryCode
  );

  return (
    <Grid
      container
      style={{ marginBottom: 12 }}
      direction={input.labelDirection}
    >
      <Grid item xs={!input.labelDirection && 4}>
        <Typography className={classes.labelText}>
          {input.label}
          {input.required ? <RequiredIndicator /> : null}
        </Typography>
      </Grid>
      <Grid item xs={!input.labelDirection && 8}>
        <Grid container direction="row" item xs={12} spacing={0}>
          <Grid item xs={12}>
            <MuiPhoneNumber
              country={"ng"}
              defaultCountry="ng"
              isValid={(value, country) => {
                if (!formState[input.key]) return true;

                return valid;
              }}
              value={
                (formState[input.key] && formState[input.key].value) ||
                formState[input.key] ||
                ""
              }
              inputProps={{ style: input.style || {}, ...input.inputProps }}
              onChange={(value, country) => {
                setFormState({ [input.key]: { value, country } });
                setFormStateValidation({
                  [input.key]: validateNumber(
                    value,
                    country.countryCode?.toUpperCase()
                  ),
                });
              }}
              onBlur={input.onBlur}
              // helperText={valid ? '' : `Please Enter a valid ${countryCode} Phone  Number`}
              variant="outlined"
              fullWidth
            />
          </Grid>
        </Grid>
        {!valid && formState[input.key] ? (
          <FormHelperText style={{ marginTop: 0 }}>
            <Typography
              className={clsx(
                { [classes.errorText]: !valid },
                classes.subInputText
              )}
            >
              {valid
                ? ``
                : `Please Enter a valid ${countryCode.toUpperCase()} Phone  Number`}
            </Typography>
          </FormHelperText>
        ) : null}
      </Grid>
    </Grid>
  );
};

const LabelComp = ({ input, setFormState, formState }) => {
  const classes = useStyles();

  return (
    <Typography className={clsx(classes.subText)}>{input.label}</Typography>
  );
};

const InputSingle = ({ input, setFormState, formState }) => {
  const classes = useStyles();
  return InputTextComp(
    classes,
    input,
    setFormState,
    formState,
    () => "",
    () => "",
    null
  );
};

const SelectSingle = ({ input, setFormState, formState }) => {

  const classes = useStyles();
  let nonValid = false;
  if (
    formState[input.key] !== undefined &&
    formState[input.key] === "" &&
    input.required
  ) {
    nonValid = true;
  }
  return generateSelectType(input, classes, setFormState, formState, nonValid);
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

const SliderComp = ({ input, setFormState, formState }) => (
  <Slider
    value={input.value}
    onChange={() => ""}
    aria-labelledby="continuous-slider"
  />
);

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

    case "label":
      return (
        <LabelComp
          setFormState={setFormState}
          input={input}
          formState={formState}
        />
      );

    case "slider":
      return (
        <SliderComp
          setFormState={setFormState}
          input={input}
          formState={formState}
        />
      );

    case "grid":
      return <Grid container item xs={12}></Grid>;

    case "textDouble":
      return (
        <TextTransformDouble
          setFormState={setFormState}
          input={input}
          formState={formState}
          setFormStateValidation={setFormStateValidation}
          setPasswordMask={setPasswordMask}
          passwordMask={passwordMask}
        />
      );

    case "selectComp":
      return (
        <SelectSingle
          setFormState={setFormState}
          input={input}
          formState={formState}
        />
      );

    case "checkbox":
      return (
        <CheckBoxTransform
          setFormState={setFormState}
          input={input}
          formState={formState}
        />
      );

    case "textComp":
      return (
        <InputSingle
          setFormState={setFormState}
          input={input}
          formState={formState}
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

    case "phone":
      return (
        <PhoneNumber
          setFormState={setFormState}
          input={input}
          formState={formState}
          setFormStateValidation={setFormStateValidation}
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
