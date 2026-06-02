import { useEffect, useState } from "react";
import "./App.css";
import axios from "axios";

function App() {
  const [info, setInfo] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios
      .get("http://localhost:8080/home")
      .then((res) => {
        setInfo(res.data);
        setLoading(false);
      })
      .catch((err) => {
        setError(err);
        setLoading(false);
      });
  }, []);
  if (loading) return <div>Loading..</div>;

  if (error) return <div>Error!</div>;
  return <div>{info}</div>;
}

export default App;
