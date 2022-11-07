<script lang="ts">
  import { refreshRate } from "../stores/generalStore";
  import { RefreshRate } from "../data/refreshRate";
  import Dropdown from "./Dropdown.svelte";
  const refreshRecord = Object.keys(RefreshRate)
    .map((key) => ({ key, value: RefreshRate[key] }))
    .filter((elt) => !isNaN(elt.value));
</script>

<div class="navbar ">
  <div class="flex-1">
    <a id="brand">JASM</a>
  </div>
  <div class="flex-none">
    <Dropdown label={"Refresh rate"} options={refreshRecord} value={refreshRate} />
  </div>
</div>

<style>
  .navbar {
    @apply shadow-lg bg-base-100;
  }

  #brand {
    @apply btn btn-ghost normal-case text-xl;
  }
</style>
