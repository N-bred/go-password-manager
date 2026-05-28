<script lang="ts">
  let {
    message = '',
    variant = 'success',
    visible = false,
  }: {
    message?: string;
    variant?: 'success' | 'error';
    visible?: boolean;
  } = $props();
</script>

{#if visible}
  <div class="toast-host" aria-live="polite">
    <div class="toast toast-{variant}" role="status">
      <span class="toast-icon" aria-hidden="true">{variant === 'success' ? '✓' : '!'}</span>
      <p>{message}</p>
    </div>
  </div>
{/if}

<style>
  .toast-host {
    position: fixed;
    top: 1.25rem;
    right: 1.25rem;
    z-index: 1000;
    pointer-events: none;
  }

  .toast {
    display: flex;
    align-items: center;
    gap: 0.65rem;
    min-width: 240px;
    max-width: min(420px, calc(100vw - 2rem));
    padding: 0.75rem 1rem;
    border-radius: 10px;
    border: 1px solid;
    background: #030a06f2;
    box-shadow: 0 0 16px #26ff6e2b, 0 0 28px #0d9f3f1f;
    color: #caffda;
    font-size: 0.88rem;
    animation: toast-in 220ms ease-out;
  }

  .toast-success {
    border-color: #5eff8c96;
    box-shadow: inset 0 0 10px #3cff7638, 0 0 14px #4aff8233;
  }

  .toast-error {
    border-color: #ff6b6b80;
    box-shadow: inset 0 0 10px #ff4d4d28, 0 0 14px #ff4d4d22;
    color: #ffd0d0;
  }

  .toast-icon {
    display: grid;
    place-items: center;
    width: 1.35rem;
    height: 1.35rem;
    border-radius: 999px;
    font-size: 0.78rem;
    font-weight: 700;
    flex-shrink: 0;
  }

  .toast-success .toast-icon {
    background: #0f2919;
    color: #adffca;
    border: 1px solid #5eff8c96;
  }

  .toast-error .toast-icon {
    background: #2a1010;
    color: #ffb4b4;
    border: 1px solid #ff6b6b80;
  }

  .toast p {
    margin: 0;
    line-height: 1.35;
  }

  @keyframes toast-in {
    from {
      opacity: 0;
      transform: translateY(-8px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  @media (max-width: 760px) {
    .toast-host {
      left: 1rem;
      right: 1rem;
    }

    .toast {
      width: 100%;
      max-width: none;
    }
  }
</style>
