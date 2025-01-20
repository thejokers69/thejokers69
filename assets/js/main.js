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
});
