import React, { useState, useEffect } from "react";
import fetch from "fetch-hoc";
import { connect } from "react-redux";
import { mapSpacePortData } from "bundles/utils";
import { SimpleTable } from "bundles/Dashboard/components/center";
const compose = require("lodash")?.flowRight;

const SpacePort = ({ data = [{}] }) => {
  const spacePortData = data.map((item) => ({
    label: item.Name,
    value: item.Name,
    description: item.Description,
  }));

  const [formState, setFormState] = useState({});
  const [formDescription, setFormDescription] = useState({
    value: data.length ? data[0].description : "",
  });
  const [spacePortValue, setspacePortValue] = useState([]);

  useEffect(() => {
    if (data.length > 0) {
      setspacePort(data[0]);
    }
  }, [data]);

  const getInitialForm = (value) => {


    const description = data.filter((item) => item.name === value.rocket);
    setFormState({ ...formState, ...value });

    if (value.rocket) {
      if (description.length)
        setFormDescription({ value: description[0].description });

      setspacePort(description[0]);
    }
  };

  const setspacePort = (value) => {
    setspacePortValue(mapSpacePortData(value, Math.floor(Math.random() * 1000)));
  };

  return (
    <SimpleTable
      data={{ content: spacePortValue }}
      text={"spacePort"}
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
  fetch(`${process.env.REACT_APP_URL}/launchapi/api/v1/spaceport/getAll`),
  connect(null, null)
)(SpacePort);