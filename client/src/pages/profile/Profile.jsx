import "./profile.scss";
import FacebookIcon from "@mui/icons-material/Facebook";
import InstagramIcon from "@mui/icons-material/Instagram";
import EmailOutlinedIcon from "@mui/icons-material/EmailOutlined";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import Posts from "../../components/posts/Posts";
import Share from "../../components/share/Share";
import PlaceOutlinedIcon from "@mui/icons-material/PlaceOutlined";
import { useContext } from "react";
import { AuthContext } from "../../context/authContext";

const Profile = () => {
  const { currentUser } = useContext(AuthContext);
  return (
    <div className="profile">
      <div className="images">
        <img src={currentUser.profilePic} alt="" className="cover" />
        <img src={currentUser.coverPic} alt="" className="profilePic" />
      </div>
      <div className="profileContainer">
        <div className="uInfo">
          <div className="left">
            <a href="http://facebook.com" target="_blank">
              <FacebookIcon fontSize="large" />
            </a>
            <a href="http://instagram.com" target="_blank">
              <InstagramIcon fontSize="large" />
            </a>
          </div>
          <div className="center">
            <span>
              {currentUser.firstname} {currentUser.lastname}
            </span>
            <div className="info">
              <div className="item">
                <PlaceOutlinedIcon />
                <span>{currentUser.city}</span>
              </div>
            </div>
            <button>Follow</button>
          </div>
          <div className="right">
            <EmailOutlinedIcon />
            <MoreVertIcon />
          </div>
        </div>
        <Share />
        <Posts />
      </div>
    </div>
  );
};

export default Profile;
