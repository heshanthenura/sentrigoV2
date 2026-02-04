import axios from "axios";
import { Network } from "lucide-react";
import { useEffect, useState } from "react";

function SettingsTab() {
  const [interfaces, setInterfaces] = useState([]);

  const getInterfaces = async () => {
    try {
      const response = await axios.get(
        import.meta.env.VITE_API_URL + "/interfaces",
      );
      setInterfaces(response.data);
    } catch (error) {
      console.error("Error fetching interfaces:", error);
    }
  };

  useEffect(() => {
    getInterfaces();
  }, []);

  return (
    <div className="flex flex-col items-center px-[20px] gap-[30px]">
      <div className="w-full max-w-2xl">
        <h1 className="text-white font-bold text-[40px]">System Settings</h1>
        <p className="text-[#9DA6B4] text-[16px]">
          Configure your intrusion detection parameters and secure network
          interfaces.
        </p>
      </div>

      <div className="w-full max-w-2xl flex flex-col bg-[#1E293B] rounded-lg text-white py-[30px] px-[25px] gap-[20px]">
        <div className="flex gap-[12px] items-center">
          <Network className="text-blue-600" size={24} />
          <h1 className="font-bold text-[20px]">Network Interface Selection</h1>
        </div>
        <div className="flex flex-col gap-[10px]">
          <p className="text-gray-200 font-medium">Active Interface</p>
          <select
            name="activeInterface"
            id="activeInterface"
            className="bg-[#101922] p-[10px] rounded-md text-white border border-[#334155]"
          >
            {interfaces.map((iface: any) => (
              <option key={iface.Name} value={`${iface.Name}`}>
                {iface.Name}
              </option>
            ))}
          </select>
          <p className="text-[#9DA6B4] text-[12px]">
            Selecting an interface will restart the capture engine.
          </p>
        </div>
      </div>

      <div className="w-full max-w-2xl border-t-[1px] border-t-[#334155] py-[20px] flex justify-end">
        <button className="text-[16px] text-white font-semibold bg-blue-600 py-[10px] px-[40px] rounded-lg">
          Apply & Restart Service
        </button>
      </div>
    </div>
  );
}

export default SettingsTab;
