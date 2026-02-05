<script lang="ts">
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
  } from "lucide-svelte";

  interface Props {
    onselect?: (tab: string) => void;
  }

  let { onselect }: Props = $props();

  let isMenuOpen = $state(false);
  let selected = $state("dashboard");

  $effect(() => {
    onselect?.(selected);
  });

  function toggleMenu() {
    isMenuOpen = !isMenuOpen;
  }

  function selectTab(tab: string) {
    selected = tab;
  }
</script>

<div
  class="flex flex-col border-[1px] border-r-[#1E293B] py-[10px] transition-all duration-200 ease-in-out {isMenuOpen
    ? 'px-[20px]'
    : 'px-[12px]'}"
>
  <div
    class="flex transition-all duration-200 ease-in-out {isMenuOpen
      ? 'items-center justify-between gap-[10px]'
      : 'flex-col items-center gap-[8px]'}"
  >
    <div class="flex items-center gap-[6px]">
      <div class="bg-[#137FEC] text-white p-1 rounded-md">
        <Shield />
      </div>
      {#if isMenuOpen}
        <div>
          <div class="text-white font-bold">SentriGO V2</div>
          <div class="text-[#9DA6B4] text-sm">Intrusion Prevention System</div>
        </div>
      {/if}
    </div>

    <button class="text-[#9DA6B4]" onclick={toggleMenu}>
      {#if isMenuOpen}
        <PanelRightOpen />
      {:else}
        <PanelLeftOpen />
      {/if}
    </button>
  </div>

  <div
    class="flex flex-col gap-[20px] mt-[40px] text-[#9DA6B4] transition-all duration-200 ease-in-out {isMenuOpen
      ? 'items-start'
      : 'items-center'}"
  >
    <button
      class="flex items-center w-full p-[5px] rounded-md gap-[10px] {selected ===
      'dashboard'
        ? 'bg-blue-900 text-blue-600'
        : ''}"
      onclick={() => selectTab("dashboard")}
    >
      <LayoutDashboard />
      {#if isMenuOpen}
        <div class="font-regular">Dashboard</div>
      {/if}
    </button>

    <button
      class="flex items-center w-full p-[5px] rounded-md gap-[10px] {selected ===
      'network'
        ? 'bg-blue-900 text-blue-600'
        : ''}"
      onclick={() => selectTab("network")}
    >
      <Radar />
      {#if isMenuOpen}
        <div class="font-regular">Network Monitor</div>
      {/if}
    </button>

    <button
      class="flex items-center w-full p-[5px] rounded-md gap-[10px] {selected ===
      'settings'
        ? 'bg-blue-900 text-blue-600'
        : ''}"
      onclick={() => selectTab("settings")}
    >
      <Settings />
      {#if isMenuOpen}
        <div class="font-regular">System Settings</div>
      {/if}
    </button>

    <button
      class="flex items-center w-full p-[5px] rounded-md gap-[10px] {selected ===
      'rules'
        ? 'bg-blue-900 text-blue-600'
        : ''}"
      onclick={() => selectTab("rules")}
    >
      <ShieldHalf />
      {#if isMenuOpen}
        <div class="font-regular">Detection Rules</div>
      {/if}
    </button>

    <button
      class="flex items-center w-full p-[5px] rounded-md gap-[10px] {selected ===
      'alerts'
        ? 'bg-blue-900 text-blue-600'
        : ''}"
      onclick={() => selectTab("alerts")}
    >
      <TriangleAlert />
      {#if isMenuOpen}
        <div class="font-regular">Alert Logs</div>
      {/if}
    </button>
  </div>

  <div class="flex-1"></div>

  <div
    class="bg-green-900 flex items-center p-[5px] mb-[20px] rounded-md gap-[10px] transition-all duration-200 ease-in-out {isMenuOpen
      ? ''
      : 'justify-center'}"
  >
    <div class="bg-green-500 h-[10px] w-[10px] rounded-full"></div>
    {#if isMenuOpen}
      <div class="text-green-500">System Online</div>
    {/if}
  </div>

  <div
    class="border-t-[1px] border-t-[#1E293B] flex py-[10px] text-[#9DA6B4] transition-all duration-200 ease-in-out {isMenuOpen
      ? 'items-center'
      : 'justify-center'}"
  >
    {#if isMenuOpen}
      <div class="text-white">Admin</div>
      <div class="flex-1"></div>
    {/if}
    <LogOut />
  </div>
</div>
