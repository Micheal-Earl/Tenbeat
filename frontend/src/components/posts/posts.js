import { h } from 'preact';
import { Link } from 'preact-router/match';
import axios from 'axios';
import { useEffect, useState } from "preact/hooks";
import style from './style.css';

const data = [
  {
    id: 1,
    title: "This is a post",
    content: "Here is some content. On the other hand, we denounce with righteous indignation and dislike men who are so beguiled and demoralized by the charms of pleasure of the moment, so blinded by desire, that they cannot foresee the pain and trouble that are bound to ensue; and equal blame belongs to those who fail in their duty through weakness of will, which is the same as saying through shrinking from toil and pain. These cases are perfectly simple and easy to distinguish. In a free hour, when our power of choice is untrammelled and when nothing prevents our being able to do what we like best, every pleasure is to be welcomed and every pain avoided. But in certain circumstances and owing to the claims of duty or the obligations of business it will frequently occur that pleasures have to be repudiated and annoyances accepted. The wise man therefore always holds in these matters to this principle of selection: he rejects pleasures to secure other greater pleasures, or else he endures pains to avoid worse pains.",
    owner: {
      username: "Mike",
    },
    created_at: "Oct 1",
    updated_at: "Oct 1",
    deleted_at: "",
  },
  {
    id: 2,
    title: "A rant",
    content: "Morbi vel malesuada ligula. Integer varius laoreet diam. Vivamus viverra auctor commodo. In faucibus, augue non vulputate blandit, elit nulla euismod tellus, eget suscipit tortor ante nec est. Proin porttitor malesuada sollicitudin. Mauris tellus libero, efficitur sed ultricies accumsan, congue non nisl. Etiam aliquet orci a justo feugiat sodales. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque sem ligula, congue eget nunc sed, auctor vehicula nunc. Fusce vitae eleifend ex, non viverra arcu. Curabitur in lobortis nunc. Integer sollicitudin fermentum odio nec fringilla. Proin facilisis, massa at facilisis fringilla, velit tellus semper velit, sit amet pulvinar ante diam at libero. Vivamus nec orci imperdiet, iaculis ligula non, tempus nunc. Sed vel sapien interdum, fringilla ipsum nec, imperdiet risus. Sed non augue id justo ultricies luctus et nec urna. Proin tempus sapien nec nisi dignissim finibus. Praesent et consequat purus. Aliquam quam purus, ullamcorper sed ullamcorper at, dapibus nec felis. Vivamus scelerisque, sapien vitae malesuada lacinia, ante lacus semper enim, a lacinia ipsum risus fermentum tortor. Praesent vel ex vel nunc gravida laoreet a vitae lacus. Etiam congue mi non tortor venenatis aliquet. Nam cursus quis ligula pharetra feugiat. Aliquam nec tortor vel eros pharetra bibendum. Morbi justo ante, mollis nec dapibus id, molestie in tellus. Fusce vitae tincidunt lacus. Nulla porttitor ex nisi, at accumsan dui varius et. Nam nunc orci, vulputate non volutpat eget, consectetur ac ipsum. Nullam ac magna ante. Aliquam erat volutpat. Vestibulum sit amet risus urna. Vivamus lacus mauris, tincidunt a dignissim sit amet, egestas in mi. Aenean tellus velit, tincidunt sed justo sit amet, tempus tristique sem. Nam at sollicitudin ex. ",
    owner: {
      username: "Kane",
    },
    created_at: "Sept 28",
    updated_at: "Oct 1",
    deleted_at: "",
  },
]

const baseURL = "http://localhost:9090"

function Posts() {
  const [postList, setPostList] = useState([])

  useEffect(() => {
    axios.get(baseURL + "/posts").then((response) => {
      setPostList(response.data);
    });
  }, []);

  const listPosts = postList.map(post =>
    <li class={style.postitem} key={post.id}>
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
    </li>
  );

  return (
    <ul class={style.posts}>
      {listPosts}
    </ul>
  );
}

export default Posts;
