import axios from "axios";
import { List, Network } from "lucide-react";
import { useEffect, useState } from "react";

type startParameters = {
  iface_name: string;
  snapshot_len: number;
  promiscuous: boolean;
  timeout: number;
};

function SettingsTab() {
  const [interfaces, setInterfaces] = useState([]);
  const [startParams, setStartParams] = useState<startParameters>({
    iface_name: "",
    snapshot_len: 65535,
    promiscuous: false,
    timeout: -1,
  });

  const getInterfaces = async () => {
    try {
      const response = await axios.get(
        import.meta.env.VITE_API_URL + "/interfaces",
      );
      setInterfaces(response.data);
      if (response.data.length > 0) {
        setStartParams({ ...startParams, iface_name: response.data[0].Name });
      }
    } catch (error) {
      console.error("Error fetching interfaces:", error);
    }
  };

  const startCapture = async () => {
    try {
      const response = await axios.post(
        import.meta.env.VITE_API_URL + "/capture/start",
        startParams,
      );
      console.log(response.data);
    } catch (error) {
      console.error("Error starting capture:", error);
    }
  };

  useEffect(() => {
    getInterfaces();
  }, []);

  return (
    <div className="flex flex-col items-center px-[20px] gap-[30px] overflow-y-auto h-full">
      <div className="w-full max-w-2xl">
        <h1 className="text-white font-bold text-[40px]">System Settings</h1>
        <p className="text-gray-400 text-[16px]">
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
            name="iface_name"
            id="iface_name"
            value={startParams.iface_name}
            className="bg-[#101922] p-[10px] rounded-md text-white border border-[#334155]"
            onChange={(e) =>
              setStartParams({ ...startParams, iface_name: e.target.value })
            }
          >
            {interfaces.map((iface: any) => (
              <option key={iface.Name} value={`${iface.Name}`}>
                {iface.Name}
              </option>
            ))}
          </select>
          <p className="text-gray-400 text-[12px]">
            Selecting an interface will restart the capture engine.
          </p>
        </div>
        <div className="flex gap-[20px] items-start">
          <div className="flex flex-col gap-[10px] flex-1">
            <p className="text-gray-200 font-medium">Promiscuous Mode</p>
            <label className="flex items-center gap-[10px] cursor-pointer">
              <input
                type="checkbox"
                name="promiscuous"
                id="promiscuousMode"
                checked={startParams.promiscuous}
                onChange={(e) =>
                  setStartParams({
                    ...startParams,
                    promiscuous: e.target.checked,
                  })
                }
                className="w-[18px] h-[18px] bg-[#101922] border border-[#334155] rounded cursor-pointer accent-blue-600"
              />
              <span className="text-gray-400 text-[14px]">
                Enable Promiscuous Mode
              </span>
            </label>
          </div>
          <div className="flex flex-col gap-[10px] flex-1">
            <label htmlFor="snaplen" className="text-gray-200 font-medium">
              Snaplen (Bytes)
            </label>
            <input
              type="number"
              id="snaplen"
              name="snapshot_len"
              value={startParams.snapshot_len}
              onChange={(e) =>
                setStartParams({
                  ...startParams,
                  snapshot_len: Number.parseInt(e.target.value) || 0,
                })
              }
              className="bg-[#101922] p-[10px] rounded-md text-white border border-[#334155]"
            />
          </div>
          <div className="flex flex-col gap-[10px] flex-1">
            <label htmlFor="timeout" className="text-gray-200 font-medium">
              Timeout (ms)
            </label>
            <input
              type="number"
              id="timeout"
              name="timeout"
              value={startParams.timeout}
              onChange={(e) =>
                setStartParams({
                  ...startParams,
                  timeout: Number.parseInt(e.target.value) || 0,
                })
              }
              className="bg-[#101922] p-[10px] rounded-md text-white border border-[#334155]"
            />
          </div>
        </div>
      </div>

      <div className="w-full max-w-2xl flex flex-col bg-[#1E293B] rounded-lg text-white py-[30px] px-[25px] gap-[20px]">
        <div className="flex justify-between">
          <div className="flex gap-[12px] items-center">
            <List className="text-blue-600" size={24} />
            <h1 className="font-bold text-[20px]">IP Whitelist</h1>
          </div>
          <button className="text-[16px] text-white font-semibold bg-blue-600 py-[5px] px-[10px] rounded-lg">
            + Add New IP
          </button>
        </div>
        <table className="rounded-lg">
          <thead className="text-[#9DA6B4] bg-[#101922]">
            <tr>
              <th>IP ADDRESS</th>
              <th>LABEL</th>
              <th>ACTION</th>
            </tr>
          </thead>
        </table>
      </div>

      <div className="w-full max-w-2xl border-t-[1px] border-t-[#334155] py-[20px] flex justify-end">
        <button
          className="text-[16px] text-white font-semibold bg-blue-600 py-[10px] px-[40px] rounded-lg"
          onClick={startCapture}
        >
          Apply & Restart Service
        </button>
      </div>
    </div>
  );
}

export default SettingsTab;
