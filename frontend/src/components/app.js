import { h } from 'preact';
import { Router } from 'preact-router';

import Header from './header/header';

// Code-splitting is automated for `routes` directory
import Home from '../routes/home/home';

import BookHome from '../routes/books/book-home';
import BookList from '../routes/books/book-list';
import BookGet from '../routes/books/book-get';
import BookAdd from '../routes/books/book-add';
import BookUpdate from '../routes/books/book-update';
import BookDelete from '../routes/books/book-delete';

import UserHome from '../routes/user/user-home'
import UserLogin from '../routes/user/user-login';
import UserLogout from '../routes/user/user-logout';

import PrivateHome from '../routes/private/private-home'
import PrivateMe from '../routes/private/private-me';
import PrivateStatus from '../routes/private/private-status';

const App = () => (
	<div id="app">
		<Header />
		<Router>
			<Home path="/" />
			<BookHome path="/books/" />
			<BookList path="/books/list" />
			<BookGet path="/books/get/" />
			<BookAdd path="/books/add/" />
			<BookUpdate path="/books/update/" />
			<BookDelete path="/books/delete/" />
			<UserHome path="/user/" />
			<UserLogin path="/user/login" />
			<UserLogout path="/user/logout" />
			<PrivateHome path="/private/" />
			<PrivateMe path="/private/me" />
			<PrivateStatus path="/private/status" />
		</Router>
	</div>
)

export default App;

// 			<Books path="/books/delete/:id" />