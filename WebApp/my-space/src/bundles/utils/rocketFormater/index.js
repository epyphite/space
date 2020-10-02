
import _ from 'lodash';

const buildTextTitleObjects = (label, title, defaultValue, key, symbol) => {
  return {
    key,
    left: {
      label,
      title,
      type: "label",
      key,
    },
    right: {
      type: "grid",
    },
    middle: {
      label: "",
      type: "textComp",
      defaultValue,
      key: `${key}-${label}`,
      style: {
        textAlign: 'right'
    }
    },
    center: {
      type: "grid",
    },
    symbol: {
      label: symbol,
    },
  };
};

const buildLabelObject = (label, title, defaultValue, key, symbol) => {
  return  {
    key,
    left: {
      label,
      type: "label",
      title,
      key,
    },
    right: {
      type: "grid",
    },
    middle: {
        label: defaultValue,
        type: "label",
        key,
        style: {
          paddingRight: 5, color: 'black'
      }
    },
    center: {
      type: "grid",
    },
    symbol: {
      label: symbol,
      style: {
       
    }
    },
  };
}

const buildSliderObjects = (label, title, defaultValue, key, symbol) => {
  return  {
    key,
    left: {
      label,
      type: "label",
      title,
      key,
    },
    right: {
      type: "grid",
    },
    middle: {
        label: "",
        type: "textComp",
        defaultValue,
        key: `${key}-${label}`,
        style: {
            textAlign: 'right'
        }
    },
    center: {
      type: "slider",
    },
    symbol: {
      label: symbol,
    },
  };
}

const buildDropFuelStagesObjects = (label, title, defaultValue, key, symbol) => {
    return {
      key,
      left: {
        label,
        placeholder: "Enter your email",
        type: "label",
        title,
        key,
      },
      right: {
        label: "",
        fields: ["Staged Combustion", "Gas-Generator", "Electric", "Pressure Fed"],
        type: "selectComp",
        key: `${key}-${label}`,
      },
      middle: {
        label: "",
        fields: ["Lox/Kerosene", "Lox/Methane", "Lox/LH2", "Monopropellant", "Solid"],
        type: "selectComp",
        key: `${key}-${label}-grid`,
      },
      symbol: {
        label: "",
      },
      center: {
        type: "grid",
      },
    }
}

const fuelFields = ["Staged Combustion", "Gas-Generator", "Electric", "Pressure Fed"]

let secondSetObject = (defaultValue, key, keyPair) => {

  
    if (defaultValue == fuelFields.length) return null

    return {
        middle: {
            label: "",
            fields: fuelFields,
            defaultValue: fuelFields[defaultValue],
            type: "selectComp",
            key: `${keyPair}-select`,
        }
    }

}

const cycleFields = ["Lox/Kerosene", "Lox/Methane", "Lox/LH2", "Monopropellant", "Solid"]

let firstSetObject = (label, title, defaultValue, key, symbol, keyPair) => {

    return {
        key,
        right: {
            label,
            fields: cycleFields,
            type: "selectComp",
            defaultValue: cycleFields[defaultValue],
            title,
            key: keyPair,
          },
          symbol: {
            label: "",
          },
          center: {
            type: "grid",
          },
          left: {
            label,
            placeholder: "Enter your email",
            type: "label",
            title,
            key,
          },
    }
}

let firstStage = {}


let secondStage = {}

let thirdStage = {}

const parseKeyType = (key, value, rocketName) => {

  const keyName = `${key}-${rocketName}`
   

  switch (key) {
    case "rocketmass":
      return buildTextTitleObjects("Rocket mass (Lift off excluding Payload and Fairing)", "The initial mass of a rocket excluding payload and fairing", value, keyName , "kg")
    case "maxrocketbodydiameter":
      return buildTextTitleObjects("Max Rocket Body Diameter", "", value, keyName, "m");

    case "fairingmass":
        return buildTextTitleObjects("Fairing mass", "", value, keyName, "kg");

    case "assumedpayloadmass": 
        return buildTextTitleObjects("Assumed payload mass", "", value, keyName, "kg");

    case "secondstagetorocketmassratio":
        return buildTextTitleObjects("2nd stage to Rocket mass ratio", "", value, keyName, "%");
    
    case "firststagedrytowetmassratio": return buildSliderObjects("1st stage Dry to Wet mass ratio", "", value, keyName, "%");

    case "secondstagedrytowetmassratio": return buildSliderObjects("2nd stage Dry to Wet mass ratio", "", value, keyName, "%");

    case "unusedpropellantoffirststage": return buildSliderObjects("Unused propellant of 1st stage", "", value, keyName, "%");

    case "unusedpropellantofseconstage": return buildSliderObjects("Unused propellant of 2nd stage", "", value, keyName, "%");

    case "firststageispstartaltitude": return buildSliderObjects("1st stage Isp sea level or at the start altitude", "", value, keyName, "%");

    case "firststageispvacuum": return buildSliderObjects("1st stage Isp vacuum", "", value, keyName, "%");

    case "secondstageisp": return buildSliderObjects("2nd stage Isp vacuum", "", value, keyName, "%");

    case "transferorbitstagetorocketmassratio": return buildTextTitleObjects("Transfer Orbit stage to Rocket mass ratio", "", value, keyName, "%");

    

    case "firststagefuel": {

            firstStage = {}
            firstStage = { 
                ...firstSetObject("1st stage Fuel & Cycle", "", value, key, '', keyName)
            }
    }

    case "firststagecycle": 

    return {
        ...firstStage, ...secondSetObject(value, "firststagecycle", keyName)
    }
        
    
    case "secondstagefuel": {

        secondStage = {}
            secondStage = { 
                ...firstSetObject("2nd stage Fuel & Cycle", "", value, key, '', keyName)
            }
    }

    case "secondstagecycle":    {

        return {
            ...secondStage, ...secondSetObject(value, "secondstagecycle", keyName)
        }

    }


    case "thirdstagefuel":  {
        thirdStage = {}
            thirdStage = { 
                ...firstSetObject("3rd stage Fuel & Cycle", "", value, key, "", keyName)
            }
    }

    case "thirdstagecycle": return {
        ...thirdStage, 
        ...secondSetObject(value, "thirdstagecycle", keyName)
    }

    default:
        return null
  }
};

const parseData = (key, value) => {};

const generateOrder = (key) => {
    switch (key) {
        case "rocketmass":
          return 7
        case "maxrocketbodydiameter":
          return 8
    
        case "fairingmass":
            return 9
    
        case "assumedpayloadmass": 
            return 10
    
        case "secondstagetorocketmassratio":
            return 11
        
        case "firststagedrytowetmassratio": return 12
    
        case "secondstagedrytowetmassratio": return 13
    
        case "unusedpropellantoffirststage": return 14
    
        case "unusedpropellantofseconstage": return 15
    
        case "firststageispstartaltitude": return 16
    
        case "firststageispvacuum": return 17
    
        case "secondstageisp": return 18
    
        case "transferorbitstagetorocketmassratio": return 
    
        case "firststagefuel": 
            return 1;
        

        case "firststagecycle": return 2;

        case "secondstagefuel": return 3;

        case "secondstagecycle": return 4;

        case "thirdstagefuel": return 5;

        case "thirdstagecycle": return 6;
    
    
      }
}


const generateOrbitOrder = (key) => {
  switch (key) {
    case "orbitperigee": return 1;

    case "orbitapogee": return 2;

    case "orbitinclination": return 3;

    case "orbitalperiod": return 4;

    case "deltav": return 5;

    case "extravelocity": return 6;


  }
}


const parseObitKeyType = (key, value, rocketName) => {
  const keyName = `${key}-${rocketName}`
  switch (key) {
    case "orbitperigee": return buildTextTitleObjects("Orbit Perigee", "", value, keyName, "km");

    case "orbitapogee": return buildTextTitleObjects("Orbit Apogee (0 means a circular orbit)", "", value, keyName, "km");

    case "orbitinclination": return buildTextTitleObjects("Orbit Inclination", "", value, keyName, "deg");

    case "orbitalperiod": return buildLabelObject("Orbit Period", "", value, keyName, "min");

    case "deltav": return buildLabelObject("Delta-v for target orbit (ideal)", "", value, keyName, "m/s");

    case "extravelocity": return buildTextTitleObjects("Extra velocity for flight to the planets", "", value, keyName, "m/s");

  }
  
}

const parseSpacePortType = (key, value, portName) => {
  const keyName = `${key}-${portName}`;

  switch(key) {
    case "launchpointaltitude": return buildTextTitleObjects("Launch point Altitude", "". keyName, value, "km");

    case "additionalvelocity": return buildTextTitleObjects("Additional velocity (Air Launch)", "". keyName, value, "m/s");

    case "spaceportlatitude": return buildTextTitleObjects("Spaceport Latitude", "",value, keyName, "deg");
    // case "spaceportlongitude": -104.76
     case "launchazimuth": return buildLabelObject("Launch Azimuth", "", value, keyName, "deg");
     case "earthrotationvelocity": return buildLabelObject("Earth rotation velocity", "", value, keyName, "m/s");
    // case "auxiliaryangle": 0
    // case "launchpointaltitudeorbitalvelocity": 0
    // case "absoluteorbitalvelocity": 0
  }
}

const generateSpacePortOrder = (key) => {
  switch (key) {
    case "launchpointaltitude": return 1;

    case "additionalvelocity": return 2;

    case "spaceportlatitude": return 3;

    case "launchazimuth": return 4;

    case "earthrotationvelocity": return 5;

  }
}

const lossPortType = (key, value, losesName) => {
  const keyName = `${key}-${losesName}`;

  switch(key) {
    case "gravityloses": return buildLabelObject("Gravity Losses", "", value, keyName, "m/s");

    case "aerodynamicloses": return buildLabelObject("Aerodynamic Losses", "", value, keyName, "m/s");

    case "assummedloses": return buildLabelObject("Assumed Losses", "", value, keyName, "%");

    case "requiredeltaloses": return buildLabelObject("Required delta-v with Losses", "", value, keyName, "m/s");
  }
}

const generateLosesPortOrder = (key) => {
  switch(key) {
    case "gravityloses": return 1;

    case "aerodynamicloses": return 2;

    case "assummedloses": return 3;

    case "requiredeltaloses": return 4;

  }
}

const reorderOrbitRender =(orbit) => {
  const orbitOrder = orbit.map((data) => ({...data, order: generateOrbitOrder(data.key)}));
  const ordered = _.orderBy(orbitOrder, ['order']);

  return _.uniqBy(ordered, 'key');
}

export const mapOrbitData = (orbits = {}, orbitType) => {
  const dataKeys = Object.keys(orbits);

  const mappedData = dataKeys.map((key) => parseObitKeyType(key, orbits[key], orbitType)).filter((item) => item);

  return reorderOrbitRender(mappedData);
}

const reorderRocket = (rockets) => {
  const rocketOrder = rockets.map((data) => ({...data, order: generateOrder(data.key)}));
  const ordered = _.orderBy(rocketOrder, ['order']);

  return _.uniqBy(ordered, 'key');
}

export const mapData = (rockets = {}, rocketType) => {
  const dataKeys = Object.keys(rockets);
  const mappedData = dataKeys.map((key) => parseKeyType(key, rockets[key], rocketType)).filter((item) => item);

  return  reorderRocket(mappedData)
};

const reOrderSpace = (space) => {
  const spacePortOrder = space.map((data) => ({...data, order: generateSpacePortOrder(data.key)}));
  const ordered = _.orderBy(spacePortOrder, ['order']);

  return _.uniqBy(ordered, 'key');
}

export const mapSpacePortData = (space ={}, spaceportType) => {
  const dataKeys = Object.keys(space);
  const mappedData = dataKeys.map((key) => parseSpacePortType(key, space[key], spaceportType)).filter((item) => item);

  return reOrderSpace(mappedData);
}

const reOrderLoses = (loses) => {
  const losesPortOrder = loses.map((data) => ({...data, order: generateLosesPortOrder(data.key)}));
  const ordered = _.orderBy(losesPortOrder, ['order']);

  return _.uniqBy(ordered, 'key');
}

export const mapLosesSpaceData = (loses = {}, losesPortType) => {
  const dataKeys = Object.keys(loses);
  const mappedData = dataKeys.map((key) => lossPortType(key, loses[key], losesPortType)).filter((item) => item);

  return reOrderLoses(mappedData)
}

