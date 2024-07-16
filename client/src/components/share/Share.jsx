import "./share.scss";
import { AuthContext } from "../../context/authContext";
import { useContext, useState } from "react";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { makeRequest } from "../../axios";

const Share = () => {
  const { currentUser } = useContext(AuthContext);

  const [desc, setDesc] = useState("");
  const [file, setFile] = useState(null);
  const [message, setMessage] = useState("");

  const queryClient = useQueryClient();

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

  const mutation = useMutation(
    (newPost) => {
      return makeRequest.post("/posts", newPost);
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries(["posts"]);
      },
    }
  );

  const handleClick = async (e) => {
    e.preventDefault();
    let imgUrl = "";
    if (file) imgUrl = await upload();
    if ((file || desc) && !message) {
      console.log(imgUrl);
      mutation.mutate({ desc, img: imgUrl });
      setDesc("");
      setFile(null);
    }
  };

  return (
    <div className="share">
      <form>
        <div className="inputs">
          <div className="top">
            <div className="left">
              <img src={currentUser.profilePic} alt="" />
              <textarea
                value={desc}
                type="text"
                name="desc"
                placeholder={`What's on your mind, ${currentUser.firstname}?`}
                onChange={(e) => setDesc(e.target.value)}
              />
            </div>
            <div className="right">
              {file && (
                <img src={URL.createObjectURL(file)} alt="" className="file" />
              )}
            </div>
          </div>
          <hr />
          <div className="bottom">
            <label htmlFor="fileInput" style={{ cursor: "pointer" }}>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                height="24px"
                viewBox="0 0 24 24"
                width="24px"
                fill="#000000"
              >
                <path d="M0 0h24v24H0V0z" fill="none" />
                <path d="M18 20H4V6h9V4H4c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2v-9h-2v9zm-7.79-3.17l-1.96-2.36L5.5 18h11l-3.54-4.71zM20 4V1h-2v3h-3c.01.01 0 2 0 2h3v2.99c.01.01 2 0 2 0V6h3V4h-3z" />
              </svg>
              <span>Photo </span>
            </label>
            {message && message}
            <input
              name="img"
              type="file"
              id="fileInput"
              accept="image/*"
              style={{ display: "none" }}
              onChange={handleFileChange}
            />
            <button onClick={handleClick}>Send</button>
          </div>
        </div>
      </form>
    </div>
  );
};

export default Share;
