import { Modal, Button, Form } from 'react-bootstrap';
import React, { Component } from 'react';
import AsyncSelect from 'react-select/lib/Async';
import axios from 'axios';

const options = [
//  { value: 1, label: 'Movies' },
//  { value: 2, label: 'Food' },
{label: 'Movies'}

];
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

class AddQModal extends Component {
  constructor(props) {
    super(props);
    this.state = {
      user_name: "Yu Zhao",
      selectedTopics: [],
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

  componentDidMount(){

  }

  handleSelectChange = (value, { action }) => {
    console.log(value, action);
    this.setState({
      selectedTopics: value
    })
  };


 getOptions = inputValue => {
  axios.get('http://34.217.213.85:3000/msgstore/v1/topics')
  .then(response=>{
    console.log(response);
    return response.json();
  })
  .then(json => {
      return json.data.filter(i=>
        i.toLowerCase().includes(inputValue.toLowerCase()))
  }).catch(err => console.log(err))
  /*
  return new Promise(resolve => {
    setTimeout(() => {
      resolve(filterColors(inputValue));
    }, 1000);
  })*/;
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
            Add Question
            </Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form>
            <Form.Group controlId="exampleForm.ControlTextarea1">
              <Form.Label>{this.state.user_name} asked</Form.Label>
              <Form.Text className="text-muted">
                Start your question with "What", "How", "Why", etc.
              </Form.Text>
              <Form.Control as="textarea" rows="3" name="questionText" onChange={this.handleChange} />
              <Form.Text className="text-muted">
                Select topics you want to post question to
              </Form.Text>
              <AsyncSelect
                isMulti
                cacheOptions
                defaultOptions
                loadOptions={this.getOptions}
                onChange={this.handleSelectChange}
                name="select"
              />
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="light" onClick={this.props.onHide}>Cancel</Button>
          <Button onClick={this.handleAdd}>Add Question</Button>
        </Modal.Footer>
      </Modal>
    );
  }
}
// apply above mapping to Login class
export default AddQModal