<script>
  import { GetCpuData } from "../wailsjs/go/main/App";
  
  let message = "Nothing to display"

  function getData() {
    GetCpuData().then((data) => (message = `Model: ${data.Model} -- Cores: ${data.CoreCount}`));
  }
  </script>

<p>{message}</p>

<button on:click={getData}>Get Cpu Data</button>