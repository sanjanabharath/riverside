import { useEffect, useState } from "react";
import axios from "axios";

function App() {
  const [task, setTask] = useState("");
  const [description, setDescription] = useState("");
  const [allTasks, setAllTasks] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const response = await axios.get("http://localhost:8080/getTasks");

        setAllTasks(response.data);
      } catch (e) {
        console.log(e);
      } finally {
        setLoading(false);
      }
    };
    fetchTasks();
  }, []);

  const createTask = async () => {
    try {
      const response = await axios.post("http://localhost:8080/createTask", {
        task: task,
        description: description,
      });

      console.log("Created:", response.data);

      // Clear inputs after success
      setTask("");
      setDescription("");
    } catch (err) {
      console.error("Failed to create task:", err);
    }
  };

  return (
    <div style={{ padding: "20px" }}>
      <h2>Create Task</h2>

      <input
        type="text"
        placeholder="Task Name"
        value={task}
        onChange={(e) => setTask(e.target.value)}
      />

      <br />
      <br />

      <input
        type="text"
        placeholder="Description"
        value={description}
        onChange={(e) => setDescription(e.target.value)}
      />

      <br />
      <br />

      <button onClick={createTask}>Create Task</button>

      <h4>Current pending tasks are given below</h4>
      {loading && <p>loading...</p>}
      {!loading && allTasks && (
        <ul>
          {allTasks.map((task) => (
            <li key={task.ID}>
              {task.task} - {task.description}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default App;
