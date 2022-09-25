import { h } from 'preact';
import { Link } from 'preact-router/match';
import style from './style.css';

const UserHeader = () => (
  <div class={style.header}>
    <h1>User</h1>
    <nav>
      <Link activeClassName={style.active} href="/user">User Home</Link>
      <Link activeClassName={style.active} href="/user/login">Login</Link>
      <Link activeClassName={style.active} href="/user/logout">Logout</Link>
    </nav>
  </div>
);

export default UserHeader;
