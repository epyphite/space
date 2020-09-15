import { createStore, applyMiddleware } from 'redux';
import { createLogger } from 'redux-logger';
import { composeWithDevTools } from 'redux-devtools-extension';
import thunk from 'redux-thunk';
import rootReducer from '../reducers';

export default function configureStore(initialState) {
  const logger = createLogger({
    collapsed: true
  });

  let middleware = null
  let enhancer = null
  if(process.env.NODE_ENV !== 'production') {
   middleware = [logger, thunk];
   enhancer = composeWithDevTools(applyMiddleware(...middleware));
  } else {
    middleware = [thunk];
    enhancer = applyMiddleware(...middleware)
  }

  const store = createStore(rootReducer, initialState, enhancer);

  return store;
}