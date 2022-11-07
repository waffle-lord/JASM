<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import { refreshRate as storeRefreshRate } from "./stores/generalStore";
  import { RefreshRate } from "./data/refreshRate";
  import { GetCpuData } from "../wailsjs/go/main/App";
  import Header from "./components/Header.svelte";

  let dataReady = false;
  let cpuInfo = "N/A";
  let coreLoads = [];
  let totalLoad;

  //
  function precision(a) {
    if (!isFinite(a)) return 0;
    var e = 1,
      p = 0;
    while (Math.round(a * e) / e !== a) {
      e *= 10;
      p++;
    }
    return p;
  }

  async function updateAllData() {
    let data = await GetCpuData();

    cpuInfo = `${data.Model} | Cores Count: ${data.CoreCount}`;

    totalLoad = precision(data.TotalLoad);

    coreLoads = data.CoreLoads;
  }

  // update data after DOM is rendered
  let refreshCallback;
  let unsubscribe = storeRefreshRate.subscribe((refreshRate) => {
    clearInterval(refreshCallback);

    if (refreshRate === RefreshRate.NONE) return;

    refreshCallback = setInterval(async () => {
      await updateAllData();
    }, refreshRate);
  });
  onMount(async () => {
    await updateAllData();
    dataReady = true;
  });
  onDestroy(unsubscribe);
</script>

<Header />
<main>
  {#if !dataReady}
    <p>Starting up ...</p>
  {:else}
    <p>{cpuInfo}</p>
    <div class="stat-item" />
    <div class="total-cpu">
      <p>Total Load :</p>
      <div class="radial-progress" style={`--value:${totalLoad};`}>{totalLoad}%</div>
    </div>
    <div class="individual-cpu">
      {#each coreLoads as load, i}
      <div class="stats shadow">
        <div class="stat">
          <div class="stat-title">Core #{i + 1}</div>
          <div class="stat-value">{precision(load)}%</div>
          <div class="stat-desc">Load</div>
        </div>
      </div>
    {/each}
    </div>
  {/if}
</main>

<style>
  main {
    @apply pt-6 flex flex-col;
  }

  .total-cpu {
    @apply self-center;
  }
  .radial-progress {
    @apply text-orange-400;
  }

  .individual-cpu {
    @apply w-full grid grid-cols-4 gap-0;
  }
  .stats {
    @apply justify-center;
  }
  .stat {
    @apply w-40;
  }
</style>
