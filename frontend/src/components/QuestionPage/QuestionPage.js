import React, { Component } from 'react';
import { Link, NavLink } from 'react-router-dom';
import cookie from 'react-cookies';
//import { Redirect } from 'react-router';
import { userActions } from '../../_actions';
import { connect } from 'react-redux';
import { Container, Col, Card, Button } from 'react-bootstrap';
import ReactPaginate from 'react-paginate';


const Answer = (

)

const AnswerList = 
class QuestionPage extends Component {


}

const mapStateToProps = ({authentication}) => ({authentication});
export default connect(mapStateToProps)(QuestionPage);