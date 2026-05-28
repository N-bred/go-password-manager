<script lang="ts">

  type Props = {
    password: string;
    editing: boolean;
    copied: boolean;
    onCopy: () => void;
    bindValue: string;
  }

  let {password, editing, copied, onCopy, bindValue = $bindable()} : Props = $props()

  const selectPasswordText = (event: MouseEvent) => {
    const target = event.currentTarget as HTMLInputElement | null;
    target?.select();
  };

</script>

<div class="password-cell">
  {#if editing}
    <input class="input-field" bind:value={bindValue} />
  {:else}
    <input
      class="password-pill"
      value={password}
      readonly
      onclick={selectPasswordText}
      title="Hover to reveal, click to select all"
    />
  {/if}

  <button type="button" class="icon-button" onclick={onCopy} aria-label="Copy password" title="Copy password">
    {#if copied}
      ✓
    {:else}
      ⧉
    {/if}
  </button>
</div>

<style>
  .password-cell {
    display: flex;
    align-items: center;
    gap: 0.45rem;
  }

  .password-pill {
    display: inline-flex;
    width: 100%;
    min-width: 170px;
    border: 1px solid #58ff8752;
    border-radius: 999px;
    background: #0b1e13;
    color: #79ffab;
    padding: 0.3rem 0.72rem;
    font-size: 0.84rem;
    letter-spacing: 0.05em;
    box-shadow: inset 0 0 9px #2fff6d1f;
    outline: none;
    cursor: text;
    -webkit-text-security: disc;
  }

  .password-pill:hover,
  .password-pill:focus {
    -webkit-text-security: none;
    color: #d5ffe3;
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
  }

  .icon-button:hover {
    color: #d3ffe3;
    box-shadow: 0 0 12px #48ff7d54;
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
