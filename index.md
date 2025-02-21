---
layout: default
title: "Mohamed Lakssir | TheJokers69"
description: "Portfolio de Mohamed Lakssir, étudiant en ingénierie informatique passionné par le Machine Learning et le développement web."
permalink: /
keywords: "Mohamed Lakssir, TheJokers69, portfolio, ingénierie, Machine Learning, développement web"
---

# Bienvenue sur mon Portfolio

Je suis Mohamed Lakssir (@TheJokers69), étudiant en deuxième année d’ingénierie informatique. Découvrez mes compétences et projets ci-dessous :

{% for item in site.data.navigation %}
  <a href="{{ item.link | relative_url }}" class="page-link">{{ item.name }}</a>
{% endfor %}

<!-- ...autres contenus ou sections... -->
