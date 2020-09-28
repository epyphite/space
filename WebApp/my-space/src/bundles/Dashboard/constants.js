import mirrorCreator from 'mirror-creator';

export const actionTypes = mirrorCreator([
    'SEARCH_FILTERS_CLEAR_ALL',
    'ADD_SEARCH',
    'SET_ROCKET',
    'ADD_ROCKET'
], {prefix: 'ROCKET_'});

export const SORT_ENUM =  ['UPDATED_AT_DESC']