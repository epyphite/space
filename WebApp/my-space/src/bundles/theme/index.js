import { createMuiTheme } from '@material-ui/core';
import typography from './typography';



const MuiTheme = createMuiTheme({
  palette: {
    
  },
  shape: {
    // borderRadius: '0.5rem'
  },
  overrides: {
    MuiButton: {
      text: {
        paddingLeft: '14px',
        paddingRight: '14px'
      },
      containedSizeSmall: {
        paddingLeft: '14px',
        paddingRight: '14px'
      },
      root: {
        textTransform: 'none',
        fontWeight: 'normal',
        borderRadius: 4
      }
    },
    MuiOutlinedInput: {
      root: {
        borderRadius: 4,
        backgroundColor: '#F8F8F8'
      },
      input: {
        paddingTop: 12,
        paddingBottom: 12,
        color: 'rgba(0, 0, 0, 0.7);'
      }
    },
    MuiAutocomplete: {
      inputRoot: {
        '&.MuiOutlinedInput-root': {
          paddingTop: 2.5,
          paddingBottom: 2.5
        }
      }
    }
  },
  typography
});

export default MuiTheme;
