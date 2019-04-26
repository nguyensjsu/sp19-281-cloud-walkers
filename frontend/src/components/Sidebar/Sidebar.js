import React, { Component } from 'react';
import './Sidebar.css';
import { NavLink } from 'react-router-dom';
import { Nav, Button} from 'react-bootstrap';
import TopicModal from './TopicModal';
class Sidebar extends Component {
    constructor(props){
        super(props);
        this.state = {
            sidebar_links:[]
        }
    }

    selectTopics = (e) => {
        e.preventDefault();
        this.setState({
            show_topics: true
        })
    }

    componentDidMount() {
        const sidebar_links = [
            { name: "Movies", url: "topics" },
            { name: "Food", url: "topics" }
        ];

        this.setState({
            sidebar_links: sidebar_links,
        });
    }  

    render() {
        let sidebar_body = this.state.sidebar_links.map(link => {
            return (
                <Nav.Link className="sidebar" as={NavLink} to={link.url}>{link.name}</Nav.Link>
            )
        });
        let modal_T_Close = () => this.setState({ show_topics: false });
        return (
            <div>
            <TopicModal
                    show={this.state.show_topics}
                    onHide={modal_T_Close}
                />
            <Nav style={{"font-size": 14, "line-height": 10}} className="flex-column" >
                {sidebar_body}
            </Nav>
            <Button style={{ 'text-decoration': 'none', "font-size": 14, "line-height": 10 }} variant="link" onClick={this.selectTopics}>Follow More</Button>
            </div>
        );
    }
}

export default Sidebar;