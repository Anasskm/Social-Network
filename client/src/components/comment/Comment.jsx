import { useContext, useState } from "react";
import { AuthContext } from "../../context/authContext";
import "./comment.scss";
import FavoriteBorderOutlinedIcon from "@mui/icons-material/FavoriteBorderOutlined";
import ModeCommentOutlinedIcon from "@mui/icons-material/ModeCommentOutlined";
import ShareOutlinedIcon from "@mui/icons-material/ShareOutlined";
import FavoriteOutlinedIcon from "@mui/icons-material/FavoriteOutlined";
import MoreHorizTwoToneIcon from "@mui/icons-material/MoreHorizTwoTone";
import ComOfCom from "../comOfComments/ComOfCom";
import moment from "moment";
import { useQuery } from "react-query";
import { makeRequest } from "../../axios";

const Comment = ({ comment }) => {
  const [comOfComOpen, setComOfComOpen] = useState(false);

  const { currentUser } = useContext(AuthContext);

  const { isLoading, error, data } = useQuery(["comlikes", comment.id], () =>
    makeRequest.get("/comlikes?commentId=" + comment.id).then((res) => {
      return res.data;
    })
  );

  return (
    <div className="comment">
      <div className="userItems">
        <div className="left">
          <img src={comment.profilePic} alt="" />
          <span>
            {comment.firstname} {comment.lastname}
          </span>
        </div>

        <MoreHorizTwoToneIcon className="more" />
      </div>

      <div className="info">
        <p>{comment.desc}</p>
        <img src={comment.img} alt="" />
      </div>
      <span className="date">{moment(comment.createdAt).fromNow()}</span>
      <div className="inf">
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
        <div className="item" onClick={() => setComOfComOpen(!comOfComOpen)}>
          <ModeCommentOutlinedIcon />
          {data ? data[1].length : 0} Comments
        </div>
      </div>
      {comOfComOpen && <ComOfCom commentId={comment.id} key={comment.id} />}
    </div>
  );
};

export default Comment;
