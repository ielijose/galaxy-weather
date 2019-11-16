import React, { useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import ButtonGroup from '@material-ui/core/ButtonGroup';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import api from '../../services/api';
import { Link, withRouter } from 'react-router-dom';
import ViewIcon from '@material-ui/icons/Visibility';

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
    paddingTop: '56.25%',
  },
  cardContent: {
    flexGrow: 1,
  },
  cardActions: {
    justifyContent: 'flex-end',
  },
  link: {
    textDecoration: 'none',
  },
  pagination: {
    marginTop: '20px',
  },
}));

function getPage(direction, actualPage) {
  const nextPage = parseInt(actualPage, 10) + direction;

  return nextPage >= 0 ? nextPage : 1;
}

function Home({ match, history }) {
  const classes = useStyles();

  const [days, setDays] = useState([]);
  const [loading, setLoading] = useState(false);
  const page = match.params.page || 1;

  useEffect(() => {
    async function loadData() {
      setLoading(true);
      const apiResponse = await api.get(`/weather/year/${page}`);
      setDays(apiResponse.data);
      setLoading(false);
    }

    loadData();
  }, [page]);

  const prevPage = () => {
    history.push(`/${getPage(-1, page)}`);
  };

  const nextPage = () => {
    history.push(`/${getPage(1, page)}`);
  };

  return (
    <Container className={classes.cardGrid} maxWidth="xl">
      <Typography variant="h4" align="left" gutterBottom>
        Showing days of the year: {page}
      </Typography>
      {loading ? (
        <Loading />
      ) : (
        <Grid container spacing={2} justify="center">
          {days.map(day => (
            <Grid item key={day.day} xs={12} sm={6} md={2}>
              <Card className={classes.card}>
                <CardMedia
                  className={classes.cardMedia}
                  image={IMAGES[day.weather]}
                  title={day.weather}
                />
                <CardContent className={classes.cardContent}>
                  <Typography gutterBottom variant="h5" component="h2">
                    Day: {day.day}
                  </Typography>
                  <Typography>
                    <strong>Weather: </strong> {day.weather}
                  </Typography>
                </CardContent>
                <CardActions className={classes.cardActions}>
                  <Link to={`/day/${day.day}`} className={classes.link}>
                    <Button
                      size="small"
                      variant="contained"
                      color="primary"
                      startIcon={<ViewIcon />}
                    >
                      View Day
                    </Button>
                  </Link>
                </CardActions>
              </Card>
            </Grid>
          ))}
          <Grid item xs={12} md={3} className={classes.pagination}>
            <ButtonGroup fullWidth>
              <Button onClick={prevPage} disabled={page < 2}>
                Previous
              </Button>
              <Button onClick={nextPage}>Next</Button>
            </ButtonGroup>
          </Grid>
        </Grid>
      )}
    </Container>
  );
}

export default withRouter(Home);
