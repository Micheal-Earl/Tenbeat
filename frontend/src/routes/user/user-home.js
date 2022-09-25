import { h } from 'preact';
import style from './style.css';
import UserHeader from './user-header';

const UserHome = () => (
  <div class={style.user}>
    <UserHeader />
    <h1>Home</h1>
    <p>This is the User Home component.</p>
  </div>
);

export default UserHome;
