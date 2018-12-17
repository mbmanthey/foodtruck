import * as React from 'react';

export class ListItem extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        {this.props.truck.Name}
      </div>
    );
  }
} 

export default ListItem;