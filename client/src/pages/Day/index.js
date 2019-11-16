import React, { useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import api from '../../services/api';
import { withRouter } from 'react-router-dom';
import Container from '@material-ui/core/Container';
import ButtonGroup from '@material-ui/core/ButtonGroup';
import { Loading, Galaxy } from '../../components';

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
    paddingTop: '56.25%',
  },
  cardContent: {
    flexGrow: 1,
  },
  cardActions: {
    justifyContent: 'flex-end',
  },
  pagination: {
    marginTop: '20px',
  },
}));

function getPage(direction, actualPage) {
  const nextPage = parseInt(actualPage, 10) + direction;

  return nextPage >= 0 ? nextPage : 1;
}

function Day({ match, history }) {
  const [weather, setWeather] = useState({});
  const [loading, setLoading] = useState(false);
  const day = match.params.day;
  const classes = useStyles();

  useEffect(() => {
    async function loadData() {
      setLoading(true);
      const apiResponse = await api.get(`/weather/day/${day}`);
      setWeather(apiResponse.data);
      setLoading(false);
    }

    loadData();
  }, [day]);

  const scale = coord => {
    return coord / 10;
  };

  const getData = positions => {
    return (
      positions &&
      positions.map(p => {
        return {
          x: scale(p.x),
          y: scale(p.y),
          size: 1,
        };
      })
    );
  };

  const prevPage = () => {
    history.push(`/day/${getPage(-1, day)}`);
  };

  const nextPage = () => {
    history.push(`/day/${getPage(1, day)}`);
  };

  return (
    <>
      {loading ? (
        <Loading />
      ) : (
        <Container className={classes.cardGrid} maxWidth="xl">
          <Grid container spacing={4}>
            <Grid item key={weather.day} xs={12} sm={6} md={4}>
              <Card className={classes.card}>
                <CardMedia
                  className={classes.cardMedia}
                  image={IMAGES[weather.weather]}
                  title={weather.weather}
                />
                <CardContent className={classes.cardContent}>
                  <Typography gutterBottom variant="h5" component="h2">
                    Day: {weather.day}
                  </Typography>
                  <Typography>
                    <strong>Weather: </strong> {weather.weather}
                  </Typography>
                </CardContent>
              </Card>
            </Grid>

            <Grid item key={weather.day} xs={12} sm={6} md={8} align="center">
              <Galaxy data={getData(weather.positions)}></Galaxy>
              <Grid item xs={12} md={3} className={classes.pagination}>
                <ButtonGroup fullWidth>
                  <Button onClick={prevPage} disabled={day < 1}>
                    Previous
                  </Button>
                  <Button onClick={nextPage}>Next</Button>
                </ButtonGroup>
              </Grid>
            </Grid>
          </Grid>
        </Container>
      )}
    </>
  );
}

export default withRouter(Day);
