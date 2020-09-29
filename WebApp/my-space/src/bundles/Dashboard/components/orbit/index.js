import React, { useState, useEffect } from "react";
import fetch from "fetch-hoc";
import { connect } from "react-redux";
import { mapOrbitData } from "bundles/utils";
import { SimpleTable } from "bundles/Dashboard/components/center";
const compose = require("lodash")?.flowRight;

const Orbit = ({ data = [] }) => {
  const orbitData = data.map((item) => ({
    label: item.Name,
    value: item.Name,
    description: item.Description,
  }));

  const [formState, setFormState] = useState({});
  const [formDescription, setFormDescription] = useState({
    value: data.length ? data[0].description : "",
  });
  const [orbitValue, setOrbitValue] = useState([]);

  useEffect(() => {
    if (data.length > 0) {
      setOrbit(data[0]);
    }
  }, [data]);

  const getInitialForm = (value) => {


    const description = data.filter((item) => item.name === value.rocket);
    setFormState({ ...formState, ...value });

    if (value.rocket) {
      if (description.length)
        setFormDescription({ value: description[0].description });

      setOrbit(description[0]);
    }
  };

  const setOrbit = (value) => {
    setOrbitValue(mapOrbitData(value, Math.floor(Math.random() * 1000)));
  };

  return (
    <SimpleTable
      data={{ content: orbitValue }}
      text={"Orbit"}
      getInitialForm={getInitialForm}
      formState={formState}
      values={data.map((item) => item.name)}
      defaultValue={
        (formState && formState["rocket"]) || (data && data[0]?.name)
      }
      description={formDescription.value}
    />
  );
};

export default compose(
  fetch("https://space.epyphite.com/launchapi/api/v1/orbit/getAll"),
  connect(null, null)
)(Orbit);
