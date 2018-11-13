import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import Marker from './marker'
import mapboxgl from 'mapbox-gl';
import { config } from '@fortawesome/fontawesome-svg-core';

const keys = require('./../config.json');
mapboxgl.accessToken = keys.mapbox;

class MapContainer extends Component {
  markerContainer;
  constructor(props) {
    super(props);
    this.state = {
      longitude: null,
      latitude: null,
      zoom: 15,
      error: null,
    };
  }

  setMarker() {
    if (true) {
      ReactDOM.render(
        React.createElement(
          Marker
        ),
        this.markerContainer
      )
    } else {
      this.markerContainer.innerHTML = '';
    }
  }

  componentDidMount() {
    navigator.geolocation.getCurrentPosition(
      (position) => {
        this.setState({
          longitude: position.coords.longitude.toFixed(4),
          latitude: position.coords.latitude.toFixed(4),
          error: '',
        }),
        this.initializeMap()
        this.getLocalTrucks()
      },
      (error) => this.setState({ error: error.message }),
      {enableHighAccuracy: true, timeout: 10000, maximumAge:60000},
    )
  }

  initializeMap() {
    this.markerContainer = document.createElement('div')
    this.map = new mapboxgl.Map({
      container: this.mapContainer,
      style: 'mapbox://styles/mapbox/streets-v9',
      center: [this.state.longitude, this.state.latitude],
      zoom: this.state.zoom
    })

    const marker = new mapboxgl.Marker(this.markerContainer).setLngLat([0,0]).addTo(this.map)

    this.map.on('move', () => {
      const { lng, lat } = this.map.getCenter();

      this.setState({
        longitude: lng.toFixed(4),
        latitude: lat.toFixed(4),
        zoom: this.map.getZoom().toFixed(2)
      })
    })

    this.map.on('click', (event) => {
      marker.setLngLat(event.lngLat);
      console.log("clicked");
      this.setMarker();
    })
  }

  getLocalTrucks() {
    fetch('http://localhost:50051/api/truck', {
      method: "GET",
    }).then(result => result.json()
    ).then(data => {
      console.log(data)
    })
  }

  componentWillUnmount() {
    if (this.map) {
      this.map.remove()
    }
  }

  render() {
    const style = {
      position: 'absolute',
      top: 0,
      bottom: 0,
      width: '100%'
    };

    return <div style={style} ref={el => this.mapContainer = el}/>
  }
}
export default MapContainer;



