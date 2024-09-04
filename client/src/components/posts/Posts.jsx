import Post from "../post/Post";
import "./posts.scss";
import { useState } from "react";
import { InfinitySpin } from "react-loader-spinner";
import { useQuery } from "react-query";
import { makeRequest } from "../../axios";

const Posts = () => {

  const [openComments, setOpenComments] = useState({});
 



  const toggleComments = (postId) => {
    setOpenComments((prevState) => ({
      ...prevState,
      [postId]: !prevState[postId],
    }));
  };
  //using react query

  const { isLoading, error, data } = useQuery(["posts"], () =>
    makeRequest.get("/posts").then((res) => {
      return res.data;
    })
  );

  return (
    <div className="posts">
      {isLoading ? (
        <div className="loading">
          <InfinitySpin
            visible={true}
            width="200"
            color="purple"
            ariaLabel="infinity-spin-loading"
          />
        </div>
      ) : error ? (
        "oops Something went wrong!"
      ) : data ? (
        data.map((post) => (
          <Post
            post={post}
            postId={post.id}
            isCommentOpen={openComments[post.id] || false} // to see
            onToggleComments={() => toggleComments(post.id)}
          />
        ))
      ) : (
        ""
      )}
    </div>
  );
};

export default Posts;
