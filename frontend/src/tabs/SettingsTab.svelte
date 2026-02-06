<script lang="ts">
  import axios from "axios";
  import { List, Network } from "lucide-svelte";
  import { onMount } from "svelte";
  import { captureConfig } from "../stores";
  import { startCapture } from "../lib/utils";
  let interfaces: any[] = [];

  async function getInterfaces() {
    try {
      const response = await axios.get(
        import.meta.env.VITE_API_URL + "/interfaces",
      );
      interfaces = response.data;
      if (response.data.length > 0) {
        captureConfig.update((current) => ({
          ...current,
          iface_name: response.data[0].Name,
        }));
      }
    } catch (error) {
      console.error("Error fetching interfaces:", error);
    }
  }

  async function handleApply() {
    await startCapture($captureConfig);
  }



  onMount(() => {
    getInterfaces();
  });
</script>

<div
  class="flex flex-col items-center px-[20px] gap-[30px] overflow-y-auto h-full"
>
  <div class="w-full max-w-[1000px]">
    <h1 class="text-white font-bold text-[40px]">System Settings</h1>
    <p class="text-gray-400 text-[16px]">
      Configure your intrusion detection parameters and secure network
      interfaces.
    </p>
  </div>

  <div
    class="w-full max-w-[1000px] flex flex-col bg-[#1E293B] rounded-lg text-white py-[30px] px-[25px] gap-[20px]"
  >
    <div class="flex gap-[12px] items-center">
      <Network class="text-blue-600" size={24} />
      <h1 class="font-bold text-[20px]">Network Interface Selection</h1>
    </div>
    <div class="flex flex-col gap-[10px]">
      <p class="text-gray-200 font-medium">Active Interface</p>
      <select
        name="iface_name"
        id="iface_name"
        bind:value={$captureConfig.iface_name}
        class="bg-[#101922] p-[10px] rounded-md text-white border border-[#334155]"
      >
        {#each interfaces as iface}
          <option value={iface.Name}>
            {iface.Name}
          </option>
        {/each}
      </select>
      <p class="text-gray-400 text-[12px]">
        Selecting an interface will restart the capture engine.
      </p>
    </div>
    <div class="flex gap-[20px] items-start">
      <div class="flex flex-col gap-[10px] flex-1">
        <p class="text-gray-200 font-medium">Promiscuous Mode</p>
        <label class="flex items-center gap-[10px] cursor-pointer">
          <input
            type="checkbox"
            name="promiscuous"
            id="promiscuousMode"
            bind:checked={$captureConfig.promiscuous}
            class="w-[18px] h-[18px] bg-[#101922] border border-[#334155] rounded cursor-pointer accent-blue-600"
          />
          <span class="text-gray-400 text-[14px]">
            Enable Promiscuous Mode
          </span>
        </label>
      </div>
      <div class="flex flex-col gap-[10px] flex-1">
        <label for="snaplen" class="text-gray-200 font-medium">
          Snaplen (Bytes)
        </label>
        <input
          type="number"
          id="snaplen"
          name="snapshot_len"
          bind:value={$captureConfig.snapshot_len}
          class="bg-[#101922] p-[10px] rounded-md text-white border border-[#334155]"
        />
      </div>
      <div class="flex flex-col gap-[10px] flex-1">
        <label for="timeout" class="text-gray-200 font-medium">
          Timeout (ms)
        </label>
        <input
          type="number"
          id="timeout"
          name="timeout"
          bind:value={$captureConfig.timeout}
          class="bg-[#101922] p-[10px] rounded-md text-white border border-[#334155]"
        />
      </div>
    </div>
  </div>

  <div
    class="w-full max-w-[1000px] flex flex-col bg-[#1E293B] rounded-lg text-white py-[30px] px-[25px] gap-[20px]"
  >
    <div class="flex justify-between">
      <div class="flex gap-[12px] items-center">
        <List class="text-blue-600" size={24} />
        <h1 class="font-bold text-[20px]">IP Whitelist</h1>
      </div>
      <button
        class="text-[16px] text-white font-semibold bg-blue-600 py-[5px] px-[10px] rounded-lg"
      >
        + Add New IP
      </button>
    </div>
    <table class="rounded-lg">
      <thead class="text-[#9DA6B4] bg-[#101922]">
        <tr>
          <th>IP ADDRESS</th>
          <th>LABEL</th>
          <th>ACTION</th>
        </tr>
      </thead>
    </table>
  </div>

  <div
    class="w-full max-w-[1000px] border-t-[1px] border-t-[#334155] py-[20px] flex justify-end"
  >
    <button
      class="text-[16px] text-white font-semibold bg-blue-600 py-[10px] px-[40px] rounded-lg"
      on:click={handleApply}
    >
      Apply & Restart Service
    </button>
  </div>
</div>
