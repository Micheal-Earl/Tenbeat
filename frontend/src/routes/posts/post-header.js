import { Link } from 'preact-router/match';
import style from './style.css';

const PostHeader = () => (
  <div class={style.header}>
    <h1>posts</h1>
    <nav>
      <Link activeClassName={style.active} href="/posts/">Posts Home</Link>
      <Link activeClassName={style.active} href="/posts/get">Get Post</Link>
      <Link activeClassName={style.active} href="/posts/add">Add Post</Link>
      <Link activeClassName={style.active} href="/posts/update">Update Post</Link>
      <Link activeClassName={style.active} href="/posts/delete">Delete Post</Link>
    </nav>
  </div>
);

export default PostHeader;
