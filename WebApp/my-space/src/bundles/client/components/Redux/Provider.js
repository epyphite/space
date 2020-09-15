import * as React from 'react';
import{ Provider as ReduxProvider } from 'react-redux';
import configureStore from 'store/configureStore';

export const store = configureStore()

const Provider = (props) => <ReduxProvider store={store} {...props} />

export default Provider;