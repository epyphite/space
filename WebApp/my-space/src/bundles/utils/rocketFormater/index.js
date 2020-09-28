
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

let secondSetObject = (defaultValue, key) => {
  
    return {
        middle: {
            label: "",
            fields: fuelFields,
            defaultValue: defaultValue ? fuelFields[defaultValue] : '',
            type: "selectComp",
            key: `${key}-select`,
        }
    }

}

const cycleFields = ["Lox/Kerosene", "Lox/Methane", "Lox/LH2", "Monopropellant", "Solid"]

let firstSetObject = (label, title, defaultValue, key, symbol) => {
    return {
        key,
        right: {
            label,
            fields: cycleFields,
            type: "selectComp",
            defaultValue: defaultValue ? cycleFields[defaultValue]: '',
            title,
            key,
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

const parseKeyType = (key, value) => {

   

  switch (key) {
    case "rocketmass":
      return buildTextTitleObjects("Rocket mass (Lift off excluding Payload and Fairing)", "The initial mass of a rocket excluding payload and fairing", value, key, "kg")
    case "maxrocketbodydiameter":
      return buildTextTitleObjects("Max Rocket Body Diameter", "", value, key, "m");

    case "fairingmass":
        return buildTextTitleObjects("Fairing mass", "", value, key, "kg");

    case "assumedpayloadmass": 
        return buildTextTitleObjects("Assumed payload mass", "", value, key, "kg");

    case "secondstagetorocketmassratio":
        return buildTextTitleObjects("2nd stage to Rocket mass ratio", "", value, key, "%");
    
    case "firststagedrytowetmassratio": return buildSliderObjects("1st stage Dry to Wet mass ratio", "", value, key, "%");

    case "secondstagedrytowetmassratio": return buildSliderObjects("2nd stage Dry to Wet mass ratio", "", value, key, "%");

    case "unusedpropellantoffirststage": return buildSliderObjects("Unused propellant of 1st stage", "", value, key, "%");

    case "unusedpropellantofseconstage": return buildSliderObjects("Unused propellant of 2nd stage", "", value, key, "%");

    case "firststageispstartaltitude": return buildSliderObjects("1st stage Isp sea level or at the start altitude", "", value, key, "%");

    case "firststageispvacuum": return buildSliderObjects("1st stage Isp vacuum", "", value, key, "%");

    case "secondstageisp": return buildSliderObjects("2nd stage Isp vacuum", "", value, key, "%");

    case "transferorbitstagetorocketmassratio": return buildTextTitleObjects("Transfer Orbit stage to Rocket mass ratio", "", value, key, "%");

    

    case "firststagefuel": {

            firstStage = {}
            firstStage = { 
                ...firstSetObject("1st stage Fuel & Cycle", "", value, key, "")
            }
    }

    case "firststagecycle": 

    return {
        ...firstStage, ...secondSetObject(value, "firststagecycle")
    }
        
    
    case "secondstagefuel": {

        secondStage = {}
            secondStage = { 
                ...firstSetObject("2nd stage Fuel & Cycle", "", value, key, "")
            }
    }

    case "secondstagecycle":    {

        return {
            ...secondStage, ...secondSetObject(value, "secondstagecycle")
        }

    }


    case "thirdstagefuel":  {
        thirdStage = {}
            thirdStage = { 
                ...firstSetObject("3rd stage Fuel & Cycle", "", value, key, "")
            }
    }

    case "thirdstagecycle": return {
        ...thirdStage, 
        ...secondSetObject(value, "thirdstagecycle")
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

const reorderRocket = (rockets) => {
  const rocketOrder = rockets.map((data) => ({...data, order: generateOrder(data.key)}));
  const ordered = _.orderBy(rocketOrder, ['order']);

  return _.uniqBy(ordered, 'key');
}

export const mapData = (rockets = {
    "ID": "0626e09c-64a8-4aa2-b516-e602d8fd7f68",
    "Name": "Blue Origin New Glenn 3st",
    "Description": "Blue Origin 3-Stages (Superceded) (project)",
    "thrusttoweightratioone": 1.2,
    "thrusttoweightratio": 0.83,
    "rocketmass": 1435000,
    "maxrocketbodydiameter": 7,
    "fairingmass": 4000,
    "fairingjettisonvelocity": 3500,
    "jettisonedbattery": 0,
    "assumedpayloadmass": 20000,
    "secondstagetorocketmassratio": 22.647,
    "transferorbitstagetorocketmassratio": 5.761,
    "firststagedrytowetmassratio": 9.91,
    "secondstagedrytowetmassratio": 7.037,
    "transferorbitstagedrytowetmassratio": 10.909,
    "unusedpropellantoffirststage": 1,
    "unusedpropellantofsecondstage": 0,
    "unusedpropellantoftransferorbitstage": 1,
    "firststageispstartaltitude": 310,
    "firststageispvacuum": 335,
    "secondstageisp": 358,
    "transferorbitstageisp": 440,
    "firststagefuel": 1,
    "firststagecycle": 0,
    "secondstagefuel": 1,
    "secondstagecycle": 0,
    "thirdstagefuel": 2,
    "thirdstagecycle": 1}) => {
  const dataKeys = Object.keys(rockets);
  const mappedData = dataKeys.map((key) => parseKeyType(key, rockets[key])).filter((item) => item);

  return  reorderRocket(mappedData)


};

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
