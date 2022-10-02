import { h } from 'preact';
import { Router } from 'preact-router';

import Header from './header/header';
import Posts from './posts/posts';

// Code-splitting is automated for `routes` directory
// import Home from '../routes/home/home';

import PostHome from '../routes/posts/post-home';
import GetPost from '../routes/posts/get-post';
import CreatePost from '../routes/posts/create-post';
import UpdatePost from '../routes/posts/update-post';
import DeletePost from '../routes/posts/delete-post';

import UserHome from '../routes/user/user-home'
import UserRegister from '../routes/user/user-register';
import UserLogin from '../routes/user/user-login';
import UserLogout from '../routes/user/user-logout';

import PrivateHome from '../routes/private/private-home'
import PrivateMe from '../routes/private/private-me';
import PrivateStatus from '../routes/private/private-status';

const App = () => (
	<div id="app">
		<Header />
		<Router>
			<Posts path="/" />
			<PostHome path="/posts/" />
			<GetPost path="/posts/get/" />
			<CreatePost path="/posts/add/" />
			<UpdatePost path="/posts/update/" />
			<DeletePost path="/posts/delete/" />
			<UserHome path="/user/" />
			<UserRegister path="/user/register" />
			<UserLogin path="/user/login" />
			<UserLogout path="/user/logout" />
			<PrivateHome path="/private/" />
			<PrivateMe path="/private/me" />
			<PrivateStatus path="/private/status" />
		</Router>
	</div>
)

export default App;

// 			<Posts path="/posts/delete/:id" />