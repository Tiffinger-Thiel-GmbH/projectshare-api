import React from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import MatAppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1
    },
    menuButton: {
      marginRight: theme.spacing(2)
    }
  })
);

export default function AppBar() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <MatAppBar position="static">
        <Toolbar variant="dense">
          <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" color="inherit">
            OfficeShare - projektspezifische Dokumentenverwaltung
          </Typography>
        </Toolbar>
      </MatAppBar>
    </div>
  );
}
