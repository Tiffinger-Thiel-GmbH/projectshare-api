import * as React from 'react';
import AppBar from './AppBar';
import CssBaseline from '@material-ui/core/CssBaseline';
import Content from './Content';

const Home = () => {
  return (
    <>
      <CssBaseline />

      <AppBar />
      <Content />
    </>
  );
};

export default Home;
