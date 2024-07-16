import "./rightBar.scss";

const RightBar = () => {
  return (
    <div className="rightBar">
      <div className="container">
        <div className="item">
          <span>Suggestions For You</span>
          <div className="user">
            <div className="userInfo">
              <img
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS9YlEWl5uVeMF2q13gF6I8Auz2ZtUDRCSzHg&s"
                alt=""
              />
              <span>Hamid Zamel</span>
            </div>
            <div className="buttons">
              <button>Add</button>
              <button>Ignore</button>
            </div>
          </div>
          <div className="user">
            <div className="userInfo">
              <img
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS9YlEWl5uVeMF2q13gF6I8Auz2ZtUDRCSzHg&s"
                alt=""
              />
              <span>Hamid Zamel</span>
            </div>
            <div className="buttons">
              <button>Add</button>
              <button>Ignore</button>
            </div>
          </div>
        </div>
        <div className="item">
          <span>Online Friends</span>
          <div className="user">
            <div className="userInfo">
              <img
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS9YlEWl5uVeMF2q13gF6I8Auz2ZtUDRCSzHg&s"
                alt=""
              />
              <div className="online"/>
              <span>Hamid Zamel</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
export default RightBar;
