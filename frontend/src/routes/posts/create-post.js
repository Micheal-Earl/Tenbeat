import axios from 'axios';
import { useState } from "preact/hooks";
import style from './style.css';

import PostHeader from './post-header'

const baseURL = "http://localhost:9090"

const CreatePost = () => {
	const [title, setTitle] = useState(null)
	const [content, setContent] = useState(null)

	function addBook() {
		axios
			.post(baseURL + "/posts", {
				title: title,
				content: content,
			}, { crossDomain: true, withCredentials: true })
	}

	return (
		<div class={style.posts}>
			<PostHeader />
			<form>
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
				<input type="button" value="Submit" onClick={addBook} />
			</form>
		</div>
	);
}

export default CreatePost;
