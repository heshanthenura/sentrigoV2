import { useEffect, useState } from "react";
import { NavLink } from "react-router-dom";

function NavBar() {
  const [isHealthy, setIsHealthy] = useState(false);

  const checkSystemHealth = () => {
    fetch("/api/health")
      .then((response) => response.json())
      .then((data) => {
        setIsHealthy(data.status === "ok");
      })
      .catch(() => {
        setIsHealthy(false);
      });
  };

  useEffect(() => {
    checkSystemHealth();
  }, []);
  setInterval(checkSystemHealth, 60000);

  return (
    <nav className="w-full bg-zinc-900 border-b border-zinc-800">
      <div className="max-w-7xl mx-auto px-6 h-14 flex items-center justify-between">
        <div className="flex items-center gap-3">
          <div className="w-2.5 h-2.5 rounded-full bg-emerald-500 animate-pulse" />
          <h1 className="text-lg font-semibold text-white tracking-wide">
            SentriGOV2
          </h1>
        </div>

        <div className="flex items-center gap-6 text-sm">
          <NavLink
            to="/"
            className={({ isActive }) =>
              `transition ${
                isActive ? "text-emerald-400" : "text-zinc-400 hover:text-white"
              }`
            }
          >
            Dashboard
          </NavLink>

          <NavLink
            to="/metrics"
            className={({ isActive }) =>
              `transition ${
                isActive ? "text-emerald-400" : "text-zinc-400 hover:text-white"
              }`
            }
          >
            Metrics
          </NavLink>

          <NavLink
            to="/alerts"
            className={({ isActive }) =>
              `transition ${
                isActive ? "text-emerald-400" : "text-zinc-400 hover:text-white"
              }`
            }
          >
            Alerts
          </NavLink>

          <NavLink
            to="/settings"
            className={({ isActive }) =>
              `transition ${
                isActive ? "text-emerald-400" : "text-zinc-400 hover:text-white"
              }`
            }
          >
            Settings
          </NavLink>
        </div>

        <div className="flex items-center gap-2 text-xs">
          {isHealthy ? (
            <>
              <span className="w-2 h-2 rounded-full bg-green-500" />
              <span className="text-green-400">System Healthy</span>
            </>
          ) : (
            <>
              <span className="w-2 h-2 rounded-full bg-red-500" />
              <span className="text-red-400">System Down</span>
            </>
          )}
        </div>
      </div>
    </nav>
  );
}

export default NavBar;
