import React, {useState} from "react";

function App() {
  const [progress, setProgress] = useState(0)
  const getProgress = () => {
    const evtSource = new EventSource("http://localhost:9999/api/event")
    evtSource.addEventListener("sse-message", function (e) {
      setProgress(JSON.parse(e.data)["progress"])
    });
    evtSource.addEventListener("error", function (e) {
      e.preventDefault()
      evtSource.close()
    });
  }
  return (
    <div className="App">
      <div>
        Progress App {progress * 10}
      </div>
      <button onClick={getProgress}>getProgress</button>
    </div>
  );
}

export default App;
