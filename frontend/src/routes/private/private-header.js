import { h } from 'preact';
import { Link } from 'preact-router/match';
import style from './style.css';

const PrivateHeader = () => (
  <div class={style.header}>
    <h1>Private</h1>
    <nav>
      <Link activeClassName={style.active} href="/private">Private Home</Link>
      <Link activeClassName={style.active} href="/private/me">Me</Link>
      <Link activeClassName={style.active} href="/private/status">Status</Link>
    </nav>
  </div>
);

export default PrivateHeader;
