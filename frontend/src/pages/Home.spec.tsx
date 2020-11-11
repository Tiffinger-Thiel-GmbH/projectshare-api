import { render } from 'enzyme';
import React from 'react';
import Home from './Home';

describe('Home', () => {
  it('renders', () => {
    const home = render(<Home />);
    expect(home.text()).toEqual('Hello World');
  });
});
