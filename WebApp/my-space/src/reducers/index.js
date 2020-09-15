import { combineReducers } from "redux";
const formReducers = {};

const rootReducer = combineReducers({
  ...formReducers,
});

export default (state, action) =>
  rootReducer(action.type === "STAFF_LOGOUT" ? undefined : state, action);
