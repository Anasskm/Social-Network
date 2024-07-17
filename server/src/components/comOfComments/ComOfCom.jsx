import { useContext, useState } from "react";
import { AuthContext } from "../../context/authContext";
import AddOutlinedIcon from "@mui/icons-material/AddOutlined";
import "./comOfCom.scss";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { makeRequest } from "../../axios";
import moment from "moment";
import MoreHorizTwoToneIcon from "@mui/icons-material/MoreHorizTwoTone";

const ComOfCom = ({ commentId }) => {
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
      return makeRequest.post("/comOfComs", newComment);
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries(["comOfComs", commentId]);
        queryClient.invalidateQueries(["comlikes", commentId]);
      },
    }
  );
  const { isLoading, error, data } = useQuery(["comOfComs", commentId], () =>
    makeRequest.get(`/comOfComs?commentId=${commentId}`).then((res) => res.data)
  );

  const handleClick = async (e) => {
    e.preventDefault();
    if ((file || desc) && !message) {
      let imgUrl = "";
      if (file) imgUrl = await upload();
      mutation.mutate({ desc, img: imgUrl, commentId });
      setDesc("");

      setFile(null);
    }
  };

  return (
    <div className="comOfComs">
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
        <label htmlFor={`fileInput_${commentId}`} style={{ cursor: "pointer" }}>
          <AddOutlinedIcon />
        </label>
        <input
          type="file"
          id={`fileInput_${commentId}`}
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
        ? data.map((comment) => (
            <div className="comOfCom">
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
              <span className="date">
                {moment(comment.createdAt).fromNow()}
              </span>
            </div>
          ))
        : ""}
    </div>
  );
};

export default ComOfCom;
