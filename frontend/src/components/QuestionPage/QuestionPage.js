import React, { Component } from 'react';
import { Link, NavLink } from 'react-router-dom';
import cookie from 'react-cookies';
//import { Redirect } from 'react-router';
import { userActions } from '../../_actions';
import { connect } from 'react-redux';
import './QuestionPage.css';
import { ListGroup, Container, Badge, ButtonToolbar, Button, Collapse, Card, Form, Col, Row } from 'react-bootstrap';
import ReactPaginate from 'react-paginate';
import Moment from 'react-moment';
import PropTypes from 'prop-types';
import AnswerEditor from '../AnswerEditor/AnswerEditor';
import renderHTML from 'react-render-html'

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
            "answerText": "<ul><li>dag</li><li>sdgd</li><li>sgdg</li></ul>",
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
            "answerText": "<p>Test2 Test2 Test2</p>",
            "createdBy": "createdBy",
            "_id": "_id",
            "createdOn": "2000-01-23T04:56:07.000+00:00"
        }
    }],
    "_id": "_id",
    "createdOn": "2000-01-23T04:56:07.000+00:00",
    "questionText": "questionText"
}

export class CommentPanel extends Component {
    static propTypes = {
        data: PropTypes.array.isRequired,
    };

    constructor(props) {
        super(props);
        this.state = {
            comment_text: false,
            show_comments: false,
        }
    }
    onChange = (e) => {
        //        console.log(e.target.value)
        this.setState({ 'comment_text': e.target.value })
    }
    render() {
        const { comment_text } = this.state

        let comment_list = null;
        if (this.state.show_comments === true)
            comment_list = this.props.data.map((comment, idx) => {
                return (
                    <ListGroup.Item key={idx} style={{ border: 'none' }}>
                        <ul className="list-unstyled">
                            <li><small>{comment.replies.createdBy}</small></li>
                            <li><small>  {comment.replies.commentText} </small> </li>
                        </ul>
                    </ListGroup.Item>
                )
            });
        return (
            <div className="threaded_comments">
                <Form inline>
                    <Form.Group >
                        <Button size="sm" className="btn-circle">
                            <span className="fa fa-user fa-lg" style={{ color: 'FFFFFF' }}></span>
                        </Button>
                        <Form.Control style={{ 'margin-left': 12, 'width': 600 }}
                            as="textarea" rows="1" size="sm" type="text" placeholder="Add a comment..."
                            onChange={this.onChange} />
                        <Button style={{ 'margin-left': 12 }} size="sm" disabled={!comment_text}> Add Comment</Button>


                        <Button style={{ 'margin-left': 12 }} size="sm" variant="link" onClick={() => this.setState({ "show_comments": true })}> All </Button>
                    </Form.Group>
                </Form>
                <Collapse in={this.state.show_comments}>
                    <ListGroup>
                        {comment_list}
                    </ListGroup>
                </Collapse>
            </div>
        )
    }
}

export class AnswerList extends Component {
    static propTypes = {
        data: PropTypes.array.isRequired,
    };

    constructor(props) {
        super(props);
        this.state = {
        }
    }
    handleUpvote = (value) => {
        console.log('upvote this answer', value)
    }

    handleDownvote = (value) => {
        console.log('downvote this answer', value)
    }

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
                            {renderHTML(post.answers.answerText)}
                        </p>
                        <ButtonToolbar style={{ 'margin-left': -10 }}>
                            <Button className="q_page_button" variant="link" onClick={() => this.handleUpvote(post.answers._id)}>
                                <span className="fa fa-arrow-up"></span> Upvote</Button>
                            <Button className="q_page_button pull-right" variant="link" onClick={() => this.handleDownvote(post.answers._id)}>
                                <span className="fa fa-arrow-down"></span> Downvote</Button>
                        </ButtonToolbar>
                        <CommentPanel data={post.answers.comments}></CommentPanel>
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
            answer_string: null,
        }
    }
    static propTypes = {
        onChange: PropTypes.func
    };

    handleSubmit = (e) => {
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
        this.handleFollow = this.handleFollow.bind(this);
    }
    componentWillMount() {
        this.setState({
            question: fake_reponse
        })
    }

    handleFollow() {
        return (
            console.log("follow this question")
        )
    }

    render() {
        return (
            <div>
                <Container>
                    <ListGroup variant="flush">
                        <ListGroup.Item>
                            <BadgeGroup data={this.state.question.topics} />
                            <h4><b>{this.state.question.questionText}</b></h4>
                            <ButtonToolbar style={{ 'margin-left': -10 }}>
                                <Button className="q_page_button" variant="link" onClick={() => this.setState({ answer_input: !this.state.answer_input })}>
                                    <span className="fa fa-edit"></span> Answer</Button>
                                <Button className="q_page_button" variant="link" onClick={this.handleFollow}>
                                    <span className="fa fa-plus-square"></span> Follow</Button>
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
            </div >
        )
    }

}

const mapStateToProps = ({ authentication }) => ({ authentication });
export default connect(mapStateToProps)(QuestionPage);