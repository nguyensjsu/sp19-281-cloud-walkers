import { Modal, Button, Form } from 'react-bootstrap';
import React, { Component } from 'react';
import AsyncSelect from 'react-select/lib/Async';
import axios from 'axios';
import _ from "lodash";
import {msgstore_apis, david_test_apis} from '../../config';

/*const options = [
//  { value: 1, label: 'Movies' },
//  { value: 2, label: 'Food' },
{label: 'Movies'}

];*/
/*
const filterColors = (inputValue) => {
  return options.filter(i =>
    i.label.toLowerCase().includes(inputValue.toLowerCase())
  );
};

const getOptions = inputValue => {
  axios.get('http://34.217.213.85:3000/msgstore/v1/topics')
  .then(response=>{
    console.log(response);
    return response.json();
  })
  .then(json => {
      return json.data.filter(i=>
        i.toLowerCase().includes(inputValue.toLowerCase()))
  }).catch(err => console.log(err))
*/
/*
return new Promise(resolve => {
  setTimeout(() => {
    resolve(filterColors(inputValue));
  }, 1000);
});
}*/

class TopicModal extends Component {
  constructor(props) {
    super(props);
    this.state = {
      selectedTopics: [],
      options: [],
      token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTY2ODI3MzksImlkIjoiNWNjNTIzNjA3MmM5YmYzNDM2ODJiNGIwIn0.2yvfGmvutYPygv_oPbj7QUdiDxVvxbh6o5eHYZ2CBUU'
    }
    this.handlePost = this.handlePost.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.handleSelectChange = this.handleSelectChange.bind(this);
  }

  handleChange = (e) => {
    this.setState({ [e.target.name]: e.target.value });
    console.log(e.target.name, e.target.value);
  }


  handlePost = (e) => {

  }

  componentDidMount() {
    /*
    axios.get('http://35.164.157.104:8000/msgstore/v1/topics', {
      headers:
      {
        'Authorization': this.state.token
      }
    })
      .then(response => {
        console.log(response.data);
        this.setState({
          options: response.data
        })
      })*/
  }

  handleSelectChange = (value, { action }) => {
    console.log(value, action);
    this.setState({
      selectedTopics: value
    })
  };


  getOptions = inputValue => {
    return axios.get(david_test_apis + '/topics', {
      headers: {
        'Authorization': `JWT ${this.state.token}`
      },
       params: {
         excludeFollowed: false
       }
     })
      .then(response => {
        console.log(response.data);
        //     this.setState({
        //       options: response.data
        //     })
        return response.data
      }).then(options => {
        const filtered = _.filter(options, o =>
          _.startsWith(_.toLower(o.label), _.toLower(inputValue))
        );
        return filtered.slice(0, 10);
      })
  }

  render() {

    //  const { selectedOption } = this.state;

    return (
      <Modal
        {...this.props}
        size="md"
        aria-labelledby="contained-modal-title-vcenter"
        centered
      >
        <Modal.Header closeButton>
          <Modal.Title id="contained-modal-title-vcenter">
            What are your interests?
            </Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group controlId="exampleForm.ControlTextarea1">
              <Form.Label className="text-muted">Select topics you want to follow</Form.Label>
              <AsyncSelect
                isMulti
                cacheOptions
                defaultOptions
                loadOptions={inputValue => this.getOptions(inputValue)}
                onChange={this.handleSelectChange}
                getOptionValue={option => option.label}
              />
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button style={{ 'color': '#949494', 'text-decoration': 'none', 'fontWeight': 400 }} variant="link" onClick={this.props.onHide}>Not now</Button>
          <Button onClick={this.handlePost}>Done</Button>
        </Modal.Footer>
      </Modal>
    );
  }
}
// apply above mapping to Login class
export default TopicModal