import axios from "axios";
import type { StartParameters } from "../types/types";
import { systemStatus } from "../stores";

export async function startCapture(captureConfig: StartParameters) {
  try {
    const response = await axios.post(
      import.meta.env.VITE_API_URL + "/capture/start",
      captureConfig,
    );
    console.log(response.data);
    if (response.status === 200) {
      systemStatus.set({ ...systemStatus, isRunning: true });
    }
  } catch (error) {
    console.error("Error starting capture:", error);
  }
}

export async function stopCapture() {
  try {
    const response = await axios.post(
      import.meta.env.VITE_API_URL + "/capture/stop",
    );
    if (response.status === 200) {
      systemStatus.set({ ...systemStatus, isRunning: false });
    }
  } catch (error) {
    console.error("Error stopping capture:", error);
  }
}
