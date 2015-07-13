import React, { Component } from 'react';

export default class Results extends Component {
  constructor(props) {
    super(props);
  }
  render() {
    const {results} = this.props;
    return (
      <div>
        {results.length} results
        <ul>
        {results.map(result =>
          <li key={result.id}>- {JSON.stringify(result.toJSON())}</li>
        )}
        </ul>
      </div>
    );
  }
}
