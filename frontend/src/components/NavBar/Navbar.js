import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import cookie from 'react-cookies';
//import { Redirect } from 'react-router';
import './Navbar.css';
import { userActions } from '../../_actions';
import { connect } from 'react-redux';

//create the Navbar Component
class Navbar extends Component {
    constructor(props) {
        super(props);
        this.handleLogout = this.handleLogout.bind(this);
    }
    //handle logout to destroy the cookie
    handleLogout = () => {
        this.props.dispatch(userActions.logout());
        cookie.remove('JWT', { path: '/' });
    }
    render() {
        //if Cookie is set render Logout Button

        const { authentication } = this.props;

        let navLogin = null;
        if (authentication.loggedIn === true) {
 //           let userId = cookie.load('userId');
 //           console.log("Able to read cookie");
            navLogin = (
                <ul className="nav navbar-nav navbar-right">
                    <li><Link to={`/profile`}><span className="glyphicon glyphicon-user"></span> Profile</Link></li>
                    <li><Link to="/" onClick={this.handleLogout}><span className="glyphicon glyphicon-off"></span> Logout</Link></li>
                </ul>
            );
        } else {
            //Else display login button
//            console.log("Not Able to read cookie");
            navLogin = (
                <ul className="nav navbar-nav navbar-right">
                    <li><Link to="/login"><span className="glyphicon glyphicon-log-in"></span> Login</Link></li>
                    <li><Link to="/signup"><span className="glyphicon glyphicon-plus"></span> SignUp</Link></li>
                </ul>
            )
        }

        let redirectVar = null;
 //       if (this.props.authentication.loggedIn !== true) {
 //           redirectVar = <Redirect to="/login" />
 //       }

        let welcome = null;
        if ('role' in authentication && authentication.role !== null) {
            let role = authentication.role;
            //redirectVar = <Redirect to="/courses" />;
            welcome = <p className="navbar-text">Signed in as {role}</p>
        }
        return (
            <div>
                {redirectVar}
                <nav className="navbar navbar-default navbar-static-top" id="nav_bar">
                    <div className="container-fluid">
                        <div className="navbar-header">
                            <a className="navbar-brand" href="#">Canvas</a>
                            {welcome}
                        </div>
                        <ul className="nav navbar-nav">

                            <li><Link to="/courses"><span className="glyphicon glyphicon-home"></span> Dashboard</Link></li>
                            <li><Link to="/messages"><span className="glyphicon glyphicon-envelope"></span> Message</Link></li>
                            {/*{nav_role}
                                                   
                        <li><Link to="/create">Add a Book</Link></li>
                        <li><Link to="/delete">Delete a Book</Link></li>
*/}
                        </ul>
                        {navLogin}

                    </div>
                </nav>
            </div>
        )
    }
}

//export default Navbar;
const mapStateToProps = ({ authentication }) => ({ authentication });
// apply above mapping to Login class
export default connect(mapStateToProps)(Navbar);