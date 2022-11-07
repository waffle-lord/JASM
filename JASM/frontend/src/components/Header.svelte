<script lang="ts">
  import { refreshRate } from "../stores/generalStore";
  import { RefreshRate } from "../data/refreshRate";
  const refreshRecord = Object.keys(RefreshRate)
    .map((key) => ({ key, value: RefreshRate[key] }))
    .filter((elt) => !isNaN(elt.value));
</script>

<div class="navbar bg-base-100">
  <div class="flex-1">
    <a class="btn btn-ghost normal-case text-xl">JASM</a>
  </div>
  <div class="flex-none">
    <ul class="menu menu-horizontal p-0">
      <li>
        <div class="form-control">
          <select class="select select-bordered w-full max-w-xs" bind:value={$refreshRate}>
            <option disabled>Refresh rate</option>
            {#each refreshRecord as rate}
              <option value={rate.value}>{rate.key}</option>
            {/each}
          </select>
        </div>
      </li>
    </ul>
  </div>
</div>

<style>
  .navbar {
    @apply shadow-lg;
  }
</style>
