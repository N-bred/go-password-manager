<script module>
  export type CredentialRowReq = {
    id: number;
    domain: string;
    username: string;
    password: string;
  }

  export type CredentialRow = {
    editing: boolean;
    copied: boolean;
  } & CredentialRowReq;

  export type HistoryEntry = {
    id: string;
    value: string;
    createdAt: string;
  };

  export type CredentialHistory = {
    id: number;
    credential_id: string;
    value: string;
    created_at: string;
  }
</script>

<script lang="ts">
  import { onMount } from 'svelte';
  import CredentialTable from './lib/CredentialTable.svelte';
  import HistoryTable from './lib/HistoryTable.svelte';
  import AddCredentialModal from './lib/AddCredentialModal.svelte';
  import Toast from './lib/Toast.svelte';

  let credentials = $state<CredentialRow[]>([]);
  let credentialsHistory = $state<CredentialHistory[]>([]);
  let selectedHistoryForId = $state("");
  let updateSuccess = $state(false);
  let createSuccess = $state(false);
  let createNotify = $state(0);
  let updateNotify = $state(0);
  let showAddModal = $state(false);
  let isCreating = $state(false);
  let toastVisible = $state(false);
  let toastMessage = $state('');
  let toastVariant = $state<'success' | 'error'>('success');

  let toastTimeout: ReturnType<typeof setTimeout> | undefined;
  let domainSearch = $state('');

  const fuzzyMatchDomain = (query: string, domain: string) => {
    const q = query.trim().toLowerCase();
    if (!q) return true;

    const d = domain.toLowerCase();
    let qi = 0;
    for (let i = 0; i < d.length && qi < q.length; i++) {
      if (d[i] === q[qi]) qi++;
    }
    return qi === q.length;
  };

  const filteredCredentials = $derived(
    credentials.filter((c) => fuzzyMatchDomain(domainSearch, c.domain))
  );

  const showToast = (message: string, variant: 'success' | 'error') => {
    if (toastTimeout) clearTimeout(toastTimeout);
    toastMessage = message;
    toastVariant = variant;
    toastVisible = true;
    toastTimeout = setTimeout(() => {
      toastVisible = false;
    }, 3200);
  };

  async function getCredentials() {
    const req = await fetch("http://localhost:3000/credentials")
    const res = await req.json() as CredentialRowReq[];
    credentials = res.map(c => ({...c, editing: false, copied: false}))
  }

  async function getHistoryById(id: number) {
    const req = await fetch(`http://localhost:3000/credential-history/${id}`)
    const res = await req.json() as CredentialHistory[];
    credentialsHistory = res;
  }

  async function updateCredential(c: CredentialRowReq) {
    const req = await fetch(`http://localhost:3000/credential/update`, {
      method: "PUT",
      headers: {"Content-Type": "Application/json"},
      body: JSON.stringify(c)
    });

    const res = await req.json();

    if (res.error) {
      updateSuccess = false;
      updateNotify++;
      console.error(res.error);
      return;
    }

    updateSuccess = true;
    updateNotify++;
  }

  async function createCredential(c: CredentialRowReq) {
    const req = await fetch(`http://localhost:3000/credential/create`, {
      method: "POST",
      headers: {"Content-Type": "Application/json"},
      body: JSON.stringify(c)
    });

    const res = await req.json();

    if (res.error) {
      createSuccess = false;
      createNotify++;
      console.error(res.error);
      return;
    }

    createSuccess = true;
    createNotify++;
  }

  onMount(async () => {
    await getCredentials();
  })

  $effect(() => {
    const notified = updateNotify;
    if (notified === 0) return;

    if (updateSuccess) {
      showToast('Credential updated successfully', 'success');
    } else {
      showToast('Failed to update credential', 'error');
    }
  });

  $effect(() => {
    const notified = createNotify;
    if (notified === 0) return;

    if (createSuccess) {
      showToast('Credential created successfully', 'success');
    } else {
      showToast('Failed to create credential', 'error');
    }
  });

  const handleAddCredential = async (credential: CredentialRowReq) => {
    isCreating = true;
    await createCredential(credential);
    isCreating = false;

    if (!createSuccess) return;

    await getCredentials();
    showAddModal = false;
  };

  const toggleEditOrSave = async (id: number) => {
    const row = credentials.find(c => c.id === id);
    if (!row) return;


    if (row.editing) {
      await updateCredential(row);
    }

    row.editing = !row.editing;
    credentials = [...credentials];
  };

  const copyPassword = async (id: number) => {
    const row = credentials.find(c => c.id === id);
    if (!row) return;

    try {
      await navigator.clipboard.writeText(row.password);
      row.copied = true;
      credentials = [...credentials];
      setTimeout(() => {
        row.copied = false;
        credentials = [...credentials];
      }, 1200);
    } catch (error) {
      console.error('Clipboard copy failed', error);
    }
  };

  const closeHistory = () => {
    selectedHistoryForId = '';
    credentialsHistory = [];
  };

  const toggleHistory = async (id: number) => {
    if (selectedHistoryForId === String(id)) {
      closeHistory();
      return;
    }

    selectedHistoryForId = String(id);
    await getHistoryById(id);
  };
</script>

<main class="page">
  <section class="card">
    <header class="card-header">
      <div>
        <p class="eyebrow">Secure Matrix Vault</p>
        <h1>Password Manager</h1>
      </div>
      <div class="header-actions">
        <input
          class="search-field"
          type="search"
          bind:value={domainSearch}
          placeholder="Search domain…"
          aria-label="Search credentials by domain"
        />
        <button type="button" class="add-button" onclick={() => (showAddModal = true)}>
          + Add credential
        </button>
      </div>
    </header>

    <CredentialTable
      credentials={filteredCredentials}
      activeHistoryId={selectedHistoryForId}
      onToggleEditOrSave={toggleEditOrSave}
      onCopyPassword={copyPassword}
      onToggleHistory={toggleHistory}
    />

    {#if selectedHistoryForId}
      <HistoryTable {selectedHistoryForId} {credentialsHistory} onClose={closeHistory} />
    {/if}
  </section>
</main>

<AddCredentialModal
  open={showAddModal}
  submitting={isCreating}
  onClose={() => (showAddModal = false)}
  onSubmit={handleAddCredential}
/>

<Toast message={toastMessage} variant={toastVariant} visible={toastVisible} />

<style>
  .page {
    min-height: 100dvh;
    display: grid;
    place-items: center;
    padding: 2rem 1rem;
    background:
      radial-gradient(circle at 20% 15%, #00ff7a1a 0%, transparent 38%),
      radial-gradient(circle at 80% 85%, #39ff1412 0%, transparent 35%),
      linear-gradient(145deg, #020503 0%, #040906 48%, #020402 100%);
    color: #9bffbc;
  }

  .card {
    width: min(1080px, 100%);
    border: 1px solid #3cff7540;
    border-radius: 16px;
    background: #030a06d9;
    box-shadow: 0 0 16px #26ff6e2b, 0 0 42px #0d9f3f1f;
    overflow: hidden;
  }

  .card-header {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    align-items: center;
    gap: 1rem;
    padding: 1.2rem 1.4rem;
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

  h1 {
    font-size: clamp(1.08rem, 1.45vw, 1.3rem);
    font-weight: 700;
    color: #caffda;
    text-shadow: 0 0 10px #4bff8a52;
  }

  .header-actions {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    gap: 0.65rem;
  }

  .search-field {
    width: 220px;
    min-width: 0;
    flex-shrink: 1;
    border: 1px solid #59ff8875;
    border-radius: 999px;
    background: #020703;
    color: #beffd6;
    padding: 0.5rem 0.9rem;
    font-size: 0.86rem;
    outline: none;
  }

  .search-field::placeholder {
    color: #6dff9a88;
  }

  .search-field:focus {
    border-color: #8fffb8;
    box-shadow: 0 0 0 2px #49ff7d26;
  }

  .add-button {
    border: 1px solid #5eff8c96;
    border-radius: 999px;
    background: linear-gradient(180deg, #0f2919, #0a1c12);
    color: #adffca;
    font-size: 0.88rem;
    font-weight: 600;
    padding: 0.55rem 0.95rem;
    cursor: pointer;
    box-shadow: inset 0 0 10px #3cff7638;
    white-space: nowrap;
    flex-shrink: 0;
  }

  .add-button:hover {
    color: #d7ffe7;
    border-color: #79ffab;
    box-shadow: inset 0 0 14px #4aff823b, 0 0 14px #4aff8233;
  }

  @media (max-width: 760px) {
    .page {
      padding: 1rem 0.6rem;
    }

    .card-header {
      padding: 1rem;
    }
  }
</style>
