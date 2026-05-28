<script lang="ts">
  import type { CredentialRowReq } from '../App.svelte';

  let {
    open = false,
    submitting = false,
    onClose = () => {},
    onSubmit = () => {},
  }: {
    open?: boolean;
    submitting?: boolean;
    onClose?: () => void;
    onSubmit?: (credential: CredentialRowReq) => void | Promise<void>;
  } = $props();

  let domain = $state('');
  let username = $state('');
  let password = $state('');
  let showPassword = $state(false);

  const resetForm = () => {
    domain = '';
    username = '';
    password = '';
    showPassword = false;
  };

  const handleClose = () => {
    if (submitting) return;
    resetForm();
    onClose();
  };

  const handleBackdropKeydown = (event: KeyboardEvent) => {
    if (event.key === 'Escape') handleClose();
  };

  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();
    if (submitting) return;

    await onSubmit({
      id: 0,
      domain: domain.trim(),
      username: username.trim(),
      password,
    });
  };

  $effect(() => {
    if (!open) resetForm();
  });
</script>

{#if open}
  <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
  <div
    class="modal-backdrop"
    role="presentation"
    onclick={(event) => event.target === event.currentTarget && handleClose()}
    onkeydown={handleBackdropKeydown}
  >
    <div class="modal-card" role="dialog" aria-modal="true" aria-labelledby="add-credential-title">
      <header class="modal-header">
        <div>
          <p class="eyebrow">New entry</p>
          <h2 id="add-credential-title">Add credential</h2>
        </div>
        <button type="button" class="close-button" title="Close" onclick={handleClose} disabled={submitting}>
          ×
        </button>
      </header>

      <form class="modal-form" onsubmit={handleSubmit}>
        <table>
          <thead>
            <tr>
              <th>Domain</th>
              <th>User</th>
              <th>Password</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>
                <input
                  class="input-field"
                  type="text"
                  bind:value={domain}
                  placeholder="example.com"
                  required
                  disabled={submitting}
                />
              </td>
              <td>
                <input
                  class="input-field"
                  type="text"
                  bind:value={username}
                  placeholder="username"
                  required
                  disabled={submitting}
                />
              </td>
              <td>
                <div class="password-field-wrap">
                  <input
                    class="input-field password-input"
                    type={showPassword ? 'text' : 'password'}
                    bind:value={password}
                    placeholder="••••••••"
                    required
                    disabled={submitting}
                  />
                  <button
                    type="button"
                    class="icon-button"
                    onclick={() => (showPassword = !showPassword)}
                    aria-label={showPassword ? 'Hide password' : 'Show password'}
                    title={showPassword ? 'Hide password' : 'Show password'}
                    disabled={submitting}
                  >
                    {showPassword ? '🙈' : '👁'}
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>

        <footer class="modal-actions">
          <button type="button" class="action-button" onclick={handleClose} disabled={submitting}>
            Cancel
          </button>
          <button type="submit" class="action-button primary-button" disabled={submitting}>
            {submitting ? 'Saving…' : 'Create credential'}
          </button>
        </footer>
      </form>
    </div>
  </div>
{/if}

<style>
  .modal-backdrop {
    position: fixed;
    inset: 0;
    z-index: 900;
    display: grid;
    place-items: center;
    padding: 1.25rem;
    background: #010402bf;
    backdrop-filter: blur(3px);
  }

  .modal-card {
    width: min(760px, 100%);
    border: 1px solid #3cff7540;
    border-radius: 16px;
    background: #030a06f2;
    box-shadow: 0 0 16px #26ff6e2b, 0 0 42px #0d9f3f1f;
    overflow: hidden;
  }

  .modal-header {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    gap: 1rem;
    padding: 1.1rem 1.25rem;
    border-bottom: 1px solid #35ff7040;
    background: linear-gradient(180deg, #0c1b11a8, #08140ea8);
  }

  .eyebrow {
    margin-bottom: 0.2rem;
    color: #79ff9e;
    font-size: 0.72rem;
    letter-spacing: 0.12em;
    text-transform: uppercase;
    font-weight: 700;
  }

  h2 {
    font-size: 1.05rem;
    font-weight: 700;
    color: #caffda;
    text-shadow: 0 0 10px #4bff8a52;
  }

  .close-button {
    border: 1px solid #58ff8752;
    border-radius: 8px;
    width: 2rem;
    height: 2rem;
    background: #0a1910;
    color: #95ffbb;
    font-size: 1.2rem;
    line-height: 1;
    cursor: pointer;
  }

  .close-button:hover:not(:disabled) {
    color: #dcffea;
    border-color: #72ffa5;
    box-shadow: 0 0 10px #4dff8152;
  }

  .close-button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .modal-form {
    padding: 1rem 1.25rem 1.25rem;
  }

  table {
    width: 100%;
    border-collapse: collapse;
  }

  th,
  td {
    text-align: left;
    padding: 0.55rem 0.45rem;
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
  }

  .input-field {
    width: 100%;
    min-width: 0;
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

  .input-field:disabled {
    opacity: 0.65;
  }

  .password-field-wrap {
    display: flex;
    align-items: center;
    gap: 0.45rem;
  }

  .password-input {
    flex: 1;
  }

  .icon-button {
    width: 30px;
    height: 30px;
    border-radius: 8px;
    border: 1px solid #58ff8791;
    background: #0a1a11;
    color: #8effb5;
    cursor: pointer;
    font-size: 0.95rem;
    display: grid;
    place-items: center;
    flex-shrink: 0;
  }

  .icon-button:hover:not(:disabled) {
    color: #d3ffe3;
    box-shadow: 0 0 12px #48ff7d54;
  }

  .icon-button:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .modal-actions {
    display: flex;
    justify-content: flex-end;
    gap: 0.55rem;
    margin-top: 1rem;
  }

  .action-button {
    border: 1px solid #58ff8752;
    border-radius: 8px;
    padding: 0.45rem 0.85rem;
    background: #0a1910;
    color: #95ffbb;
    font-size: 0.84rem;
    cursor: pointer;
  }

  .action-button:hover:not(:disabled) {
    color: #dcffea;
    border-color: #72ffa5;
    box-shadow: 0 0 10px #4dff8152;
  }

  .action-button:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .primary-button {
    background: linear-gradient(180deg, #0f2919, #0a1c12);
    color: #adffca;
    box-shadow: inset 0 0 10px #3cff7638;
  }

  .primary-button:hover:not(:disabled) {
    color: #d7ffe7;
    border-color: #79ffab;
    box-shadow: inset 0 0 14px #4aff823b, 0 0 14px #4aff8233;
  }
</style>
