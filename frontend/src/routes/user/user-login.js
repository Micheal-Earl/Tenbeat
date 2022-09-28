import { h } from 'preact';
import axios from 'axios';
import { useEffect, useState } from "preact/hooks";
import style from './style.css';

import UserHeader from './user-header'

const baseURL = "http://localhost:9090"

const UserLogin = () => {
	const [username, setUsername] = useState(null)
	const [password, setPassword] = useState(null)
	const [res, setRes] = useState(null)

	function login() {
		axios
			.post(baseURL + "/login", {
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
				<input type="button" value="Submit" onClick={login} />
			</form>
			<p>{messageOrError(res)}</p>
		</div>
	);
}

function messageOrError(res) {
	if (!res) return ""
	if (res.hasOwnProperty("message")) {
		return res.message
	} else if (res.hasOwnProperty("error")) {
		return res.error
	} else {
		return "invalid response"
	}
}

export default UserLogin;
