import * as React from 'react';
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTruck } from '@fortawesome/free-solid-svg-icons'

library.add(faTruck);

export class Marker extends React.Component {
  render() {
    return (
      <div>
        <FontAwesomeIcon icon={faTruck} size="4x"/>
      </div>
    );
  }
} 

export default Marker;