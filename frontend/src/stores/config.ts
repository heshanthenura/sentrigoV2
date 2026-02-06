import { writable } from "svelte/store";
export const systemStatus = writable({
  isOnline: false,
  isRunning: false,
});

export const captureConfig = writable({
  iface_name: "",
  snapshot_len: 65535,
  promiscuous: false,
  timeout: -1,
});
