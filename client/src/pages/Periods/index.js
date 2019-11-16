import React, { useState, useEffect } from 'react';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import api from '../../services/api';
import { Link } from 'react-router-dom';

import { Loading } from '../../components';

import { IMAGES } from '../../constants';

const useStyles = makeStyles(theme => ({
  cardGrid: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  card: {
    height: '100%',
    display: 'flex',
    flexDirection: 'column',
  },
  cardMedia: {
    paddingTop: '56.25%', // 16:9
  },
  cardContent: {
    flexGrow: 1,
  },
  link: {
    textDecoration: 'none',
    cursor: 'pointer',
  },
}));

function Periods({ match, history }) {
  const classes = useStyles();

  const [days, setDays] = useState([]);
  const [loading, setLoading] = useState(false);
  const page = match.params.page || 1;

  useEffect(() => {
    async function loadData() {
      setLoading(true);
      const apiResponse = await api.get(`/periods`);
      setDays(apiResponse.data);
      setLoading(false);
    }

    loadData();
  }, [page]);

  return (
    <Container className={classes.cardGrid} maxWidth="xl">
      <Typography variant="h4" align="left" gutterBottom>
        Periods
      </Typography>
      {loading ? (
        <Loading />
      ) : (
        <Grid container spacing={2} justify="left">
          {days.map(day => (
            <Grid item key={day.start} xs={12} sm={6} md={2}>
              <Card className={classes.card}>
                <CardMedia
                  className={classes.cardMedia}
                  image={IMAGES[day.weather]}
                  title={day.weather}
                />
                <CardContent className={classes.cardContent}>
                  <Typography gutterBottom variant="h5" component="h2">
                    Days:{' '}
                    <Link to={`/day/${day.start}`} className={classes.link}>
                      {day.start}
                    </Link>{' '}
                    -{' '}
                    <Link to={`/day/${day.end}`} className={classes.link}>
                      {day.end}
                    </Link>
                  </Typography>
                  <Typography>
                    <strong>Weather: </strong> {day.weather}
                  </Typography>
                  {day.peak && (
                    <Typography>
                      <strong>Peak Day: </strong>
                      <Link to={`/day/${day.peak}`} className={classes.link}>
                        {day.peak}
                      </Link>
                    </Typography>
                  )}
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
      )}
    </Container>
  );
}

export default Periods;
