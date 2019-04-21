import React, { Component } from 'react';
import { Link, NavLink } from 'react-router-dom';
import cookie from 'react-cookies';
//import { Redirect } from 'react-router';
//import './Navbar.css';
import { userActions } from '../../_actions';
import { connect } from 'react-redux';
import Sidebar from '../Sidebar/Sidebar';
import { Container,Col } from 'react-bootstrap';

class Home extends Component {


    render() {
        const sidebar_links = [
            { name: "Movies", url: "topics/1" }, 
            { name: "Food", url: "topics/2" }
        ];

        const questions = [
            {"questionText": "questionText1",
             "top_answer": 
            {
               answerText: 'AnswerText1',
               createdOn: "2000-01-23T04:56:07.000+00:00",
               createdBy: "createdBy1"
            } },
            {"questionText": "questionText2",
            "top_answer": 
           {
              answerText: 'AnswerText2',
              createdOn: "2000-01-23T04:56:07.000+00:00",
              createdBy: "createdBy2"
           } }
        ]

        return (

            <Container fluid>
                <Col xs={3} style={{ "margin-top": 50}}>
                    <Sidebar links={sidebar_links} />
                </Col>
                <Col md={{offset: 3 }} style={{ "margin-top": 50}}> 
                     Home Page </Col>
            </Container>
        )
    }
}

const mapStateToProps = ({ authentication }) => ({ authentication });
export default connect(mapStateToProps)(Home);