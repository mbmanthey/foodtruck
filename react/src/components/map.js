import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import Marker from './marker'
import mapboxgl from 'mapbox-gl';
import Overlay from './overlay';

const keys = require('./../config.json');
mapboxgl.accessToken = keys.mapbox;

class MapContainer extends Component {
  constructor(props) {
    super(props);
    this.state = {
      longitude: null,
      latitude: null,
      zoom: 15,
      error: null,
      mapLoaded: false,
      trucks: [],
    };
  }

  setMarker(marker) {
    ReactDOM.render(
      React.createElement(
        Marker
      ),
      marker
    )
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
    this.map = new mapboxgl.Map({
      container: this.mapContainer,
      style: 'mapbox://styles/mapbox/streets-v9',
      center: [this.state.longitude, this.state.latitude],
      zoom: this.state.zoom
    })
    this.setState({mapLoaded: true})
    this.map.on('move', () => {
      const { lng, lat } = this.map.getCenter();
      this.setState({
        longitude: lng.toFixed(4),
        latitude: lat.toFixed(4),
        zoom: this.map.getZoom().toFixed(2),
      })
    })

    this.map.on('click', (event) => {
      // this.addMarker(event.lngLat)
      this.addTrucks()
      console.log("clicked");
    })
  }

  addMarker(lngLat) {
    console.log(lngLat)
    var markerContainer = document.createElement('div')
    var marker = new mapboxgl.Marker(markerContainer).setLngLat([0,0]).addTo(this.map) 
    marker.setLngLat(lngLat);
    this.setMarker(markerContainer);
  }

  addTrucks() {
    console.log(this.state.trucks)
    for (let truck of this.state.trucks) {
      console.log(truck)
      let lngLat = new mapboxgl.LngLat(truck.Location.longitude, truck.Location.latitude);
      this.addMarker(lngLat)
    }

  }
  getLocalTrucks() {
    fetch('http://localhost:50051/api/truck', {
      method: "GET",
    }).then(result => result.json()
    ).then(data => {
      // console.log(data);
      this.setState({trucks: data})
      //this.addTrucks()
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

    return (
      <div>
          <div style={{display: 'inline-block', position: 'absolute', zIndex: '10'}}>
            <Overlay mapLoaded={this.state.mapLoaded}/>
          </div>
          <div style={style} ref={el => this.mapContainer = el}/>
      </div>
   )
  }
}
export default MapContainer;



