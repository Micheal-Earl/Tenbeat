import { h } from 'preact';
import axios from 'axios';
import { useEffect, useState } from "preact/hooks";
import style from './style.css';

import BookHeader from './book-header'

const baseURL = "http://localhost:9090"

const BookGet = () => {
	const [book, setBook] = useState(null)
	const [id, setID] = useState(null)

	function getBook() {
		axios.get(baseURL + "/books/" + id).then((response) => {
			setBook(response.data);
		});
	}

	if (!book) {
		return (
			<div class={style.books}>
				<BookHeader />
				<form>
					<label>Book ID:
						<input
							type="text"
							name="id"
							value={id}
							onInput={e => setID(e.target.value)}
						/>
					</label>
					<input type="button" value="Submit" onClick={getBook} />
				</form>
			</div>
		);
	} else {
		return (
			<div class={style.books}>
				<BookHeader />
				<form>
					<label>Book ID:
						<input
							type="text"
							name="id"
							value={id}
							onInput={e => setID(e.target.value)}
						/>
					</label>
					<input type="button" value="Submit" onClick={getBook} />
				</form>
				<div id={style.book}>
					<p>Title - {book.title}</p>
					<p>Author - {book.author}</p>
					<p>Desc - {book.desc}</p>
				</div>
			</div>
		)
	}
}

export default BookGet;
