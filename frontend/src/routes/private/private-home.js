import { h } from 'preact';
import style from './style.css';
import PrivateHeader from './private-header';

const PrivateHome = () => (
  <div class={style.user}>
    <PrivateHeader />
    <h1>Home</h1>
    <p>This is the User Home component.</p>
  </div>
);

export default PrivateHome;
