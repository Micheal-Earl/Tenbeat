import { h } from 'preact';
import PostHeader from './post-header';
import style from './style.css';

const PostHome = () => (
  <div class={style.posts}>
    <PostHeader />
    <h1>Home</h1>
    <p>This is the posts Home component.</p>
  </div>
);

export default PostHome;
