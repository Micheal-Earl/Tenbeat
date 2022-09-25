import { h } from 'preact';
import axios from 'axios';
import { useEffect, useState } from "preact/hooks";
import style from './style.css';

import UserHeader from './user-header'

const baseURL = "http://localhost:9090"

const UserLogout = () => {
	const [res, setRes] = useState(null)

	function login() {
		axios
			.get(baseURL + "/logout", { crossDomain: true, withCredentials: true }).then((response) => {
				setRes(response.data);
			}).catch((error) => {
				setRes(error.response.data);
			})
	}

	return (
		<div class={style.user}>
			<UserHeader />
			<form>
				<input type="button" value="Log Out" onClick={login} />
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

export default UserLogout;
