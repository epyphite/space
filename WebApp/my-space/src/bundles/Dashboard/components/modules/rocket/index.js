import React, { useState, useEffect } from "react";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import { connect } from "react-redux";
import fetch from 'fetch-hoc';
import { setRocket, setCurrentRocket } from "bundles/Dashboard/actions";
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

const Rocket = ({ data: rockets = [], setRocket, setCurrentRocket, ...props }) => {
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


    setRocketValue(description)
  };

  const setRocketValue = (value) => {
    setCurrentRocket(value);
  }


  useEffect(() => {
    if(rockets.length > 0)  setCurrentRocket(rockets);
  }, [rockets])

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


const mapDispatchToProps = dispatch => ({
  setCurrentRocket: value => dispatch(setCurrentRocket(value))
});



export default compose(fetch('https://space.epyphite.com/launchapi/api/v1/rocket/getAll'),connect(null, mapDispatchToProps))(Rocket);
