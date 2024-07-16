import { useEffect, useState } from "react";
import "./loading.scss";
import { InfinitySpin } from "react-loader-spinner";
import { useNavigate } from "react-router-dom";
import HoneyComb from "../../components/honeyComb/HoneyComb";

const Loading = () => {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setLoading(true);
    setTimeout(() => {
      setLoading(false);
      navigate("/");
    }, 1500);
  }, []);

  return (
    <div className="loading">
      <HoneyComb/>
    </div>
  );
};

export default Loading;
