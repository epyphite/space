import { combineReducers } from "redux";
import rocket from "bundles/Dashboard/reducer";
const formReducers = {
  rocket,
};

const rootReducer = combineReducers({
  ...formReducers,
});

export default (state, action) =>
  rootReducer(action.type === "STAFF_LOGOUT" ? undefined : state, action);
