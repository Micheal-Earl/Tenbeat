import axios from 'axios';
import { useState } from "preact/hooks";
import style from './style.css';

import PostHeader from './post-header'

const baseURL = "http://localhost:9090"

const UpdatePost = () => {
	const [title, setTitle] = useState(null)
	const [content, setContent] = useState(null)
	const [id, setID] = useState(null)
	const [post, setPost] = useState(null)

	function updatePost() {
		setPost({
			title: title,
			content: content,
		})

		axios
			.put(baseURL + "/posts/" + id, {
				title: title,
				content: content,
			}, { crossDomain: true, withCredentials: true })
	}

	if (!post) {
		return (
			<div class={style.posts}>
				<PostHeader />
				<form>
					<label>ID:
						<input
							type="text"
							name="id"
							value={id}
							onInput={e => setID(e.target.value)}
						/>
					</label>
					<label>Title:
						<input
							type="text"
							name="title"
							value={title}
							onInput={e => setTitle(e.target.value)}
						/>
					</label>
					<label>Content:
						<input
							type="text"
							name="Content"
							value={content}
							onInput={e => setContent(e.target.value)}
						/>
					</label>
					<input type="button" value="Submit" onClick={updatePost} />
				</form>
			</div>
		);
	} else {
		return (
			<div class={style.posts}>
				<PostHeader />
				<form>
					<label>ID:
						<input
							type="text"
							name="id"
							value={id}
							onInput={e => setID(e.target.value)}
						/>
					</label>
					<label>Title:
						<input
							type="text"
							name="title"
							value={title}
							onInput={e => setTitle(e.target.value)}
						/>
					</label>
					<label>Content:
						<input
							type="text"
							name="Content"
							value={content}
							onInput={e => setContent(e.target.value)}
						/>
					</label>
					<input type="button" value="Submit" onClick={updatePost} />
				</form>
				<div id={style.post}>
					<div class={style.postthumbnail}>
						<img src="https://i.imgur.com/r1FFPz9.jpeg"></img>
					</div>
					<div class={style.postcontent}>
						<h2>{post.title}</h2>
						<span class={style.postmetadata}>
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
		);
	}
}

export default UpdatePost;