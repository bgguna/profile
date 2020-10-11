<template>
  <div class="about">
    <h1>{{ msg }}</h1>
    <p>
      My name is Bogdan and I'm a 💻 Software Engineer based in 🇬🇧 London, UK.
    </p>
    <p>
      I graduated from The University of Manchester with a Master of Engineering in Computer Science in 2015.
      Since then, I've worked mostly as a backend engineer in the world of fintech, playing with technologies like Go, Java, and occasionally, JavaScript, Python, Perl and C++.
      My goal is to become a confident full-stack engineer.
    </p>
    <p>
      Apart from being a geek and getting excited about technology, I'm a big ⛰ outdoors enthusiast, ☕️ coffee seeker and 🍺 beer drinker. I sometimes fiddle with my 📷 camera and annoy my neighbours with my 🎸 guitar.
    </p>
  </div>
</template>

<script>
export default {
  name: 'About',
  props: {
      msg: String
  }
}
</script>

<style scoped>

.about {
  padding-left: 10px;
  padding-right: 10px;
  padding-bottom: 20px;
}
</style>
