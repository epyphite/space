import React, { useState } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Paper from "@material-ui/core/Paper";
import { rocketFilter } from "bundles/Dashboard/selectors";
import { connect } from "react-redux";
import { TitleText, FormBuilder, mapData } from "bundles/utils";
import { Grid, Typography, Box, Button } from "@material-ui/core";
import Orbit from "bundles/Dashboard/components/orbit";
import SpacePort from "bundles/Dashboard/components/spaceport";
import Loses from "bundles/Dashboard/components/loses";
import Output from "bundles/Dashboard/components/output";
import Advanced from "bundles/Dashboard/components/advanced";
const compose = require("lodash")?.flowRight;

const useStyles = makeStyles({
  table: {
    // minWidth: 650,
  },
});

function createData(name, calories, fat, carbs, protein) {
  return { name, calories, fat, carbs, protein };
}

const rows = [
  createData("Frozen yoghurt", 159, 6.0, 24, 4.0),
  createData("Ice cream sandwich", 237, 9.0, 37, 4.3),
  createData("Eclair", 262, 16.0, 24, 6.0),
  createData("Cupcake", 305, 3.7, 67, 4.3),
  createData("Gingerbread", 356, 16.0, 49, 3.9),
];

const loses = {
  leftTitle: "Loses",
  rightTitle: "",
  content: [
    {
      name: "rocket space",
      left: {
        label: "1st stage Fuel & Cycle",
        placeholder: "Enter your email",
        type: "label",
      },
      right: {
        label: "",
        fields: ["A", "B", "C"],
        type: "selectComp",
      },
      middle: {
        label: "",
        fields: ["A", "B", "C"],
        type: "selectComp",
      },
      symbol: {
        label: "",
      },
      center: {
        type: "grid",
      },
    },
    {
      name: "rocket data",
      left: {
        label: "2nd stage Fuel & Cycle ",
        placeholder: "Enter your email",
        type: "label",
      },
      middle: {
        label: "1st",
        type: "textDouble",
      },
      right: {
        label: "",
        fields: ["A", "B", "C"],
        type: "selectComp",
      },
      symbol: {
        label: "",
      },
      center: {
        type: "grid",
      },
    },
  ],
};

export const SimpleTable = ({
  data = { title: "", content: "" },
  getInitialForm = () => "",
  formState = {},
  text,
  values,
  defaultValue = "",
  description,
}) => {
  const classes = useStyles();

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>
              <Grid container>
                <Grid item xs>
                  <TitleText text={data.leftTitle || text} />
                </Grid>

                {values && (
                  <Grid item xs>
                    <FormBuilder
                      formInput={{
                        label: "Select",
                        type: "select",
                        defaultValue,
                        fields: values,
                        placeholder: "",
                        labelDirection: "column",
                        key: "rocket",
                      }}
                      formState={formState}
                      setFormState={getInitialForm}
                    />
                  </Grid>
                )}
                <Grid item xs={12}>
                  {description}
                </Grid>
              </Grid>
            </TableCell>
            <TableCell align="right">
              <Grid container>
                <Grid item xs>
                  <Typography> {data.rightTitle} </Typography>
                </Grid>
              </Grid>
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {data?.content?.map((row) => (
            <TableRow key={row.name}>
              <TableCell component="th" style={{ width: "50%" }} scope="row">
                <Grid container>
                  <Grid item md={12} >
                    <Grid
                      container
                      justify="space-between"
                      alignContent="center"
                      alignItems="center"
                    >
                      <Grid item md={7}>
                        <FormBuilder
                          formInput={{ ...row.left }}
                          setFormState={getInitialForm}
                          formState={formState}
                        />
                      </Grid>
                      <Grid item md={5}>
                        <FormBuilder
                          formInput={{ ...row.right }}
                          setFormState={getInitialForm}
                          formState={formState}
                        />
                      </Grid>
                    </Grid>
                  </Grid>
                </Grid>
              </TableCell>
              <TableCell component="th" scope="row">
                <Grid container>
                  <Grid item md={12}>
                    <Grid container alignContent="center" alignItems="center">
                      <Grid item md={5} style={{ paddingRight: 5 }}>
                        <FormBuilder
                          formInput={{ ...row.center }}
                          setFormState={getInitialForm}
                          formState={formState}
                        />
                      </Grid>
                      <Grid item md={6} style={{ textAlign: "right" }}>
                        <FormBuilder
                          formInput={{ ...row.middle }}
                          setFormState={getInitialForm}
                          formState={formState}
                        />
                      </Grid>
                      <Grid item md={1}>
                        <Typography style={{ textAlign: "center" }}>
                          {row?.symbol?.label}
                        </Typography>
                      </Grid>
                    </Grid>
                  </Grid>
                </Grid>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

const ComplexTable = ({ data = [[]] }) => {
  const classes = useStyles();

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>
              <Grid container>
                <Grid item xs>
                  Space Rocket
                </Grid>
              </Grid>
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {data.map((row) => (
            <TableRow key={row.name}>
              <TableCell component="th" scope="row">
                <Grid container>
                  <Grid item xs={12}>
                    <Grid container justify="space-between">
                      <Typography> Hello</Typography>
                      <Typography> World</Typography>
                    </Grid>
                  </Grid>
                </Grid>
              </TableCell>
            </TableRow>
          ))}
          {rows.map((row) => (
            <TableRow key={row.name}>
              <TableCell component="th" scope="row">
                <Grid container>
                  <Grid item xs={12}>
                    <Grid container justify="space-between">
                      <Typography> Hello</Typography>
                      <Typography> World</Typography>
                    </Grid>
                  </Grid>
                </Grid>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

const WrapComp = ({ rocket }) => {
  const [formState, setFormState] = useState({});

  const getInitialForm = (value) => {
    setFormState({ ...formState, ...value });
  };

  return (
    <>
     <Advanced />
    <Grid container spacing={4} style={{ padding: '1%' }}>
      <Grid item xs={12} md={7}>
        <Grid container>
          <Grid item xs={12} md={12}>
            <SimpleTable
              text={"Rocket"}
              data={{ content: rocket }}
              getInitialForm={getInitialForm}
              formState={formState}
            />
          </Grid>
        </Grid>
      </Grid>
      <Grid item xs={12} md={5}>
        <Grid container spacing={2} direction="column">
          <Grid item xs={12} md={12}>
            <SpacePort />
            <div style={{ marginTop: 15 }}>
              <Orbit />
            </div>
            <div style={{ marginTop: 15 }}>
              <Loses />
            </div>
            <div style={{ marginTop: 15 }}>
              <Output />
            </div>
          </Grid>
          {/* <Grid  item md={12}>
            <Orbit />
          </Grid>
          <Grid item md={12}>
            <Loses />
          </Grid>
          <Grid item md={12}>
            <Output />
          </Grid> */}
        </Grid>
      </Grid>
     
    </Grid>
    </>
  );
};

const CenterDashboard = ({ rocket }) => {
  return <WrapComp rocket={rocket} />;
};

const mapStateToProps = (state) => ({
  rocket: rocketFilter.getSingleRocket(state),
});

export default compose(connect(mapStateToProps, null))(CenterDashboard);
