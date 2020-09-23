import React, { useState } from "react";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import { connect } from "react-redux";
import { rocketFilter } from "bundles/Dashboard/selectors";
import { TitleText, FormBuilder } from "bundles/utils";
const compose = require("lodash")?.flowRight;

const Text = ({ text }) => {
  return (
    <Grid item>
      <Typography>{text}</Typography>
    </Grid>
  );
};

const Rocket = ({ rockets }) => {
  const rocketData = rockets.map((item) => ({
    label: item.Name,
    value: item.Name,
    description: item.Description,
  }));
  const [formState, setFormState] = useState({});
  const [formDescription, setFormDescription] = useState({
    value: rocketData.length ? rocketData[0].description : "",
  });

  const setInitialForm = (value) => {
    setFormState({ ...formState, ...value });

    const description = rockets.filter((item) => item.Name === value.rocket);
    if (description.length)
      setFormDescription({ value: description[0].Description });
  };

  return (
    <Grid container justify="center" direction="column">
      <Grid item>
        <TitleText text={"Select Rocket"} />
      </Grid>
      <Grid item>
        <FormBuilder
          formInput={{
            label: "Select",
            type: "select",
            defaultValue: rocketData.length ? rocketData[0].label : "",
            fields: rocketData,
            placeholder: "",
            labelDirection: "column",
            key: "rocket",
          }}
          formState={formState}
          setFormState={setInitialForm}
        />
      </Grid>
      <Grid item>
        <Grid container direction="row" spacing={2}>
          <Text text={formDescription.value} />
        </Grid>
      </Grid>
    </Grid>
  );
};

const mapStateToProps = (state) => ({
  rockets: rocketFilter.getRocket(state),
});

export default compose(connect(mapStateToProps, null))(Rocket);
