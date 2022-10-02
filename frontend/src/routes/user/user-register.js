import axios from 'axios';
import { useState } from "preact/hooks";
import style from './style.css';

import UserHeader from './user-header'

const baseURL = "http://localhost:9090"

const UserRegister = () => {
  const [email, setEmail] = useState(null)
  const [username, setUsername] = useState(null)
  const [password, setPassword] = useState(null)
  const [res, setRes] = useState(null)

  function register() {
    axios
      .post(baseURL + "/user/register", {
        email: email,
        username: username,
        password: password,
      }, { crossDomain: true, withCredentials: true }).then((response) => {
        setRes(response.data);
      }).catch((error) => {
        setRes(error.response.data);
      })
  }

  return (
    <div class={style.user}>
      <UserHeader />
      <form>
        <label>Email:
          <input
            type="text"
            name="email"
            value={email}
            onInput={e => setEmail(e.target.value)}
          />
        </label>
        <label>Username:
          <input
            type="text"
            name="username"
            value={username}
            onInput={e => setUsername(e.target.value)}
          />
        </label>
        <label>Password:
          <input
            type="password"
            name="Password"
            value={password}
            onInput={e => setPassword(e.target.value)}
          />
        </label>
        <input type="button" value="Submit" onClick={register} />
      </form>
      <p>{messageOrError(res)}</p>
    </div>
  );
}

function messageOrError(res) {
  if (!res) return ""

  if (res.hasOwnProperty("error")) {
    return res.error
  }

  if (res.hasOwnProperty("message")) {
    return res.message
  }

  return "invalid response"
}

export default UserRegister;
