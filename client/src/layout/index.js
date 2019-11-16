import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
import { Header, Footer } from '../components';
import { Route, Switch } from 'react-router-dom';
import Home from '../pages/Home';
import Day from '../pages/Day';
import Periods from '../pages/Periods';

export default function Layout() {
  return (
    <React.Fragment>
      <CssBaseline />
      <Header />
      <main>
        <Switch>
          <Route path="/periods" exact component={Periods} />
          <Route path="/:page?" exact component={Home} />
          <Route path="/day/:day" exact component={Day} />
        </Switch>
      </main>
      <Footer />
    </React.Fragment>
  );
}
