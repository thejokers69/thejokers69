// /workspaces/thejokers69/assets/js/main.js
document.addEventListener("DOMContentLoaded", function () {
  console.log("Portfolio loaded successfully!");

  const themeToggle = document.getElementById("theme-toggle");
  const body = document.body;

  // Vérifie si un thème est enregistré dans localStorage
  const savedTheme = localStorage.getItem("theme");
  if (savedTheme) {
    body.classList.add(savedTheme);
  } else {
    body.classList.add("dark"); // Mode sombre par défaut
  }

  // Basculer entre le mode clair et sombre
  themeToggle.addEventListener("click", () => {
    if (body.classList.contains("dark")) {
      body.classList.replace("dark", "light");
      localStorage.setItem("theme", "light");
    } else {
      body.classList.replace("light", "dark");
      localStorage.setItem("theme", "dark");
    }
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

  // Improved theme toggle with localStorage
  function setTheme(theme) {
    document.body.className = theme;
    localStorage.setItem('theme', theme);
    document.documentElement.style.setProperty('color-scheme', theme);
  }

  const currentTheme = localStorage.getItem('theme') || 'dark';
  setTheme(currentTheme);

  document.getElementById('theme-toggle').addEventListener('click', () => {
    const newTheme = document.body.classList.contains('dark') ? 'light' : 'dark';
    setTheme(newTheme);
  });
});