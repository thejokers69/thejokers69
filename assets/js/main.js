// /workspaces/thejokers69/assets/js/main.js
document.addEventListener("DOMContentLoaded", function () {
  console.log("Portfolio loaded successfully!");

  // Sélectionne le bouton et le body
  const themeToggle = document.getElementById('theme-toggle');
  const body = document.body;

  // Vérifie si un thème est défini dans localStorage
  const savedTheme = localStorage.getItem('theme');
  if (savedTheme) {
    body.classList.add(savedTheme);
  } else {
    body.classList.add('dark'); // Mode sombre par défaut
  }

  // Gestionnaire d'événement pour basculer le thème
  themeToggle.addEventListener('click', () => {
    if (body.classList.contains('dark')) {
      body.classList.replace('dark', 'light');
      localStorage.setItem('theme', 'light');
    } else {
      body.classList.replace('light', 'dark');
      localStorage.setItem('theme', 'dark');
    }
  });
});