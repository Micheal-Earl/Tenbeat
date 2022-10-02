import { h } from 'preact';
import axios from 'axios';
import { useEffect, useState } from "preact/hooks";
import style from './style.css';

import PrivateHeader from './private-header';

const baseURL = "http://localhost:9090"

const PrivateStatus = () => {
	const [res, setRes] = useState(null)

	useEffect(() => {
		axios
			.get(baseURL + "/status", { crossDomain: true, withCredentials: true }).then((response) => {
				setRes(response.data);
			}).catch((error) => {
				setRes(error.response.data);
			})
	}, []);

	if (!res) return null

	return (
		<div class={style.user}>
			<PrivateHeader />
			<p>{messageOrError(res)}</p>
		</div>
	);
}

function messageOrError(res) {
	if (!res) return ""
	if (res.hasOwnProperty("status")) {
		return res.status
	} else if (res.hasOwnProperty("error")) {
		return res.error
	} else {
		return "invalid response"
	}
}

export default PrivateStatus;
