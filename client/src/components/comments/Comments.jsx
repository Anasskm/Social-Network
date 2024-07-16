import { useContext, useState } from "react";
import { AuthContext } from "../../context/authContext";
import "./comments.scss";
import Comment from "../comment/Comment";
import AddOutlinedIcon from "@mui/icons-material/AddOutlined";
import { useQueryClient, useMutation, useQuery } from "react-query";
import { makeRequest } from "../../axios";

const Comments = ({ postId }) => {
  const { currentUser } = useContext(AuthContext);
  const [desc, setDesc] = useState("");
  const queryClient = useQueryClient();

  //--------------------------------------------------------------------UPLOAD IMG--------------------------------------------------------------------

  const [file, setFile] = useState(null);
  const [message, setMessage] = useState("");
  const handleFileChange = (e) => {
    const selectedFile = e.target.files[0];
    const validTypes = ["image/png", "image/jpeg", "image/gif", "image/webp"];

    if (selectedFile && validTypes.includes(selectedFile.type)) {
      setFile(selectedFile);
      setMessage("");
    } else {
      setFile(null);
      setTimeout(() => {
        setMessage("Please select a valid image file (PNG, JPEG, or GIF)");
        setMessage("");
      }, 2000);
    }
  };
  const upload = async () => {
    try {
      const formData = new FormData();
      formData.append("file", file);
      const res = await makeRequest.post("/upload", formData);
      return res.data;
    } catch (error) {
      console.log(error);
    }
  };

  //----------------------------------------------------FETCH COMMENTS--------------------------------------------------------------------

  const mutation = useMutation(
    (newComment) => {
      return makeRequest.post("/comments", newComment);
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries(["comments", postId]);
        queryClient.invalidateQueries(["likes", postId]);
      },
    }
  );

  const { isLoading, error, data } = useQuery(["comments", postId], () =>
    makeRequest.get(`/comments?postId=${postId}`).then((res) => res.data)
  );

  const handleClick = async (e) => {
    e.preventDefault();
    if ((file || desc) && !message) {
      let imgUrl = "";
      if (file) imgUrl = await upload();
      mutation.mutate({ desc, img: imgUrl, postId });
      setDesc("");
      setFile(null);
    }
  };

  return (
    <div className="comments" id={`comments_${postId}`}>
      <div className="write">
        <img src={currentUser.profilePic} alt="" />

        <input
          type="text"
          placeholder="Write a comment"
          value={desc}
          onChange={(e) => setDesc(e.target.value)}
        />
        <div className="preview">
          {file && (
            <img src={URL.createObjectURL(file)} alt="" className="file" />
          )}
        </div>
        <label htmlFor={`fileInput_${postId}`} style={{ cursor: "pointer" }}>
          <AddOutlinedIcon />
        </label>
        <input
          type="file"
          id={`fileInput_${postId}`}
          accept="image/*"
          style={{ display: "none" }}
          onChange={handleFileChange}
        />
        <button type="button" onClick={handleClick}>
          Send
        </button>
      </div>
      {isLoading
        ? "loading"
        : error
        ? "Something went wrong!"
        : data
        ? data.map((comment) => <Comment comment={comment} key={comment.id} />)
        : ""}
    </div>
  );
};

export default Comments;
