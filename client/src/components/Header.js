import React from 'react';
import AppBar from '@material-ui/core/AppBar';
import CloudIcon from '@material-ui/icons/Cloud';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import Link from '@material-ui/core/Link';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  icon: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
}));

export default function Header() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <CloudIcon className={classes.icon} />

          <Typography
            variant="h6"
            color="inherit"
            noWrap
            className={classes.title}
          >
            Galaxy Weather
          </Typography>

          <Link href="/" color="inherit">
            <Button color="inherit">Home</Button>
          </Link>
          <Link href="/periods" color="inherit">
            <Button color="inherit">Periods</Button>
          </Link>
        </Toolbar>
      </AppBar>
    </div>
  );
}
