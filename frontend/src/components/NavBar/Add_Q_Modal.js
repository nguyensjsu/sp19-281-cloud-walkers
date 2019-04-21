import { Modal, Button, Form } from 'react-bootstrap';
import React, { Component } from 'react';

class AddQModal extends Component {
  render() {
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
              <Form.Label>{this.props.user_name} asked</Form.Label>
              <Form.Control as="textarea" rows="3" />
              <Form.Text className="text-muted">
                Start your question with "What", "How", "Why", etc.
    </Form.Text>
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="light" onClick={this.props.onHide}>Cancel</Button>
          <Button onClick={this.props.handleAdd}>Add Question</Button>
        </Modal.Footer>
      </Modal>
    );
  }
}
// apply above mapping to Login class
export default AddQModal