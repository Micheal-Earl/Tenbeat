import { h } from 'preact';
import { Link } from 'preact-router/match';
import style from './style.css';

const Header = () => (
	<header class={style.header}>
		<h1>Tenbeat</h1>
		<nav>
			<Link activeClassName={style.active} href="/">Home</Link>
			<Link activeClassName={style.active} href="/posts">Posts</Link>
			<Link activeClassName={style.active} href="/user">User</Link>
			<Link activeClassName={style.active} href="/private">Private</Link>
		</nav>
	</header>
);

export default Header;
