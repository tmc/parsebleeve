import React, { Component } from 'react';

import Results from './Results';

export default class ResultsContainer extends Component {
  constructor(props) {
    super(props);
    this.state = {results: []};
  }
  componentDidMount() {
    this._populate(this.props.ids);
  }
  componentWillReceiveProps(nextProps) {
    if (nextProps.ids == this.props.ids) { return; }
    this._populate(nextProps.ids);
  }
  _populate(ids) {
    var Object = Parse.Object.extend(this.props.className);
    var query = new Parse.Query(Object);
    query.containedIn("objectId", ids);
    query.find({
      success: (results) => {
        this.setState({results: results, error: null});
      },
      error: (error) => {
        this.setState({error: error, results: null});
      },
    })
  }
  render() {
    const {results} = this.state;
    return (
      <Results results={results} />
    );
  }
}
