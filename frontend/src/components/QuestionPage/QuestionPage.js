import React, { Component } from 'react';
import { Link, NavLink } from 'react-router-dom';
import cookie from 'react-cookies';
//import { Redirect } from 'react-router';
import { userActions } from '../../_actions';
import { connect } from 'react-redux';
import { ListGroup, Container, ListGroupItem } from 'react-bootstrap';
import ReactPaginate from 'react-paginate';
import Moment from 'react-moment';
import PropTypes from 'prop-types';

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
            details = this.props.data.map(post => {
                return (
                    <ListGroup.Item>
                        <ul class="list-unstyled">
                            <li>{post.answers.createdBy}</li>
                            <li><small class="text-muted">Answered<Moment fromNow>{post.answers.createdOn}</Moment></small></li>
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



class QuestionPage extends Component {
    constructor(props) {
        super(props);
        this.state = {

        }
    }

    render() {
        return (
            <div>
                <Container>
                    <ListGroup variant="flush">
                        <ListGroupItem>
                            <h4><b>{fake_reponse.questionText}</b></h4>
                        </ListGroupItem>
                    </ListGroup>

                    <AnswerList data={fake_reponse.answers} />
                </Container>
            </div>
        )
    }

}

const mapStateToProps = ({ authentication }) => ({ authentication });
export default connect(mapStateToProps)(QuestionPage);