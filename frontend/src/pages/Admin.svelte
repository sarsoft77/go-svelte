<script>
  import { onMount } from 'svelte';

  let products = [];
  let loading = true;
  let error = '';
  let editing = null;
  let form = { name: '', description: '', price: '' };
  let showModal = false;
  let sortKey = 'id';
  let sortDir = 'asc';
  let search = '';
  let currentPage = 1;
  let pageSize = 10;
  let total = 0;
  let totalPages = 0;
  let searchTimeout;

  onMount(() => {
    loadProducts();
  });

  async function loadProducts() {
    loading = true;
    error = '';
    try {
      const params = new URLSearchParams({
        page: currentPage,
        pageSize: pageSize,
        sort: sortKey,
        sortDir: sortDir,
        search: search
      });
      const res = await fetch(`/api/products?${params}`);
      if (!res.ok) throw new Error('Ошибка сервера: ' + res.status);
      const data = await res.json();
      if (data.error) throw new Error(data.error);
      products = data.products || [];
      total = data.total || 0;
      totalPages = data.totalPages || 0;
    } catch (e) {
      error = 'Ошибка загрузки: ' + e.message;
      products = [];
    } finally {
      loading = false;
    }
  }

  let previousSearch = '';
  
  function handleSearch() {
    if (search === '' && previousSearch !== '') {
      previousSearch = '';
      currentPage = 1;
      loadProducts();
    } else if (search !== previousSearch) {
      previousSearch = search;
      clearTimeout(searchTimeout);
      searchTimeout = setTimeout(() => {
        currentPage = 1;
        loadProducts();
      }, 300);
    }
  }

  function sortBy(key) {
    if (sortKey === key) {
      sortDir = sortDir === 'asc' ? 'desc' : 'asc';
    } else {
      sortKey = key;
      sortDir = 'asc';
    }
    currentPage = 1;
    loadProducts();
  }

  function setPage(p) {
    if (p >= 1 && p <= totalPages) {
      currentPage = p;
      loadProducts();
    }
  }

  function changePageSize() {
    currentPage = 1;
    loadProducts();
  }

  function getSortIcon(key) {
    if (sortKey !== key) return '';
    return sortDir === 'asc' ? ' ▲' : ' ▼';
  }

  function openAdd() {
    editing = null;
    form = { name: '', description: '', price: '' };
    showModal = true;
  }

  function openEdit(p) {
    editing = p;
    form = { name: p.name, description: p.description, price: p.price };
    showModal = true;
  }

  function closeModal(e) {
    if (e && e.target !== e.currentTarget) return;
    showModal = false;
    editing = null;
  }

  async function saveProduct() {
    try {
      let res;
      if (editing) {
        res = await fetch(`/api/products/${editing.id}`, {
          method: 'PUT',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(form)
        });
      } else {
        res = await fetch('/api/products', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(form)
        });
      }
      const data = await res.json();
      if (data.error) throw new Error(data.error);
      closeModal();
      await loadProducts();
    } catch (e) {
      error = e.message;
    }
  }

  async function deleteProduct(id) {
    if (!confirm('Удалить продукт?')) return;
    try {
      const res = await fetch(`/api/products/${id}`, { method: 'DELETE' });
      const data = await res.json();
      if (data.error) throw new Error(data.error);
      await loadProducts();
    } catch (e) {
      error = e.message;
    }
  }

  function handleKeydown(e) {
    if (e.key === 'Escape') closeModal();
  }
</script>

<svelte:window on:keydown={handleKeydown} />

<div class="page">
  <h1>Администрирование</h1>
  
  {#if error}
    <p class="error">{error}</p>
  {/if}

  <div class="header-row">
    <div class="search-box">
      <input bind:value={search} placeholder="Поиск..." on:input={handleSearch} />
    </div>
    <div class="header-right">
      <span class="subtitle">Всего: {total}</span>
      <select bind:value={pageSize} on:change={changePageSize}>
        <option value={5}>5 на странице</option>
        <option value={10}>10 на странице</option>
        <option value={20}>20 на странице</option>
        <option value={50}>50 на странице</option>
      </select>
      <button class="btn-success" on:click={openAdd}>+ Добавить</button>
    </div>
  </div>

  {#if loading}
    <p class="loading">Загрузка...</p>
  {:else}
    <table>
      <thead>
        <tr>
          <th class="sortable" class:active={sortKey === 'id'} on:click={() => sortBy('id')}>
            ID{getSortIcon('id')}
          </th>
          <th class="sortable" class:active={sortKey === 'name'} on:click={() => sortBy('name')}>
            Название{getSortIcon('name')}
          </th>
          <th class="sortable" class:active={sortKey === 'description'} on:click={() => sortBy('description')}>
            Описание{getSortIcon('description')}
          </th>
          <th class="sortable" class:active={sortKey === 'price'} on:click={() => sortBy('price')}>
            Цена{getSortIcon('price')}
          </th>
          <th>Действия</th>
        </tr>
      </thead>
      <tbody>
        {#each products as p}
          <tr>
            <td>{p.id}</td>
            <td>{p.name}</td>
            <td>{p.description}</td>
            <td>{p.price}</td>
            <td class="actions">
              <button class="btn-edit" on:click={() => openEdit(p)}>Изменить</button>
              <button class="btn-danger" on:click={() => deleteProduct(p.id)}>Удалить</button>
            </td>
          </tr>
        {/each}
        {#if products.length === 0}
          <tr>
            <td colspan="5" class="empty">Нет данных</td>
          </tr>
        {/if}
      </tbody>
    </table>

    {#if totalPages > 1}
      <div class="pagination">
        <button disabled={currentPage === 1} on:click={() => setPage(1)}>&laquo;</button>
        <button disabled={currentPage === 1} on:click={() => setPage(currentPage - 1)}>&lsaquo;</button>
        <span class="page-info">Страница {currentPage} из {totalPages}</span>
        <button disabled={currentPage === totalPages} on:click={() => setPage(currentPage + 1)}>&rsaquo;</button>
        <button disabled={currentPage === totalPages} on:click={() => setPage(totalPages)}>&raquo;</button>
      </div>
    {/if}
  {/if}
</div>

{#if showModal}
  <!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
  <div class="modal-overlay" on:click={closeModal} on:keydown={(e) => e.key === 'Escape' && closeModal()} role="presentation">
    <div class="modal" role="dialog" aria-modal="true">
      <h2>{editing ? 'Редактировать продукт' : 'Новый продукт'}</h2>
      <div class="form-group">
        <label for="name">Название</label>
        <input id="name" class="input-field" bind:value={form.name} placeholder="Введите название" />
      </div>
      <div class="form-group">
        <label for="description">Описание</label>
        <textarea id="description" class="input-field" bind:value={form.description} placeholder="Введите описание" rows="3"></textarea>
      </div>
      <div class="form-group">
        <label for="price">Цена</label>
        <input id="price" class="input-field" bind:value={form.price} placeholder="Введите цену" />
      </div>
      <div class="modal-actions">
        <button class="btn-success" on:click={saveProduct}>Сохранить</button>
        <button class="btn-secondary" on:click={closeModal}>Отмена</button>
      </div>
    </div>
  </div>
{/if}

<style>
  .header-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1.5rem;
    flex-wrap: wrap;
    gap: 1rem;
  }
  .search-box input {
    padding: 10px 15px;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 14px;
    width: 250px;
  }
  .header-right {
    display: flex;
    align-items: center;
    gap: 1rem;
  }
  .header-right select {
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 4px;
  }
  .actions {
    display: flex;
    gap: 8px;
  }
  .btn-edit {
    background: #007bff;
    color: white;
    border: none;
    padding: 6px 12px;
    border-radius: 4px;
    cursor: pointer;
  }
  .btn-edit:hover {
    background: #0056b3;
  }
  .empty {
    text-align: center;
    color: #999;
  }
  .modal {
    width: 450px;
  }
  .modal h2 {
    margin-top: 0;
    margin-bottom: 1.5rem;
  }
  .form-group {
    margin-bottom: 1rem;
    text-align: left;
  }
  .form-group label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 500;
    color: #333;
  }
  .form-group textarea {
    resize: vertical;
    font-family: inherit;
  }
  .modal-actions {
    display: flex;
    gap: 10px;
    justify-content: flex-end;
    margin-top: 1.5rem;
  }
</style>