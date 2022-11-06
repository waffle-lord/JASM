<script>
  import { onMount } from "svelte";
  import { GetCpuData } from "../wailsjs/go/main/App";
  import { EventsOn } from "../wailsjs/runtime";

  let dataReady = false;
  let cpuInfo = "N/A";
  let coreLoads = [];
  let totalLoad;

  //
  function precision(a) {
    if (!isFinite(a)) return 0;
    var e = 1, p = 0;
    while (Math.round(a * e) / e !== a) { e *= 10; p++; }
    return p;
  }

  async function updateAllData() {
      let data = await GetCpuData();

      cpuInfo = `${data.Model} | Cores Count: ${data.CoreCount}`;

      totalLoad = precision(data.TotalLoad);

      coreLoads = data.CoreLoads;
  }

  // update data when backend fires dataRefresh event (maybe provide data through that event instead of retreieving it here?)
  EventsOn("dataRefresh", async () => {
    await updateAllData();
    dataReady = true;
  });

  // update data after DOM is rendered
  onMount(async () => {
    await updateAllData();
    dataReady = true;
  });
  </script>

{#if !dataReady}
<p>Starting up ...</p>
{:else}
  <p>{cpuInfo}</p>
  <p>Total Load: {totalLoad}%</p>
  {#each coreLoads as load, i}
    <p>Core #{i+1} :: Load: {precision(load)}%</p>
  {/each}
{/if}