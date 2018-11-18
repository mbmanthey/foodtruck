import * as React from 'react';
import "./overlay.css";
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMapPin } from '@fortawesome/free-solid-svg-icons'

library.add(faMapPin);

export class Overlay extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    if (this.props.mapLoaded) {
      return (
        <div className='overlay'>
          <div className='identify'>
            <FontAwesomeIcon icon={faMapPin} size="3x"/>
          </div>
        </div>
      );
    } else {
      return (null)
    }
  }
} 

export default Overlay;