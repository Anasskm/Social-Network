import { useState } from "react";
import "./register.scss";
import { Link, useNavigate } from "react-router-dom";
import axios from "axios";

const Register = () => {
  const navigate = useNavigate();
  const [inputs, setInputs] = useState({
    email: "",
    username: "",
    password: "",
    firstname: "",
    lastname: "",
    dateOfBirth: "",
    city: "",
  });
  const handleChange = (e) => {
    setInputs((prev) => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const [err, serErr] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      console.log(await axios.post("http://localhost:8080/signup", inputs));
      navigate("/login");
    } catch (err) {
      serErr(err.response.data.error);
      console.log(err);
    }
  };
  return (
    <div className="register">
      <div className="card">
        <div className="left">
          <h1>Sign up</h1>
          <form onSubmit={handleSubmit}>
            {err && err}
            <input
              type="text"
              name="username"
              placeholder="Username"
              onChange={handleChange}
              required
            />
            <input
              type="email"
              name="email"
              placeholder="Email"
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
            <input
              type="text"
              name="firstname"
              placeholder="Firstname"
              onChange={handleChange}
              required
            />
            <input
              type="text"
              name="lastname"
              placeholder="Lastname"
              onChange={handleChange}
            />
            <input
              type="text"
              name="city"
              placeholder="City"
              onChange={handleChange}
              required
            />
            <input
              type="date"
              name="dateOfBirth"
              placeholder="Birthday"
              onChange={handleChange}
              required
            />

            <button type="submit">Sign up</button>
          </form>
        </div>
        <div className="right">
          <h1>Social Network</h1>

          <span>Do you have an acount?</span>
          <Link to="/login">
            <button>Login</button>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Register;
