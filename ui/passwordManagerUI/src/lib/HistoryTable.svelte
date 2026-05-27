<script lang="ts">
  import type {CredentialHistory} from '../App.svelte';

  type props = {
    selectedHistoryForId: string;
    credentialsHistory: CredentialHistory[];
    onClose?: () => void;
  }

  let {selectedHistoryForId, credentialsHistory, onClose = () => {}}: props = $props();
</script>

<section class="history-panel">
  <header class="history-header">
    <h2>History for {selectedHistoryForId}</h2>
    <button type="button" class="close-button" onclick={onClose} title="Hide history" aria-label="Hide history">
      Hide
    </button>
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
        {#if credentialsHistory.length > 0}
          {#each credentialsHistory as entry}
            <tr>
              <td>{entry.id}</td>
              <td>{entry.value}</td>
              <td>{entry.created_at}</td>
            </tr>
          {/each}
        {:else}
          <tr class="empty-row">
            <td colspan="3">No history entries for this credential.</td>
          </tr>
        {/if}
      </tbody>
    </table>
  </div>
</section>

<style>
  .history-panel {
    border-top: 1px solid #2eff6e3b;
    background: #051109f0;
  }

  .history-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 1rem;
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

  .close-button {
    border: 1px solid #58ff8752;
    border-radius: 8px;
    padding: 0.38rem 0.64rem;
    background: #0a1910;
    color: #95ffbb;
    font-size: 0.82rem;
    cursor: pointer;
    white-space: nowrap;
    flex-shrink: 0;
  }

  .close-button:hover {
    color: #dcffea;
    border-color: #72ffa5;
    box-shadow: 0 0 10px #4dff8152;
  }

  .history-table {
    height: 152px;
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

  .empty-row td {
    color: #79ff9acc;
    font-style: italic;
  }
</style>
