import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';
import FeatureDetailComponent from '../components/FeatureDetail';
import { listEnvironments, getFeature, toggleFeature } from '../utils/api';

class FeatureDetail extends Component {
  constructor(props) {
    super(props);

    this.state = {
      loading: true,
      environments: [],
      feature: null,
    };

    this.handleToggle = this.handleToggle.bind(this);
  }

  componentDidMount() {
    Promise.all([
      listEnvironments(),
      getFeature(window.decodeURIComponent(this.props.match.params.name))
    ]).then(([environments, feature]) => this.setState({
        loading: false,
        environments,
        feature,
      }));
  }

  handleToggle(name, value) {
    toggleFeature(name, this.state.feature.name, value).then(() => {
      this.setState({
        feature: Object.assign({}, this.state.feature, {
          status: Object.assign({}, this.state.feature.status, {
            [name]: value
          })
        })
      });
    });
  }

  render() {
    if (this.state.loading) return null;
    return (
      <FeatureDetailComponent
        environments={this.state.environments}
        feature={this.state.feature}
        onToggle={this.handleToggle}
      />
    );
  }
}

FeatureDetail.propTypes = {
  match: PropTypes.shape({
    params: PropTypes.object,
  }).isRequired,
};

export default withRouter(FeatureDetail);
