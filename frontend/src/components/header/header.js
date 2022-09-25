import { h } from 'preact';
import { Link } from 'preact-router/match';
import style from './style.css';

const Header = () => (
	<header class={style.header}>
		<h1>Book App</h1>
		<nav>
			<Link activeClassName={style.active} href="/">Home</Link>
			<Link activeClassName={style.active} href="/books">Books</Link>
			<Link activeClassName={style.active} href="/user">User</Link>
			<Link activeClassName={style.active} href="/private">Private</Link>
		</nav>
	</header>
);

export default Header;
