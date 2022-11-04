<script>
  import { GetCpuData } from "../wailsjs/go/main/App.js";
  
  let message = "No data to display";
  let data;

  function getData() {
    GetCpuData().then((result) => (data = result));
    message = `Model: ${data.Model} Core Count: ${data.CoreCount}`;
  }
</script>

<p>{message}</p>
<button on:click={getData}>Get Cpu Data</button>