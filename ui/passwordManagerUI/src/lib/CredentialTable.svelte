<script lang="ts">
  import PasswordCell from './PasswordCell.svelte';
  import type {CredentialRow} from '../App.svelte';

  export let credentials: CredentialRow[] = [];
  export let activeHistoryId = '';
  export let onToggleEditOrSave: (index: number) => void = () => {};
  export let onCopyPassword: (index: number) => void = () => {};
  export let onToggleHistory: (index: number) => void = () => {};
</script>

<div class="table-wrap primary-table">
  <table>
    <thead>
      <tr>
        <th>Id</th>
        <th>Domain</th>
        <th>User</th>
        <th>Password</th>
        <th>Actions</th>
      </tr>
    </thead>
    <tbody>
      {#each credentials as row, index}
        <tr>
          <td>
            {row.id}
          </td>
          <td>
            {#if row.editing}
              <input class="input-field" bind:value={row.domain} />
            {:else}
              {row.domain}
            {/if}
          </td>
          <td>
            {#if row.editing}
              <input class="input-field" bind:value={row.username} />
            {:else}
              {row.username}
            {/if}
          </td>
          <td>
            <PasswordCell
              password={row.password}
              editing={row.editing}
              copied={row.copied}
              bind:bindValue={row.password}
              onCopy={() => onCopyPassword(row.id)}
            />
          </td>
          <td>
            <div class="actions">
              <button
                type="button"
                class="action-button history-button {activeHistoryId === String(row.id) ? 'active' : ''}"
                title="History"
                on:click={() => onToggleHistory(row.id)}
              >
                🕘 History
              </button>
              <button type="button" class="action-button edit-button" on:click={() => onToggleEditOrSave(row.id)}>
                {row.editing ? 'Save' : 'Edit'}
              </button>
            </div>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>

<style>
  .table-wrap {
    overflow-x: auto;
  }

  .primary-table {
    max-height: 414px;
    overflow-y: auto;
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

  tbody tr {
    transition: background-color 140ms ease;
  }

  tbody tr:hover {
    background: #0f2d1b4f;
  }

  .actions {
    display: flex;
    align-items: center;
    gap: 0.45rem;
  }

  .action-button {
    border: 1px solid #58ff8752;
    border-radius: 8px;
    padding: 0.38rem 0.64rem;
    background: #0a1910;
    color: #95ffbb;
    font-size: 0.82rem;
    cursor: pointer;
  }

  .action-button:hover {
    color: #dcffea;
    border-color: #72ffa5;
    box-shadow: 0 0 10px #4dff8152;
  }

  .history-button.active {
    color: #dcffea;
    border-color: #72ffa5;
    box-shadow: 0 0 10px #4dff8152;
    background: #103120;
  }

  .input-field {
    width: 100%;
    min-width: 130px;
    border: 1px solid #59ff8875;
    border-radius: 6px;
    background: #020703;
    color: #beffd6;
    padding: 0.38rem 0.45rem;
    outline: none;
  }

  .input-field:focus {
    border-color: #8fffb8;
    box-shadow: 0 0 0 2px #49ff7d26;
  }
</style>
