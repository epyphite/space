import { actionTypes } from 'bundles/Dashboard/constants';

export const setRocket = value => ({
    type: actionTypes.SET_ROCKET,
    payload: { value }
});

export const setCurrentRocket =  value => ({
    type: actionTypes.ADD_ROCKET,
    payload: { value }
});