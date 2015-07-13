import React, { Component } from 'react';

import ResultsContainer from './ResultsContainer';

export default class SearchApp extends Component {
  constructor(props) {
    super(props);
    this.state = {ids: []};
    this.to = null;
  }
  _onInputChange(event) {
    if (this.to) { clearTimeout(this.to); }
    if (!event || !event.target) { return; }
    this.to = setTimeout(() => {
      Parse.Cloud.run("search", {q: event.target.value}).then(
        (ids) => {
          this.setState({ids: ids, error: null});
        },
        (error) => {
          this.setState({error: error, ids: []});
        },
      );
    },
    900);
  }
  render() {
    const { ids, error } = this.state;
    return (
      <div>
        <input onChange={::this._onInputChange} />
        {error}
        <ResultsContainer className={this.props.className} ids={ids} />
      </div>
    );
  }
}
