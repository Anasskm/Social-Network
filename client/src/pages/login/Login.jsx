import { useContext, useEffect, useState } from "react";
import { AuthContext } from "../../context/authContext";
import "./login.scss";
import { Link, useNavigate } from "react-router-dom";
import axios from "axios";

const Login = () => {
  const navigate = useNavigate();
  const { login } = useContext(AuthContext);

  const [inputs, setInputs] = useState({
    username: "",
    password: "",
  });
  const [err, setErr] = useState(null);

  const handleChange = (e) => {
    setInputs((prev) => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const HandleLogin = async (e) => {
    e.preventDefault();
    try {
      await login(inputs);
      navigate("/l");
    } catch (error) {
      console.log(error);
      setErr(error.response.data.error);
    }
  };

  return (
    <div className="login">
      <div className="card">
        <div className="left">
          <h1>Social Network</h1>

          <span>Don't you have an acount?</span>
          <Link to="/register">
            <button>Sign up</button>
          </Link>
        </div>

        <div className="right">
          <h1>Login</h1>
          <form onSubmit={HandleLogin}>
            {err && err}
            <input
              type="text"
              name="username"
              placeholder="Username or Email"
              onChange={handleChange}
              required
            />
            <input
              type="password"
              name="password"
              placeholder="Password"
              onChange={handleChange}
              required
            />
            <button type="submit">Login</button>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Login;
