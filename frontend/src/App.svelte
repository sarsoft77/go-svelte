<script>
  import Home from './pages/Home.svelte';
  import Products from './pages/Products.svelte';
  import About from './pages/About.svelte';
  import Admin from './pages/Admin.svelte';

  let currentPage = 'home';

  const pages = {
    home: Home,
    products: Products,
    about: About,
    admin: Admin
  };
</script>

<nav>
  <div class="nav-brand">Моя Компания</div>
  <div class="nav-links">
    <button class:active={currentPage === 'home'} on:click={() => currentPage = 'home'}>Главная</button>
    <button class:active={currentPage === 'products'} on:click={() => currentPage = 'products'}>Продукция</button>
    <button class:active={currentPage === 'about'} on:click={() => currentPage = 'about'}>О нас</button>
    <button class:active={currentPage === 'admin'} on:click={() => currentPage = 'admin'}>Админ</button>
  </div>
</nav>

<main>
  <svelte:component this={pages[currentPage]} />
</main>