<script lang="ts">
  import NavBar from "./components/NavBar.svelte";
  import DashboardTab from "./tabs/DashboardTab.svelte";
  import SettingsTab from "./tabs/SettingsTab.svelte";

  let selected = "dashboard";

  function handleSelect(tab: string) {
    selected = tab;
  }
</script>

<div class="bg-[#101922] h-screen w-screen flex">
  <NavBar onselect={handleSelect} />
  <div class="bg-[#101922] flex-1">
    {#if selected === "dashboard"}
      <DashboardTab />
    {:else if selected === "network"}
      <div class="text-white p-8">Network Monitor</div>
    {:else if selected === "settings"}
      <SettingsTab />
    {:else if selected === "rules"}
      <div class="text-white p-8">Detection Rules</div>
    {:else if selected === "alerts"}
      <div class="text-white p-8">Alert Logs</div>
    {:else}
      <div class="text-white p-8">Dashboard</div>
    {/if}
  </div>
</div>
