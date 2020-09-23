import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Paper from "@material-ui/core/Paper";
import { TitleText, FormBuilder, mapData } from "bundles/utils";
import { Grid, Typography } from "@material-ui/core";

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

/*
[
 {
   "name": "Custom Rocket...",
   "thrusttoweightratioone": 1.2,
   "thrusttoweightratio": 1.1,
   "rocketmass": 32475,
   "maxrocketbodydiameter": 1.7,
   "fairingmass": 150,
   "fairingjettisonvelocity": 3500,
   "jettisonedbattery": 0,
   "assumedpayloadmass": 775,
   "secondstagetorocketmassratio": 24.86,
   "transferorbitstagetorocketmassratio": 0,
   "firststagedrytowetmassratio": 8,
   "secondstagedrytowetmassratio": 10,
   "transferorbitstagedrytowetmassratio": 20,
   "unusedpropellantoffirststage": 2,
   "unusedpropellantofseconstage": 2,
   "unusedpropellantoftransferorbitstage": 1,
   "firststageispstartaltitude": 290.5,
   "firststageispvacuum": 327,
   "secondstageisp": 365,
   "transferorbitstageisp": 320,
   "firststagefuel": 0,
   "firststagecycle": 0,
   "secondstagefuel": 0,
   "secondstagecycle": 0,
   "thirdstagefuel": 0,
   "thirdstagecycle": 3,
   "spaceport": "Cape Canaveral",
   "orbit": "LEO",
   "airlaunch": 0,
   "description": "Designed Rocket - engine based",
 },
 {
   "name": "NASA Saturn V",
   "thrusttoweightratioone": 1.165,
   "thrusttoweightratio": 0.78,
   "rocketmass": 2909200,
   "maxrocketbodydiameter": 10.1,
   "fairingmass": 8000,
   "fairingjettisonvelocity": 3500,
   "jettisonedbattery": 0,
   "assumedpayloadmass": 48600,
   "secondstagetorocketmassratio": 21.284,
   "transferorbitstagetorocketmassratio": 4.228,
   "firststagedrytowetmassratio": 5.677,
   "secondstagedrytowetmassratio": 8.081,
   "transferorbitstagedrytowetmassratio": 10.976,
   "unusedpropellantoffirststage": 2.6,
   "unusedpropellantofseconstage": 2,
   "unusedpropellantoftransferorbitstage": 2,
   "firststageispstartaltitude": 263,
   "firststageispvacuum": 304,
   "secondstageisp": 421,
   "transferorbitstageisp": 421,
   "firststagefuel": 0,
   "firststagecycle": 1,
   "secondstagefuel": 2,
   "secondstagecycle": 1,
   "thirdstagefuel": 2,
   "thirdstagecycle": 1,
   "spaceport": "Cape Canaveral",
   "orbit": "Moon",
   "airlaunch": 0,
   "description": "Apollo lunar program launcher",
 }
]
*/



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

const orbit = {
  leftTitle: "Orbit",
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
        type: "grid",
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
    {
      name: "rocket space test",
      left: {
        label: "1st stage Fuel & Cycle",
        placeholder: "Enter your email",
        type: "label",
      },
      center: {
        type: "grid",
      },
      right: {
        label: "",
        fields: ["Gas Generator", "B", "C"],
        type: "selectComp",
      },
      middle: {
        type: "label",
        label: "370",
      },
      symbol: {
        label: "%",
      },
    },
  ],
};

const spaceport = {
  leftTitle: "Space Port",
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
    {
      name: "rocket space test",
      left: {
        label: "1st stage Fuel & Cycle",
        placeholder: "Enter your email",
        type: "label",
      },
      center: {
        type: "grid",
      },
      right: {
        label: "",
        fields: ["Gas Generator", "B", "C"],
        type: "selectComp",
      },
      middle: {
        type: "label",
        label: "370",
      },
      symbol: {
        label: "%",
      },
    },
  ],
};

const rockedData = {
  leftTitle: "Rocket",
  rightTitle: "Fixed Design",
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
    {
      name: "rocket space test",
      left: {
        label: "1st stage Fuel & Cycle",
        placeholder: "Enter your email",
        type: "label",
      },
      center: {
        type: "grid",
      },
      right: {
        label: "",
        fields: ["Gas Generator", "B", "C"],
        type: "selectComp",
      },
      middle: {
        type: "label",
        label: "370",
      },
      symbol: {
        label: "%",
      },
    },
    {
      name: "rocket space thrust",
      left: {
        label: "Thrust to weight ration",
        placeholder: "Enter your email",
        type: "label",
      },
      right: {
        type: "grid",
      },
      middle: {
        label: "",
        fields: ["A", "B", "C"],
        type: "selectComp",
      },
      center: {
        type: "grid",
      },
      symbol: {
        label: "",
      },
    },
    {
      name: "rocket space machine",
      left: {
        label: "Stages",
        placeholder: "",
        type: "label",
      },
      right: {
        type: "grid",
      },
      middle: {
        label: "Limit to reasonable",
        type: "checkbox",
      },
      center: {
        type: "grid",
      },
      symbol: {
        label: "",
      },
    },
    {
      name: "wet ratio",
      left: {
        label: "1st stage Dry to Wet mass ratio",
        type: "label",
      },
      right: {
        type: "grid",
      },
      middle: {
        label: "",
        fields: ["A", "B", "C"],
        type: "selectComp",
      },
      center: {
        type: "slider",
      },
      symbol: {
        label: "kg",
      },
    },
    {
      name: "wet ratio days",
      left: {
        label: "2nd stage Dry to Wet mass ratio",
        type: "label",
      },
      right: {
        type: "grid",
      },
      middle: {
        label: "",
        type: "textComp",
      },
      center: {
        type: "slider",
      },
      symbol: {
        label: "kg",
      },
    },
  ],
};

const rocketDataSemi = [
  {
    left: {
      label: "1st stage Fuel & Cycle",
      placeholder: "",
      type: "label",
    },
    right: {
      label: "",
      fields: ["A", "B", "C"],
      type: "selectComp",
    },
  },
];

const SimpleTable = ({ data = { title: "", content: "" } }) => {
  const classes = useStyles();


  console.log(data, '---j--- >  --- >')

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>
              <Grid container>
                <Grid item xs>
                  <TitleText text={data.leftTitle} />
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
                  <Grid item md={12}>
                    <Grid
                      container
                      justify="space-between"
                      alignContent="center"
                      alignItems="center"
                    >
                      <Grid item md={7}>
                        <FormBuilder
                          formInput={{ ...row.left }}
                          setFormState={() => ""}
                          formState={{}}
                        />
                      </Grid>
                      <Grid item md={5}>
                        <FormBuilder
                          formInput={{ ...row.right }}
                          setFormState={() => ""}
                          formState={{}}
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
                          setFormState={() => ""}
                          formState={{}}
                        />
                      </Grid>
                      <Grid item md={6} style={{ textAlign: "right" }}>
                        <FormBuilder
                          formInput={{ ...row.middle }}
                          setFormState={() => ""}
                          formState={{}}
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

const CenterDashboard = () => {

  const dataStore = {
    content: mapData()
  }

  console.log(dataStore, '---> --> ');
  
  return (
    <Grid container spacing={4} style={{ padding: 20 }}>
      <Grid item xs={12} md={7}>
        <Grid container>
          <Grid item md={12}>
            <SimpleTable data={{ content: mapData()}} />
          </Grid>
        </Grid>
      </Grid>
      <Grid item xs={12} md={5}>
        <Grid container spacing={2} direction="column">
          <Grid item md={12}>
            <SimpleTable data={spaceport} />
          </Grid>
          <Grid item md={12}>
            <SimpleTable data={orbit} />
          </Grid>
          <Grid item md={12}>
            <SimpleTable data={loses} />
          </Grid>
          <Grid item md={12}>
            <SimpleTable data={loses} />
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
};

export default CenterDashboard;
