<script>
  import { Play, Square } from "lucide-svelte";
  import { systemStatus, captureConfig } from "../stores";
  import { startCapture, stopCapture } from "../lib/utils";

  function toggleRunning() {
    if ($systemStatus.isRunning) {
      stopCapture();
    } else {
      startCapture($captureConfig);
    }
  }
</script>

<div
  class="flex flex-col items-center px-[20px] gap-[30px] overflow-y-auto h-full"
>
  <div class="w-full max-w-[1000px] flex justify-between items-center">
    <div>
      <h1 class="text-white font-bold text-[40px]">Dashboard</h1>
      <p class="text-gray-400 text-[16px] flex gap-[5px] items-center">
        <span
          class="w-[10px] h-[10px] rounded-full inline-block"
          class:bg-green-500={$systemStatus.isRunning}
          class:bg-red-500={!$systemStatus.isRunning}
        ></span><span class="font-semibold"
          >LIVE MONITORING {$systemStatus.isRunning
            ? "ACTIVE"
            : "INACTIVE"}</span
        >
      </p>
    </div>

    <div>
      <button
        class="font-semibold text-white flex gap-[5px] items-center px-[10px] py-[10px] rounded-lg"
        class:bg-red-600={$systemStatus.isRunning}
        class:bg-green-600={!$systemStatus.isRunning}
        on:click={toggleRunning}
      >
        {#if $systemStatus.isRunning}
          <Square size={16} />
        {:else}
          <Play size={16} />
        {/if}

        <span
          >{$systemStatus.isRunning
            ? "Stop Monitoring"
            : "Start Monitoring"}</span
        >
      </button>
    </div>
  </div>

  <div class="w-full max-w-[1000px]"></div>
</div>
