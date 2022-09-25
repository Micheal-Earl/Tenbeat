import { h } from 'preact';
import BookHeader from './book-header';
import style from './style.css';

const BookHome = () => (
  <div class={style.books}>
    <BookHeader />
    <h1>Home</h1>
    <p>This is the Books Home component.</p>
  </div>
);

export default BookHome;
