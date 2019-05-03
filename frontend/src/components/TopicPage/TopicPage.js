import React, { Component } from 'react';
import { Link, NavLink } from 'react-router-dom';
import cookie from 'react-cookies';
//import { Redirect } from 'react-router';
import './TopicPage.css';
import { userActions } from '../../_actions';
import { connect } from 'react-redux';
import { Container, Col, Card, Button, Row } from 'react-bootstrap';
import axios from 'axios';
import { msgstore_apis, david_test_apis, user_tracking_apis } from '../../config';
import renderHTML from 'react-render-html';
import Moment from 'react-moment';
import { BadgeGroup } from '../QuestionPage/QuestionPage'

var html_truncate = require('html-truncate');

class TopicPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            questions: [],
            user_name: 'Yu Zhao',
            topic: this.props.match.params.topic,
            token: cookie.load('JWT')
        }
    }

    componentDidMount() {
        axios.get(david_test_apis + '/questions', {
            headers: {
                'Authorization': `JWT ${this.state.token}`
            },
            params: {
                topic: this.state.topic,
                topAnswer: true,
                depth: 1,

            }
        }).then(response => {
            console.log(response.data);
            this.setState({
                questions: response.data
            })

        }).catch(error => {
            console.log(error);
        })
    }

    render() {

        let main_panel = null;

        if (this.state.questions.length !== 0) {
            main_panel = this.state.questions.map(q => {
                let answer = null;
                if ('answers' in q) {
                    console.log(html_truncate(q.answers[0].answerText, 3));
                    answer = (
                        <div>
                            <ul className="list-unstyled">
                                <li>{q.answers[0].displayName} </li>
                                <li><small className="text-muted">Answered <Moment fromNow>{q.answers[0].createdOn}</Moment></small></li>
                            </ul>
                            <p className="comment_truncate">
                                {renderHTML(html_truncate(q.answers[0].answerText, 150))}
                            </p>
                        </div>
                    )
                }
                else {
                    answer = <p>Be the first to answer this question! </p>
                }
                return (
                    <Card>
                        <Card.Body >
                            <BadgeGroup data={q.topics} />
                            <Card.Title style={{ fontWeight: 'bold' }}>{q.questionText}</Card.Title>
                            <Card.Text >
                                {answer}
                            </Card.Text>
                            <Card.Link as={NavLink} to={'/questions/' + q._id}>more</Card.Link>
                        </Card.Body>
                    </Card>

                )
            });
        }
        else {
            main_panel = (
                <div>
                    What is your question?
                </div>
            )
        }


        return (

            <div>
                <Card>
                    <Card.Body style={{ "font-size": 20, "color": '#666', 'fontWeight': 'bold' }}>
                        {this.state.topic}
                        <br />
                        <Button variant="link" onClick={this.handleFollow} disabled>
                            <span className="fa fa-plus-square"></span> Followed</Button>
                    </Card.Body>
                </Card>
                <br />
                {main_panel}
            </div>
        )
    }
}

const mapStateToProps = ({ authentication }) => ({ authentication });
export default connect(mapStateToProps)(TopicPage);