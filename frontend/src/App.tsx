import { useState } from "react";
import "./App.css";
import NavBar from "./components/NavBar";
import SettingsTab from "./tabs/SettingsTab";

function App() {
  const [selected, setSelected] = useState("dashboard");

  const handleSelectedChange = (data: string) => {
    setSelected(data);
  };

  const renderTab = () => {
    switch (selected) {
      case "dashboard":
        return "dashboard";
      case "network":
        return "network";
      case "settings":
        return <SettingsTab />;
      default:
        return "dashboard";
    }
  };
  return (
    <div className="bg-[#101922] h-screen w-screen flex">
      <NavBar onSelect={handleSelectedChange} />
      <div className="bg-[#101922] flex-1">{renderTab()}</div>
    </div>
  );
}

export default App;
