// /workspaces/thejokers69/assets/js/main.js

document.addEventListener("DOMContentLoaded", function () {
  console.log("Portfolio loaded successfully!");

  const themeToggle = document.getElementById("theme-toggle");
  const body = document.body;

  function setTheme(theme) {
    try {
      body.className = theme;
      localStorage.setItem('theme', theme);
      document.documentElement.style.setProperty('color-scheme', theme);
    } catch (e) {
      body.className = theme;
      console.warn('LocalStorage non-disponible, utilisation de la session uniquement.');
    }
  }

  const currentTheme = localStorage.getItem('theme') || 'dark';
  setTheme(currentTheme);

  themeToggle.addEventListener('click', () => {
    const newTheme = body.classList.contains('dark') ? 'light' : 'dark';
    setTheme(newTheme);
  });

  document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
      e.preventDefault();
      document.querySelector(this.getAttribute('href')).scrollIntoView({
        behavior: 'smooth',
        block: 'start'
      });
    });
  });

  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('visible');
        observer.unobserve(entry.target);
      }
    });
  }, { threshold: 0.05 });

  document.querySelectorAll('.project-card, .skill-item').forEach(el => {
    observer.observe(el);
  });

  
  function toggleSidebar() {
    const sidebar = document.querySelector('.sidebar');
    sidebar.classList.toggle('open');
  }

  
  const sidebarToggle = document.querySelector('.sidebar-toggle');
  sidebarToggle.addEventListener('click', toggleSidebar);
});
