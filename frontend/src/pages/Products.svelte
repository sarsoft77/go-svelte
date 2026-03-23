<script>
  import { onMount } from 'svelte';
  
  let products = [];
  let loading = true;
  let error = '';
  let selectedProduct = null;

  onMount(async () => {
    try {
      const res = await fetch('/api/products?pageSize=100');
      const data = await res.json();
      if (data.error) throw new Error(data.error);
      products = data.products || [];
    } catch (e) {
      error = 'Ошибка загрузки продуктов: ' + e.message;
    } finally {
      loading = false;
    }
  });

  function showDetails(product) {
    selectedProduct = product;
  }

  function closeDetails() {
    selectedProduct = null;
  }
</script>

<div class="page">
  <h1>Наша продукция</h1>
  <p class="subtitle">Лучшие решения для вашего бизнеса</p>
  
  {#if loading}
    <p class="loading">Загрузка...</p>
  {:else if error}
    <p class="error">{error}</p>
  {:else}
    <div class="products">
      {#each products as product}
        <div class="product-card">
          <div class="product-icon">📦</div>
          <h3>{product.name}</h3>
          <p class="description">{product.description}</p>
          <p class="price">{product.price}</p>
          <button class="btn-primary" on:click={() => showDetails(product)}>Подробнее</button>
        </div>
      {/each}
    </div>
  {/if}
</div>

{#if selectedProduct}
  <div class="modal-overlay" on:click={closeDetails} on:keydown={(e) => e.key === 'Escape' && closeDetails()} role="button" tabindex="0">
    <div class="modal" role="dialog" aria-modal="true">
      <h2>{selectedProduct.name}</h2>
      <p class="modal-price">{selectedProduct.price}</p>
      <p class="modal-description">{selectedProduct.description}</p>
      <p class="modal-id">ID товара: {selectedProduct.id}</p>
      <button class="btn-primary" on:click={closeDetails}>Закрыть</button>
    </div>
  </div>
{/if}

<style>
  .products {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 2rem;
  }
  .product-card {
    background: white;
    border: 1px solid #ddd;
    border-radius: 8px;
    padding: 1.5rem;
    text-align: center;
    transition: box-shadow 0.2s;
  }
  .product-card:hover {
    box-shadow: 0 4px 12px rgba(0,0,0,0.1);
  }
  .product-icon {
    font-size: 3rem;
    margin-bottom: 1rem;
  }
  .product-card h3 {
    margin: 0 0 0.5rem;
  }
  .description {
    color: #666;
    margin-bottom: 1rem;
  }
  .price {
    font-size: 1.5rem;
    font-weight: bold;
    color: #333;
    margin-bottom: 1rem;
  }
  .modal-price {
    font-size: 2rem;
    font-weight: bold;
    color: #333;
    margin: 1rem 0;
  }
  .modal-id {
    color: #999;
    font-size: 0.9rem;
    margin-bottom: 1.5rem;
  }
  .modal {
    pointer-events: auto;
  }
</style>