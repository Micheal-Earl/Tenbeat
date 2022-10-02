import axios from 'axios';
import { useState } from "preact/hooks";
import style from './style.css';

import PostHeader from './post-header'

const baseURL = "http://localhost:9090"

const DeletePost = () => {
	const [id, setID] = useState(null)

	function deletePost() {
		axios.delete(baseURL + "/posts/" + id, { crossDomain: true, withCredentials: true })
	}

	return (
		<div class={style.posts}>
			<PostHeader />
			<form>
				<label>Post ID:
					<input
						type="text"
						name="id"
						value={id}
						onInput={e => setID(e.target.value)}
					/>
				</label>
				<input type="button" value="Submit" onClick={deletePost} />
			</form>
		</div>
	);
}

export default DeletePost;
