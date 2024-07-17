import "./post.scss";
import { Link } from "react-router-dom";
import FavoriteBorderOutlinedIcon from "@mui/icons-material/FavoriteBorderOutlined";
import ModeCommentOutlinedIcon from "@mui/icons-material/ModeCommentOutlined";
import ShareOutlinedIcon from "@mui/icons-material/ShareOutlined";
import FavoriteOutlinedIcon from "@mui/icons-material/FavoriteOutlined";
import MoreHorizTwoToneIcon from "@mui/icons-material/MoreHorizTwoTone";
import Comments from "../comments/Comments";
import moment from "moment";
import { useQuery } from "react-query";
import { makeRequest } from "../../axios";
import { colors } from "@mui/material";
import { useContext } from "react";
import { AuthContext } from "../../context/authContext";

const Post = ({ post, isCommentOpen, onToggleComments }) => {
  const { currentUser } = useContext(AuthContext);

  const { isLoading, error, data } = useQuery(["likes", post.id], () =>
    makeRequest.get("/likes?postId=" + post.id).then((res) => {
      return res.data;
    })
  );


 

  return (
    <div className="post">
      <div className="container">
        <div className="user">
          <div className="userInfo">
            <img src={post.profilePic} alt="" />
            <div className="details">
              <Link
                to={`/profile/${post.userId}`}
                style={{ textDecoration: "none", color: "inherit" }}
              >
                <span className="name">
                  {post.firstname} {post.lastname}
                </span>
              </Link>
              <span className="date">{moment(post.createdAt).fromNow()}</span>
            </div>
          </div>
          <MoreHorizTwoToneIcon />
        </div>
        <div className="content">
          <p>{post.desc}</p>
          <img src={post.img} alt="" />
        </div>

        <div className="info">
          <div className="item">
            {data ? (
              data[0].includes(currentUser.id) ? (
                <FavoriteOutlinedIcon style={{ color: "red" }} />
              ) : (
                <FavoriteBorderOutlinedIcon />
              )
            ) : (
              ""
            )}	

            {data ? data[0].length : 0} Likes
          </div>
          <div className="item" onClick={onToggleComments}>
            <ModeCommentOutlinedIcon />
            {data ? data[1].length : 0} Comments
          </div>
          <div className="item">
            <ShareOutlinedIcon />
            Share
          </div>
        </div>
        {isCommentOpen && <Comments postId={post.id} />}
      </div>
    </div>
  );
};

export default Post;
