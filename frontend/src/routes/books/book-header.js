import { h } from 'preact';
import { Link } from 'preact-router/match';
import style from './style.css';

const BookHeader = () => (
  <div class={style.header}>
    <h1>Books</h1>
    <nav>
      <Link activeClassName={style.active} href="/books/">Books Home</Link>
      <Link activeClassName={style.active} href="/books/list">List Books</Link>
      <Link activeClassName={style.active} href="/books/get">Get Book</Link>
      <Link activeClassName={style.active} href="/books/add">Add Book</Link>
      <Link activeClassName={style.active} href="/books/update">Update Book</Link>
      <Link activeClassName={style.active} href="/books/delete">Delete Book</Link>
    </nav>
  </div>
);

export default BookHeader;
