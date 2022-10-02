import axios from 'axios';
import { useState } from "preact/hooks";
import style from './style.css';

import PostHeader from './post-header'

const baseURL = "http://localhost:9090"

const GetPost = () => {
	const [post, setPost] = useState(null)
	const [id, setID] = useState(null)

	function getPost() {
		axios.get(baseURL + "/posts/" + id).then((response) => {
			setPost(response.data);
		});
	}

	if (!post) {
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
					<input type="button" value="Submit" onClick={getPost} />
				</form>
			</div>
		);
	} else {
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
					<input type="button" value="Submit" onClick={getPost} />
				</form>
				<div id={style.post}>
					<div class={style.postthumbnail}>
						<img src="https://i.imgur.com/r1FFPz9.jpeg"></img>
					</div>
					<div class={style.postcontent}>
						<h2>{post.title}</h2>
						<span class={style.postmetadata}>
							<p>Author: {post.owner.username}</p>
							<p>Created on: {post.createdAt}</p>
							<p>Updated on: {post.updatedAt}</p>
						</span>
						<details>
							<summary>Open Post</summary>
							<p>{post.content}</p>
						</details>
					</div>
				</div>
			</div>
		)
	}
}

export default GetPost;
