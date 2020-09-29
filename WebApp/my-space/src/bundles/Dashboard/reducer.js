import Immutable from 'immutable';
import { actionTypes } from './constants';
import { mapData } from "bundles/utils";

const initialRocketFilters = Immutable.Map({
  rocket: Immutable.List([]),
});  

const initialState = Immutable.Map({
  rocketFilters: initialRocketFilters,
  currentRocket: []
}); 

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case actionTypes.SEARCH_FILTERS_CLEAR_ALL:
      return handleClearAllSearch(state, action);

    case actionTypes.ADD_ROCKET:
      return handleAddRocket(state, action); 

    default:
      return state;
  }
}; 

const handleAddRocket = (state, action) => {
  const { value } = action.payload;
  return state.set('currentRocket', mapData(value[0], Math.floor(Math.random() * 1000)));
}

const handleClearAllSearch = (state, action) => {
  return state.set('rocketFilters', initialRocketFilters);
};

export default reducer;
