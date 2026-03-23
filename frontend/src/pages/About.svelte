<div class="page">
  <h1>О нас</h1>
  <p class="subtitle">Наша история и миссия</p>
  
  <div class="content">
    <section class="about-section">
      <h2>Наша история</h2>
      <p>
        Мы работаем на рынке с 2014 года. За это время мы помогли тысячам клиентов 
        найти качественные решения для их бизнеса. Наша команда состоит из 
        профессионалов, которые любят своё дело.
      </p>
    </section>

    <section class="about-section">
      <h2>Наша миссия</h2>
      <p>
        Создавать инновационные продукты, которые упрощают жизнь наших клиентов 
        и помогают бизнесу расти.
      </p>
    </section>

    <section class="stats">
      <div class="stat">
        <span class="number">10+</span>
        <span class="label">Лет опыта</span>
      </div>
      <div class="stat">
        <span class="number">5000+</span>
        <span class="label">Клиентов</span>
      </div>
      <div class="stat">
        <span class="number">50+</span>
        <span class="label">Продуктов</span>
      </div>
    </section>

    <section class="contact">
      <h2>Контакты</h2>
      <p>📍 Москва, ул. Примерная, д. 123</p>
      <p>📞 +7 (999) 123-45-67</p>
      <p>✉️ info@example.com</p>
    </section>
  </div>
</div>

<style>
  .content {
    text-align: left;
    max-width: 800px;
    margin: 0 auto;
  }
  .about-section {
    margin-bottom: 2rem;
    padding: 1.5rem;
    background: #f9f9f9;
    border-radius: 8px;
  }
  .about-section h2 {
    margin-top: 0;
  }
  .about-section p {
    color: #555;
    line-height: 1.6;
  }
  .stats {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 2rem;
    margin: 3rem 0;
  }
  .stat {
    text-align: center;
    padding: 1.5rem;
    background: #333;
    color: white;
    border-radius: 8px;
  }
  .stat .number {
    display: block;
    font-size: 2.5rem;
    font-weight: bold;
  }
  .stat .label {
    font-size: 0.9rem;
    opacity: 0.8;
  }
  .contact {
    text-align: center;
    padding: 2rem;
    background: #f5f5f5;
    border-radius: 8px;
  }
  .contact h2 {
    margin-top: 0;
  }
  .contact p {
    margin: 0.5rem 0;
    color: #555;
  }
  @media (max-width: 600px) {
    .stats {
      grid-template-columns: 1fr;
    }
  }
</style>