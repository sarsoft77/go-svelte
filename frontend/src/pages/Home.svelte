<script>
  let apiMessage = '';

  async function fetchMessage() {
    const res = await fetch('/api/hello');
    const data = await res.json();
    apiMessage = data.message;
  }
</script>

<div class="page">
  <h1>Добро пожаловать!</h1>
  <p class="subtitle">Мы создаём качественные продукты для вас</p>
  
  <div class="features">
    <div class="feature">
      <h3>Качество</h3>
      <p>Мы гарантируем высокое качество нашей продукции</p>
    </div>
    <div class="feature">
      <h3>Надёжность</h3>
      <p>Более 10 лет на рынке</p>
    </div>
    <div class="feature">
      <h3>Поддержка</h3>
      <p>24/7 поддержка клиентов</p>
    </div>
  </div>

  <button class="btn-primary" on:click={fetchMessage}>
    Проверить API
  </button>
  
  {#if apiMessage}
    <p class="response">{apiMessage}</p>
  {/if}
</div>

<style>
  .features {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 2rem;
    margin: 3rem 0;
  }
  .feature {
    padding: 1.5rem;
    background: #f5f5f5;
    border-radius: 8px;
  }
  .feature h3 {
    margin-top: 0;
  }
  .feature p {
    color: #666;
  }
  .response {
    margin-top: 1rem;
    color: green;
    font-weight: bold;
  }
</style>