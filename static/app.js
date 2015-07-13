/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};

/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {

/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId])
/******/ 			return installedModules[moduleId].exports;

/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			exports: {},
/******/ 			id: moduleId,
/******/ 			loaded: false
/******/ 		};

/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);

/******/ 		// Flag the module as loaded
/******/ 		module.loaded = true;

/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}


/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;

/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;

/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";

/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(0);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	'use strict';

	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

	var _react = __webpack_require__(1);

	var _react2 = _interopRequireDefault(_react);

	var _SearchApp = __webpack_require__(2);

	var _SearchApp2 = _interopRequireDefault(_SearchApp);

	_react2['default'].render(_react2['default'].createElement(_SearchApp2['default'], { className: ClassName }), document.getElementById('root'));

/***/ },
/* 1 */
/***/ function(module, exports) {

	module.exports = React;

/***/ },
/* 2 */
/***/ function(module, exports, __webpack_require__) {

	'use strict';

	Object.defineProperty(exports, '__esModule', {
	  value: true
	});

	var _createClass = (function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ('value' in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; })();

	var _get = function get(_x, _x2, _x3) { var _again = true; _function: while (_again) { var object = _x, property = _x2, receiver = _x3; desc = parent = getter = undefined; _again = false; if (object === null) object = Function.prototype; var desc = Object.getOwnPropertyDescriptor(object, property); if (desc === undefined) { var parent = Object.getPrototypeOf(object); if (parent === null) { return undefined; } else { _x = parent; _x2 = property; _x3 = receiver; _again = true; continue _function; } } else if ('value' in desc) { return desc.value; } else { var getter = desc.get; if (getter === undefined) { return undefined; } return getter.call(receiver); } } };

	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

	function _inherits(subClass, superClass) { if (typeof superClass !== 'function' && superClass !== null) { throw new TypeError('Super expression must either be null or a function, not ' + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) subClass.__proto__ = superClass; }

	var _react = __webpack_require__(1);

	var _react2 = _interopRequireDefault(_react);

	var _ResultsContainer = __webpack_require__(3);

	var _ResultsContainer2 = _interopRequireDefault(_ResultsContainer);

	var SearchApp = (function (_Component) {
	  _inherits(SearchApp, _Component);

	  function SearchApp(props) {
	    _classCallCheck(this, SearchApp);

	    _get(Object.getPrototypeOf(SearchApp.prototype), 'constructor', this).call(this, props);
	    this.state = { ids: [] };
	    this.to = null;
	  }

	  _createClass(SearchApp, [{
	    key: '_onInputChange',
	    value: function _onInputChange(event) {
	      var _this = this;

	      if (this.to) {
	        clearTimeout(this.to);
	      }
	      if (!event || !event.target) {
	        return;
	      }
	      this.to = setTimeout(function () {
	        Parse.Cloud.run('search', { q: event.target.value }).then(function (ids) {
	          _this.setState({ ids: ids, error: null });
	        }, function (error) {
	          _this.setState({ error: error, ids: [] });
	        });
	      }, 900);
	    }
	  }, {
	    key: 'render',
	    value: function render() {
	      var _state = this.state;
	      var ids = _state.ids;
	      var error = _state.error;

	      return _react2['default'].createElement(
	        'div',
	        null,
	        _react2['default'].createElement('input', { onChange: this._onInputChange.bind(this) }),
	        error,
	        _react2['default'].createElement(_ResultsContainer2['default'], { className: this.props.className, ids: ids })
	      );
	    }
	  }]);

	  return SearchApp;
	})(_react.Component);

	exports['default'] = SearchApp;
	module.exports = exports['default'];

/***/ },
/* 3 */
/***/ function(module, exports, __webpack_require__) {

	'use strict';

	Object.defineProperty(exports, '__esModule', {
	  value: true
	});

	var _createClass = (function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ('value' in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; })();

	var _get = function get(_x, _x2, _x3) { var _again = true; _function: while (_again) { var object = _x, property = _x2, receiver = _x3; desc = parent = getter = undefined; _again = false; if (object === null) object = Function.prototype; var desc = Object.getOwnPropertyDescriptor(object, property); if (desc === undefined) { var parent = Object.getPrototypeOf(object); if (parent === null) { return undefined; } else { _x = parent; _x2 = property; _x3 = receiver; _again = true; continue _function; } } else if ('value' in desc) { return desc.value; } else { var getter = desc.get; if (getter === undefined) { return undefined; } return getter.call(receiver); } } };

	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

	function _inherits(subClass, superClass) { if (typeof superClass !== 'function' && superClass !== null) { throw new TypeError('Super expression must either be null or a function, not ' + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) subClass.__proto__ = superClass; }

	var _react = __webpack_require__(1);

	var _react2 = _interopRequireDefault(_react);

	var _Results = __webpack_require__(4);

	var _Results2 = _interopRequireDefault(_Results);

	var ResultsContainer = (function (_Component) {
	  _inherits(ResultsContainer, _Component);

	  function ResultsContainer(props) {
	    _classCallCheck(this, ResultsContainer);

	    _get(Object.getPrototypeOf(ResultsContainer.prototype), 'constructor', this).call(this, props);
	    this.state = { results: [] };
	  }

	  _createClass(ResultsContainer, [{
	    key: 'componentDidMount',
	    value: function componentDidMount() {
	      this._populate(this.props.ids);
	    }
	  }, {
	    key: 'componentWillReceiveProps',
	    value: function componentWillReceiveProps(nextProps) {
	      if (nextProps.ids == this.props.ids) {
	        return;
	      }
	      this._populate(nextProps.ids);
	    }
	  }, {
	    key: '_populate',
	    value: function _populate(ids) {
	      var _this = this;

	      var Object = Parse.Object.extend(this.props.className);
	      var query = new Parse.Query(Object);
	      query.containedIn('objectId', ids);
	      query.find({
	        success: function success(results) {
	          _this.setState({ results: results, error: null });
	        },
	        error: function error(_error) {
	          _this.setState({ error: _error, results: null });
	        }
	      });
	    }
	  }, {
	    key: 'render',
	    value: function render() {
	      var results = this.state.results;

	      return _react2['default'].createElement(_Results2['default'], { results: results });
	    }
	  }]);

	  return ResultsContainer;
	})(_react.Component);

	exports['default'] = ResultsContainer;
	module.exports = exports['default'];

/***/ },
/* 4 */
/***/ function(module, exports, __webpack_require__) {

	'use strict';

	Object.defineProperty(exports, '__esModule', {
	  value: true
	});

	var _createClass = (function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ('value' in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; })();

	var _get = function get(_x, _x2, _x3) { var _again = true; _function: while (_again) { var object = _x, property = _x2, receiver = _x3; desc = parent = getter = undefined; _again = false; if (object === null) object = Function.prototype; var desc = Object.getOwnPropertyDescriptor(object, property); if (desc === undefined) { var parent = Object.getPrototypeOf(object); if (parent === null) { return undefined; } else { _x = parent; _x2 = property; _x3 = receiver; _again = true; continue _function; } } else if ('value' in desc) { return desc.value; } else { var getter = desc.get; if (getter === undefined) { return undefined; } return getter.call(receiver); } } };

	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

	function _inherits(subClass, superClass) { if (typeof superClass !== 'function' && superClass !== null) { throw new TypeError('Super expression must either be null or a function, not ' + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) subClass.__proto__ = superClass; }

	var _react = __webpack_require__(1);

	var _react2 = _interopRequireDefault(_react);

	var Results = (function (_Component) {
	  _inherits(Results, _Component);

	  function Results(props) {
	    _classCallCheck(this, Results);

	    _get(Object.getPrototypeOf(Results.prototype), 'constructor', this).call(this, props);
	  }

	  _createClass(Results, [{
	    key: 'render',
	    value: function render() {
	      var results = this.props.results;

	      return _react2['default'].createElement(
	        'div',
	        null,
	        results.length,
	        ' results',
	        _react2['default'].createElement(
	          'ul',
	          null,
	          results.map(function (result) {
	            return _react2['default'].createElement(
	              'li',
	              { key: result.id },
	              '- ',
	              JSON.stringify(result.toJSON())
	            );
	          })
	        )
	      );
	    }
	  }]);

	  return Results;
	})(_react.Component);

	exports['default'] = Results;
	module.exports = exports['default'];

/***/ }
/******/ ]);