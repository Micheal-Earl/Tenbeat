import { h } from 'preact';
import axios from 'axios';
import { useEffect, useState } from "preact/hooks";
import style from './style.css';

import BookHeader from './book-header'

const baseURL = "http://localhost:9090"

const BookUpdate = () => {
	const [title, setTitle] = useState(null)
	const [author, setAuthor] = useState(null)
	const [desc, setDesc] = useState(null)
	const [id, setID] = useState(null)
	const [book, setBook] = useState(null)

	function addBook() {
		setBook({
			title: title,
			author: author,
			desc: desc
		})

		axios
			.put(baseURL + "/books/" + id, {
				title: title,
				author: author,
				desc: desc
			}, { crossDomain: true })
	}

	if (!book) {
		return (
			<div class={style.books}>
				<BookHeader />
				<form>
					<label>ID:
						<input
							type="text"
							name="title"
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
	} else {
		return (
			<div class={style.books}>
				<BookHeader />
				<form>
					<label>ID:
						<input
							type="text"
							name="title"
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
				<div id={style.book}>
					<p>Updated Book:</p>
					<p>Title - {book.title}</p>
					<p>Author - {book.author}</p>
					<p>Desc - {book.desc}</p>
				</div>
			</div>
		);
	}
}

export default BookUpdate;