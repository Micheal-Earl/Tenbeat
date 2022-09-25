import { h } from 'preact';
import axios from 'axios';
import { useEffect, useState } from "preact/hooks";
import style from './style.css';

import BookHeader from './book-header'

const baseURL = "http://localhost:9090"

const BookList = () => {
	const [bookList, setBookList] = useState([])

	useEffect(() => {
		axios.get(baseURL + "/books").then((response) => {
			setBookList(response.data);
		});
	}, []);

	if (!bookList) return null;

	return (
		<div class={style.books}>
			<BookHeader />
			<ul>
				{
					bookList.map((book) => {
						return (
							<li>
								<p>Title - {book.title}</p>
								<p>Author - {book.author}</p>
								<p>Desc - {book.desc}</p>
							</li>
						)
					})
				}
			</ul>
		</div>
	);
}

export default BookList;
