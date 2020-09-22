export const isValidState = (data, setState, state) => {
  let nonValidState = {};
  //const formQuestions = data.reduce((agg, val) => [...agg, ...val?.content], []);

  data.forEach((field) => {
    if (
      (field.required && !state[field.key]) ||
      (field.required && state[field.key] && state[field.key].length < 1)
    ) {
      nonValidState[field.key] = "";
    }
  });

  if (Object.keys(nonValidState).length > 0) {
    setState({ ...state, ...nonValidState });

    return false;
  }

  return true;
};
