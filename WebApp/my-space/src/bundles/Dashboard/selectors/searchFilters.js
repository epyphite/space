import { getRocketState } from './state';

export const getRocket = (state) => getRocketState(state).getIn(['rocketFilters', 'rocket']).toJS()

export const getSingleRocket = (state) => getRocketState(state).get('currentRocket')