import { h } from 'preact';
import axios from 'axios';
import { useEffect, useState } from "preact/hooks";
import style from './style.css';

import BookHeader from './book-header'

const baseURL = "http://localhost:9090"

const BookAdd = () => {
	const [title, setTitle] = useState(null)
	const [author, setAuthor] = useState(null)
	const [desc, setDesc] = useState(null)

	// useEffect(() => {
	// 	axios.get(baseURL + "/books").then((response) => {
	// 		setBookList(response.data);
	// 	});
	// }, []);

	function addBook() {
		axios
			.post(baseURL + "/books", {
				title: title,
				author: author,
				desc: desc
			}, { crossDomain: true })
	}

	// if (!book) return null;

	return (
		<div class={style.books}>
			<BookHeader />
			<form>
				<label>Title:
					<input
						type="text"
						name="title"
						value={title}
						onInput={e => setTitle(e.target.value)}
					/>
				</label>
				<label>Author:
					<input
						type="text"
						name="Author"
						value={author}
						onInput={e => setAuthor(e.target.value)}
					/>
				</label>
				<label>Desc:
					<input
						type="text"
						name="Desc"
						value={desc}
						onInput={e => setDesc(e.target.value)}
					/>
				</label>
				<input type="button" value="Submit" onClick={addBook} />
			</form>
		</div>
	);
}

export default BookAdd;
