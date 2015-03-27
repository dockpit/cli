var React = require('react');

var DepActions  = require('../actions/DepActions')
var DepStore = require('../stores/DepStore')
var IsolationStore = require('../stores/IsolationStore')
var DepItem = require("./DepItem.jsx")
var DepForm = require('./DepForm.jsx')

module.exports = React.createClass({

  getInitialState: function() {
    return {
      data: DepStore.state(),
      isoData: IsolationStore.state(),
      showForm: false,
    }
  },

  componentDidMount: function() {
  	DepStore.on(DepStore.CHANGED, this.onStoreChange)
    IsolationStore.on(IsolationStore.CHANGED, this.onStoreChange)
    DepActions.refresh()
  },

  componentDidUnmount: function() {
  	DepStore.removeListener(DepStore.CHANGED, this.onStoreChange)
  },

  onStoreChange: function() {
  	this.setState({data: DepStore.state(), isoData: IsolationStore.state()})
  },

  openDepForm: function() {
    this.setState({ showForm: true });
  },

  closeDepForm: function(ev) {
    ev.preventDefault();
    this.setState({ showForm: false });
  },

  render: function() {
    var me = this
    var list = this.state.data.get('deps').map(function(dep){
      return <DepItem isolations={me.state.isoData.get('isolations')} key={dep.get('name')} dep={dep}/> 
    })
    
    if (this.state.data.get('deps').size < 1) {
      list = <div className="ui attached bottom segment"><div className="ui icon message">
          <i className="cube icon"></i>
          <div className="content">
            <div className="header">
              You don't have any dependencies configured
            </div>
            <p>A dependency represents another service you depend on: A database, HTTP API, message queues, etc</p>
          </div>
        </div></div>
    }

    return <div style={{position: 'relative'}}>
      <h2 className="ui top attached header">
        <div className="content">
          Dependencies          
        </div>
      </h2>

      <button style={{position: 'absolute', top: '10px', right: '10px'}} className="ui labeled icon button primary" onClick={this.openDepForm}><i className="plus icon"></i>New Dependency</button>

      { this.state.showForm ? <DepForm closeFormFn={this.closeDepForm} dep={this.state.data.get('selection')}/> : null }

      {list}
      {this.state.data.get('deps').size > 0 ? <div className="ui bottom attached segment"></div> : null }
    
    </div>;
  }
});