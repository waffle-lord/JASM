<script lang="ts">
  export let options;
  export let value;
  export let label;
</script>

<ul class="menu menu-horizontal p-0">
  <li>
    <div class="form-control">
      <select class="select select-bordered w-full max-w-xs" bind:value={$value}>
        <option disabled>{label}</option>
        {#each options as option}
          <option value={option.value}>{option.key}</option>
        {/each}
      </select>
    </div>
  </li>
</ul>
