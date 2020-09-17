import React from "react";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import { TitleText, FormBuilder } from "bundles/utils";

const Text = ({ text }) => {
  return (
    <Grid item>
      <Typography>{text}</Typography>
    </Grid>
  );
};

const Rocket = () => {
  return (
    <Grid container justify="center"   direction="column">
      <Grid item>
        <TitleText text={"Select Rocket"} />
      </Grid>
      <Grid item>
        <FormBuilder
          formInput={{
            label: "Select",
            type: "select",
            defaultValue: "",
            fields: [],
            placeholder: "",
            labelDirection: "column",
            key: "age",
          }}
          formState={{}}
          setFormState={() => ""}
        />
      </Grid>
      <Grid item>
        <Grid container direction="row" spacing={2}>
          <Text text={"Space X"} />
        </Grid>
      </Grid>
    </Grid>
  );
};

export default Rocket;
