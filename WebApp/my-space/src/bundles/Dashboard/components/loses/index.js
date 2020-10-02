import React, { useState, useEffect } from "react";
import fetch from "fetch-hoc";
import { connect } from "react-redux";
import { mapLosesSpaceData } from "bundles/utils";
import { SimpleTable } from "bundles/Dashboard/components/center";
const compose = require("lodash")?.flowRight;


const Loses = ({ data = [{
    gravityloses: 122,
    aerodynamicloses: 12233,
    assummedloses: 1.2333,
    requiredeltaloses: 122
}] }) => {
  const losesData = data.map((item) => ({
    label: item.Name,
    value: item.Name,
    description: item.Description,
  }));

  const [formState, setFormState] = useState({});
  const [formDescription, setFormDescription] = useState({
    value: data.length ? data[0].description : "",
  });
  const [losesValue, setlosesValue] = useState([]);

  useEffect(() => {
    if (data.length > 0) {
      setloses(data[0]);
    }
  }, []);

  const getInitialForm = (value) => {


    const description = data.filter((item) => item.name === value.rocket);
    setFormState({ ...formState, ...value });

    if (value.rocket) {
      if (description.length)
        setFormDescription({ value: description[0].description });

      setloses(description[0]);
    }
  };

  const setloses = (value) => {
    setlosesValue(mapLosesSpaceData(value, Math.floor(Math.random() * 1000)));
  };

  return (
    <SimpleTable
      data={{ content: losesValue }}
      text={"Loses"}
      getInitialForm={getInitialForm}
      formState={formState}
      
      description={formDescription.value}
    />
  );
};

export default compose(
  connect(null, null)
)(Loses);
