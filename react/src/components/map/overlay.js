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
        <div onClick={this.props.onClick} className='overlay'>
          <div className='pin'>
            <FontAwesomeIcon icon={faMapPin} size="3x"/>
            <h1 style={{padding:'5px'}}>Add</h1>
          </div>
        </div>
      );
    } else {
      return (null)
    }
  }
} 

export default Overlay;