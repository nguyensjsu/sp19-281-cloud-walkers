import React, { Component } from 'react';
import { Route, Switch } from 'react-router-dom';
import Login from './Login/Login';
import SignUp from './SignUp/SignUp';
import Home from './Home/Home';
import AddCourse from './Course/AddCouse';
import SearchCourse from './Course/SearchCourse';
import DropCourse from './Course/DropCourse';
import Profile from './Profile/Profile';
import UpdateProfile from './Profile/UpdateProfileForm'
//import Delete from './Delete/Delete';
//import Create from './Create/Create';
import Navbar from './LandingPage/Navbar';
import Course from './Course/Course';
import CourseAddCode from './Course/CourseAddCode';
import Announcement from './Course/Announcement';
import NewAnnouncement from './Course/NewAnnouncement';
import People from './Course/People';
import HW_List from './Course/HW_List';
import New_HW from './Course/New_HW';
import HW_Detail from './Course/HW_Detail';
import FileList from './Course/FileList';
import NewFile from './Course/NewFile';
import SubmissionList from './Course/SubmissionList';
import SubSubmissionList from './Course/SubSubmissionList';
import GradeSubmission from './Course/GradeSubmission';
import New_Quiz from './Course/AddQuiz';
import Quiz_Detail from './Course/Quiz_Detail';
import QuizList from './Course/QuizList';
import Message from './Message/Message'

//Create a Main Component
class Main extends Component {
    render() {
        return (
            <div>
                <Route path="/" component={Navbar} />

                <Switch>
                    {/*Render Different Component based on Route*/}
                    <Route path="/login" component={Login} />
                    <Route path='/signup' component={SignUp} />

                    <Route path='/courses/add' component={AddCourse} />
                    <Route path='/courses/search' component={SearchCourse} />

                    <Route path='/courses/drop' component={DropCourse} />
                    

                    <Route path='/courses/:courseId/quizzes/new' component={New_Quiz} /> 
                    <Route path='/courses/:courseId/quizzes/:quizId' component={Quiz_Detail} />                                        
                    <Route path='/courses/:courseId/quizzes' component={QuizList} />

                    <Route path='/courses/:courseId/files/new' component={NewFile} />
                    <Route path='/courses/:courseId/files' component={FileList} />

                    <Route path='/courses/:courseId/assignments/new' component={New_HW} />
                    
                    <Route path='/courses/:courseId/submissions/:assignmentId/:submissionId' component={GradeSubmission} />
                    <Route path='/courses/:courseId/submissions/:assignmentId/' component={SubSubmissionList} />
                    <Route path='/courses/:courseId/submissions/' component={SubmissionList} />


                    <Route path='/courses/:courseId/assignments/:assignmentId' component={HW_Detail} />
                    <Route path='/courses/:courseId/assignments' component={HW_List} />

                    <Route path='/courses/:courseId/people' component={People} />

                    <Route path='/courses/:courseId/announcements/new' component={NewAnnouncement} />
                    <Route path='/courses/:courseId/announcements' exact component={Announcement} />


                    <Route path='/courses/:courseId/addcode' component={CourseAddCode} />
                    <Route path='/courses/:courseId' component={Course} />

                    <Route path="/courses" exact component={Home} />

                    <Route path='/profile/edit' component={UpdateProfile} />
                    <Route path='/profile' exact component={Profile} />
                    <Route path='/messages' component={Message}/>

                </Switch>
            </div>
        )
    }
}
//Export The Main Component
export default Main;