import * as React from 'react';
import ListItem from "./ListItem"

export class List extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      longitude: null,
      latitude: null,
      zoom: 15,
      error: null,
      mapLoaded: false,
      trucks: []
    }
  }

  componentDidMount() {
    fetch('http://localhost:50051/api/truck', {
      method: "GET",
    }).then(result => result.json()
    ).then(data => {
      console.log(data)
      this.setState({trucks: data})
    })

  }



  render() {
    return (
      <ul>
        {this.state.trucks.map(i => {
          return(<ListItem key={i.ID} truck={i}/>)
        })}
      </ul>
    )
  }
} 

export default List;