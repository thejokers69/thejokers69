---
layout: default
title: "Mohamed Lakssir | TheJokers69"
description: "Portfolio de Mohamed Lakssir, étudiant en ingénierie informatique passionné par le Machine Learning et le développement web."
permalink: /
keywords: "Mohamed Lakssir, TheJokers69, portfolio, ingénierie, Machine Learning, développement web"
---

[![roadmap.sh](https://roadmap.sh/card/tall/67a29f6ef863343482fe2f28?variant=dark)](https://roadmap.sh)

## Bienvenue sur mon Portfolio

<div class="hero">
  <img
  src="/assets/images/Logo-Of-TheJokerML.jpg"
  alt="Mohamed Lakssir" class="profile-pic">
  <h1>
  Mohamed Lakssir
  <small>
  @TheJokers69
  </small></h1>
  <p>
  Étudiant en ingénierie informatique | Passionné par le ML et le web
  </p>
  <a
  href="/projects"
  class="cta-button">
  Voir mes projets
  </a>
</div>
{% for item in site.data.navigation %}
  <a
  href="{{ item.link | relative_url }}"
  class="page-link">{{ item.name }}</a>
{% endfor %}

<!-- ...autres contenus ou sections... -->
