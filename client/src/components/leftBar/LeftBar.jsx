import "./leftBar.scss";
import Friends from "../../assets/friends.png";
import Groups from "../../assets/group.png";
import AddFriend from "../../assets/add-friend.png";
import { AuthContext } from "../../context/authContext";
import { useContext } from "react";

const LeftBar = () => {
  const {currentUser} = useContext(AuthContext);




  return (
    <div className="leftBar">
      <div className="container">
        <div className="menu">
          <div className="user">
            <img
            src={currentUser.profilePic} 
            alt="" 
          /> 
            <span>{currentUser.firstname} {currentUser.lastname}</span>
          </div>
          <div className="item">
          <img src={Friends} alt="" /> 
            <span>Friends</span>
          </div>
          <div className="item">
            <img src={AddFriend} alt="" />
            <span>Add friends</span>
          </div>
          <div className="item">
            <img src={Groups} alt="" />
            <span>Groups</span>
          </div>
        </div>
        <hr />
      </div>
    </div>
  );
};
export default LeftBar;
