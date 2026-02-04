import { useState } from "react";
import "./App.css";
import NavBar from "./components/NavBar";

function App() {
  const [selected, setSelected] = useState("dashboard");

  const handleSelectedChange = (data: string) => {
    setSelected(data);
  };

  return (
    <div className="bg-[#101922] h-screen w-screen flex">
      <NavBar onSelect={handleSelectedChange} />
      <div className="bg-[#101922] flex-1">sds</div>
    </div>
  );
}

export default App;
