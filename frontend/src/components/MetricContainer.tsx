import { useEffect, useState } from "react";

function MetricContainer() {
  const [metrics, setMetrics] = useState<any>(null);

  const getWebSocketURL = () => {
    if (import.meta.env.DEV) {
      return "ws://192.168.1.101:8080/api/ws/metrics";
    } else {
      const protocol = window.location.protocol === "https:" ? "wss" : "ws";
      return `${protocol}://${window.location.host}/api/ws/metrics`;
    }
  };

  useEffect(() => {
    const ws = new WebSocket(getWebSocketURL());

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setMetrics(data);
      console.log(data);
    };

    ws.onerror = (e) => console.log("WebSocket error:", e);
    ws.onclose = () => console.log("WebSocket closed");

    return () => ws.close(); // cleanup on unmount
  }, []);

  return (
    <div>
      <h2>Metrics</h2>
      <pre>{metrics && JSON.stringify(metrics, null, 2)}</pre>
    </div>
  );
}

export default MetricContainer;
