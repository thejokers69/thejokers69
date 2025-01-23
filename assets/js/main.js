// /workspaces/thejokers69/assets/js/main.js
document.addEventListener("DOMContentLoaded", function () {
  console.log("Portfolio loaded successfully!");

  const themeToggle = document.getElementById("theme-toggle");
  const body = document.body;

  // Improved theme toggle with localStorage
  function setTheme(theme) {
    document.body.className = theme;
    localStorage.setItem('theme', theme);
    document.documentElement.style.setProperty('color-scheme', theme);
  }

  // Initialisation
  const currentTheme = localStorage.getItem('theme') || 'dark';
  setTheme(currentTheme);

  // Gestionnaire unique
  document.getElementById('theme-toggle').addEventListener('click', () => {
    const newTheme = document.body.classList.contains('dark') ? 'light' : 'dark';
    setTheme(newTheme);
  });

  // Smooth scroll and intersection observer for animations
  document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
      e.preventDefault();
      document.querySelector(this.getAttribute('href')).scrollIntoView({
        behavior: 'smooth',
        block: 'start'
      });
    });
  });

  // Intersection Observer for fade-in effects
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('visible');
      }
    });
  }, { threshold: 0.1 });

  document.querySelectorAll('.project-card, .skill-item').forEach((el) => {
    observer.observe(el);
  });
});