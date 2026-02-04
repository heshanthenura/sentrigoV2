import {
  Shield,
  LayoutDashboard,
  Radar,
  Settings,
  ShieldHalf,
  TriangleAlert,
  LogOut,
  PanelLeftOpen,
  PanelRightOpen,
} from "lucide-react";
import { useEffect, useState } from "react";

function NavBar({ onSelect }: { readonly onSelect: (data: string) => void }) {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const [selected, setSelected] = useState("dashboard");

  useEffect(() => {
    onSelect(selected);
  }, [selected]);

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  return (
    <div
      className={`flex flex-col border-[1px] border-r-[#1E293B] py-[10px] transition-all duration-200 ease-in-out ${
        isMenuOpen ? "px-[20px]" : "px-[12px]"
      }`}
    >
      <div
        className={`flex transition-all duration-200 ease-in-out ${
          isMenuOpen
            ? "items-center justify-between gap-[10px]"
            : "flex-col items-center gap-[8px]"
        }`}
      >
        <div className="flex items-center gap-[6px]">
          <div className="bg-[#137FEC] text-white p-1 rounded-md">
            <Shield />
          </div>
          {isMenuOpen && (
            <div>
              <div className="text-white font-bold">SentriGO V2</div>
              <div className="text-[#9DA6B4] text-sm">
                Intrusion Prevention System
              </div>
            </div>
          )}
        </div>

        <button className="text-[#9DA6B4]" onClick={toggleMenu}>
          {isMenuOpen ? <PanelRightOpen /> : <PanelLeftOpen />}
        </button>
      </div>
      <div
        className={`flex flex-col gap-[20px] mt-[40px] text-[#9DA6B4] transition-all duration-200 ease-in-out ${
          isMenuOpen ? "items-start" : "items-center"
        }`}
      >
        <button
          className={`flex items-center w-full p-[5px] rounded-md gap-[10px] ${selected === "dashboard" ? "bg-blue-900 text-blue-600" : ""}`}
          onClick={() => setSelected("dashboard")}
        >
          <LayoutDashboard />
          {isMenuOpen && <div className="font-regular ">Dashboard</div>}
        </button>

        <button
          className={`flex items-center w-full p-[5px] rounded-md gap-[10px] ${selected === "network" ? "bg-blue-900 text-blue-600" : ""}`}
          onClick={() => setSelected("network")}
        >
          <Radar />
          {isMenuOpen && <div className="font-regular">Network Monitor</div>}
        </button>

        <button
          className={`flex items-center w-full p-[5px] rounded-md gap-[10px] ${selected === "settings" ? "bg-blue-900 text-blue-600" : ""}`}
          onClick={() => setSelected("settings")}
        >
          <Settings />
          {isMenuOpen && <div className="font-regular">System Settings</div>}
        </button>

        <button
          className={`flex items-center w-full p-[5px] rounded-md gap-[10px] ${selected === "rules" ? "bg-blue-900 text-blue-600" : ""}`}
          onClick={() => setSelected("rules")}
        >
          <ShieldHalf />
          {isMenuOpen && <div className="font-regular">Detection Rules</div>}
        </button>

        <button
          className={`flex items-center w-full p-[5px] rounded-md gap-[10px] ${selected === "alerts" ? "bg-blue-900 text-blue-600" : ""}`}
          onClick={() => setSelected("alerts")}
        >
          <TriangleAlert />
          {isMenuOpen && <div className="font-regular">Alert Logs</div>}
        </button>
      </div>

      <div className="flex-1"></div>

      <div
        className={`bg-green-900 flex items-center p-[5px] mb-[20px] rounded-md gap-[10px] transition-all duration-200 ease-in-out ${
          isMenuOpen ? "" : "justify-center"
        }`}
      >
        <div className="bg-green-500 h-[10px] w-[10px] rounded-full"></div>
        {isMenuOpen && <div className="text-green-500">System Online</div>}
      </div>

      <div
        className={`border-t-[1px] border-t-[#1E293B] flex py-[10px] text-[#9DA6B4] transition-all duration-200 ease-in-out ${
          isMenuOpen ? "items-center" : "justify-center"
        }`}
      >
        {isMenuOpen && <div className="text-white">Admin</div>}
        {isMenuOpen && <div className="flex-1"></div>}

        <LogOut />
      </div>
    </div>
  );
}

export default NavBar;
