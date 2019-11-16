import React from 'react';

import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Link from '@material-ui/core/Link';

const useStyles = makeStyles(theme => ({
  footer: {
    backgroundColor: theme.palette.background.paper,
    padding: theme.spacing(4),
  },
}));

export default function Footer() {
  const classes = useStyles();
  return (
    <footer className={classes.footer}>
      <Typography variant="body2" color="textSecondary" align="center">
        {'Made with ❤ by '}
        <Link
          color="inherit"
          target="_blank"
          href="http://bit.ly/linkedin-ielijose"
        >
          Eli José Carrasquero
        </Link>
      </Typography>
      <Typography color="textSecondary" align="center">
        <Link
          color="inherit"
          target="_blank"
          href="http://bit.ly/galaxy-weather-gh"
        >
          <strong>Fork me on GitHub</strong>
        </Link>
      </Typography>
    </footer>
  );
}
