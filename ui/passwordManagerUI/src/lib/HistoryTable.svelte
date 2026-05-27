<script lang="ts">
  import type {CredentialHistory} from '../App.svelte';

  type props = {
    selectedHistoryForId: String;
    credentialsHistory: CredentialHistory[]
  }

  let {selectedHistoryForId, credentialsHistory}: props = $props();

</script>

<section class="history-panel">
  <header class="history-header">
    <h2>History for {selectedHistoryForId}</h2>
  </header>
  <div class="table-wrap history-table">
    <table>
      <thead>
        <tr>
          <th>Id</th>
          <th>Value</th>
          <th>Created At</th>
        </tr>
      </thead>
      <tbody>
        {#each credentialsHistory as entry}
          <tr>
            <td>{entry.id}</td>
            <td>{entry.value}</td>
            <td>{entry.created_at}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</section>

<style>
  .history-panel {
    border-top: 1px solid #2eff6e3b;
    background: #051109b0;
  }

  .history-header {
    padding: 0.9rem 1.4rem 0.65rem;
  }

  .history-header h2 {
    font-size: 0.86rem;
    letter-spacing: 0.08em;
    color: #98ffbb;
    text-transform: uppercase;
    font-weight: 700;
  }

  .table-wrap {
    overflow-x: auto;
  }

  .history-table {
    max-height: 230px;
    overflow-y: auto;
    border-top: 1px solid #2eff6e24;
  }

  .table-wrap::-webkit-scrollbar {
    width: 10px;
    height: 10px;
  }

  .table-wrap::-webkit-scrollbar-thumb {
    background: #31ff7654;
    border-radius: 999px;
    border: 2px solid #06110a;
  }

  .table-wrap::-webkit-scrollbar-track {
    background: #06110a;
  }

  table {
    width: 100%;
    min-width: 900px;
    border-collapse: collapse;
  }

  th,
  td {
    text-align: left;
    padding: 0.9rem 1.1rem;
    border-bottom: 1px solid #37ff6f29;
    color: #bcffd4;
    font-size: 0.92rem;
  }

  th {
    color: #77ff9e;
    font-size: 0.76rem;
    letter-spacing: 0.09em;
    text-transform: uppercase;
    background: #0a170f;
    text-shadow: 0 0 8px #4dff8742;
    position: sticky;
    top: 0;
    z-index: 1;
  }

  tbody tr:hover {
    background: #0f2d1b4f;
  }
</style>
