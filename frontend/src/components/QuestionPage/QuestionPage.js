import React, { Component } from 'react';
import { Link, NavLink } from 'react-router-dom';
import cookie from 'react-cookies';
//import { Redirect } from 'react-router';
import { userActions } from '../../_actions';
import { connect } from 'react-redux';
import './QuestionPage.css';
import { ListGroup, Container, Badge, ButtonToolbar, Button, Collapse, Card } from 'react-bootstrap';
import ReactPaginate from 'react-paginate';
import Moment from 'react-moment';
import PropTypes from 'prop-types';
import AnswerEditor from '../AnswerEditor/AnswerEditor';
import { throws } from 'assert';

const fake_reponse = {
    "spaceId": "spaceId",
    "createdBy": "createdBy",
    "topics": [{
        "topics": {
            "label": "label"
        }
    }, {
        "topics": {
            "label": "label"
        }
    }],
    "answers": [{
        "answers": {
            "questionId": "questionId",
            "comments": [{
                "replies": {
                    "answerId": "answerId",
                    "replies": [{}, {}],
                    "createdBy": "createdBy",
                    "parentCommentId": "parentCommentId",
                    "_id": "_id",
                    "commentText": "commentText",
                    "createdOn": "2000-01-23T04:56:07.000+00:00"
                }
            }, {
                "replies": {
                    "answerId": "answerId",
                    "replies": [{}, {}],
                    "createdBy": "createdBy",
                    "parentCommentId": "parentCommentId",
                    "_id": "_id",
                    "commentText": "commentText",
                    "createdOn": "2000-01-23T04:56:07.000+00:00"
                }
            }],
            "answerText": "answerText",
            "createdBy": "createdBy",
            "_id": "_id",
            "createdOn": "2000-01-23T04:56:07.000+00:00"
        }
    }, {
        "answers": {
            "questionId": "questionId",
            "comments": [{
                "replies": {
                    "answerId": "answerId",
                    "replies": [{}, {}],
                    "createdBy": "createdBy",
                    "parentCommentId": "parentCommentId",
                    "_id": "_id",
                    "commentText": "commentText",
                    "createdOn": "2000-01-23T04:56:07.000+00:00"
                }
            }, {
                "replies": {
                    "answerId": "answerId",
                    "replies": [{}, {}],
                    "createdBy": "createdBy",
                    "parentCommentId": "parentCommentId",
                    "_id": "_id",
                    "commentText": "commentText",
                    "createdOn": "2000-01-23T04:56:07.000+00:00"
                }
            }],
            "answerText": "answerText",
            "createdBy": "createdBy",
            "_id": "_id",
            "createdOn": "2000-01-23T04:56:07.000+00:00"
        }
    }],
    "_id": "_id",
    "createdOn": "2000-01-23T04:56:07.000+00:00",
    "questionText": "questionText"
}


export class AnswerList extends Component {
    static propTypes = {
        data: PropTypes.array.isRequired,
    };
    render() {
        let details = null;
        if (this.props.data.length != 0) {
            details = this.props.data.map((post, idx) => {
                return (
                    <ListGroup.Item key={idx}>
                        <ul className="list-unstyled">

                            <li>{post.answers.createdBy} </li>
                            <li><small className="text-muted">Answered<Moment fromNow>{post.answers.createdOn}</Moment></small></li>
                        </ul>
                        <p>
                            {post.answers.answerText}
                        </p>
                    </ListGroup.Item>

                )
            });
        }
        else {
            details = (<div>No Answer Yet</div>)
        }
        return (
            <ListGroup variant="flush">
                {details}
            </ListGroup>
        )
    }
}

export class BadgeGroup extends Component {
    static propTypes = {
        data: PropTypes.array.isRequired,
    };

    render() {
        let details = null;
        if (this.props.data.length != 0) {
            details = this.props.data.map((post, idx) => {
                return (
                    <Badge pill variant="light" className='topic_pill' key={idx}>
                        {post.topics.label}
                    </Badge>

                )
            });
        }
        else {
            details = (<div>No Answer Yet</div>)
        }
        return (
            <div >
                {details}
            </div>
        )
    }
}

export class AnswerInput extends Component {
    constructor(props) {
        super(props);
        this.state = {
            answer_string: null
        }
    }
    static propTypes = {
        onChange: PropTypes.func
    };

    handleSubmit = (e) =>{
        e.preventDefault();
        const data = {
            body: this.state.answer_string
        }

        

    }


    handleInputChange = (value) => {
        console.log(value);
        this.setState({ answer_string: value })
    }

    render() {
        return (
            <AnswerEditor onChange={this.handleInputChange} />
        )
    }
}

class QuestionPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            user_name: "Yu Zhao",
            answer_input: false,
            question: null
        }
    }
    componentWillMount(){
        this.setState({
            question: fake_reponse
        })
    }
    render() {
        return (
            <div>
                <Container>
                    <ListGroup variant="flush">
                        <ListGroup.Item>
                            <BadgeGroup data={this.state.question.topics} />
                            <h4><b>{this.state.question.questionText}</b></h4>
                            <ButtonToolbar>
                                <Button variant="link" onClick={() => this.setState({ answer_input: !this.state.answer_input })}>Answer</Button>
                                <Button variant="link">Follow</Button>
                            </ButtonToolbar>
                            <Collapse in={this.state.answer_input}>
                                <Card>
                                    <Card.Header>{this.state.user_name}</Card.Header>
                                    <Card.Body>
                                        <AnswerInput q_id={this.state.question._id} />
                                    </Card.Body>
                                    <Card.Footer className="text-muted">
                                        <Button size="sm"> Submit</Button>
                                    </Card.Footer>
                                </Card>

                            </Collapse>
                        </ListGroup.Item>
                    </ListGroup>

                    <AnswerList data={this.state.question.answers} />
                </Container>
            </div>
        )
    }

}

const mapStateToProps = ({ authentication }) => ({ authentication });
export default connect(mapStateToProps)(QuestionPage);