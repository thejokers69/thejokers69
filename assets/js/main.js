/**
 * @ai-context Main JavaScript file for theme switching and animations
 * @project thejokers69
 * @description Handles theme toggling, smooth scrolling, and element animations
 * @author Mohamed Lakssir (@TheJokers69)
 */
document.addEventListener("DOMContentLoaded", function () {
  console.log("Portfolio loaded successfully!");

  const themeToggle = document.getElementById("theme-toggle");
  const body = document.body;

  /**
   * Sets the theme for the website and stores preference in localStorage
   * @param {string} theme - 'dark' or 'light' theme value
   */
  function setTheme(theme) {
    try {
      body.className = theme;
      localStorage.setItem("theme", theme);
      document.documentElement.style.setProperty("color-scheme", theme);
      
      // Update ARIA attributes for accessibility
      themeToggle.setAttribute("aria-pressed", theme === "dark" ? "true" : "false");
      themeToggle.setAttribute("aria-label", theme === "dark" ? "Switch to light theme" : "Switch to dark theme");
    } catch (e) {
      body.className = theme;
      console.warn(
        "LocalStorage non-disponible, utilisation de la session uniquement."
      );
    }
  }

  // Load the saved theme or default to dark
  const currentTheme = localStorage.getItem("theme") || "dark";
  setTheme(currentTheme);

  // Theme toggle functionality with enhanced accessibility
  themeToggle.addEventListener("click", () => {
    const newTheme = body.classList.contains("dark") ? "light" : "dark";
    setTheme(newTheme);
  });

  // Smooth scrolling for anchor links
  document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
    anchor.addEventListener("click", function (e) {
      e.preventDefault();
      const targetElement = document.querySelector(this.getAttribute("href"));
      if (targetElement) {
        targetElement.scrollIntoView({
          behavior: "smooth",
          block: "start",
        });
      }
    });
  });

  // Element reveal animations using Intersection Observer
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add("visible");
          observer.unobserve(entry.target);
        }
      });
    },
    { threshold: 0.05, rootMargin: "0px 0px -50px 0px" }
  );

  // Apply animations to project cards and skill items
  document.querySelectorAll(".project-card, .skill-item").forEach((el) => {
    observer.observe(el);
  });
});
