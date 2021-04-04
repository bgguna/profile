<template>
  <div class="about">
    <h1>{{ msg }}</h1>
    <br>
    <p>
      My name is Bogdan and I'm a Software Engineer 💻 based in London 🇬🇧.
    </p>
    <p>
      I graduated from The University of Manchester with a Master of Engineering in Computer Science in 2015.
      Since then, I've worked as a backend engineer at a fintech company, designing, building and maintaining enterprise applications and services in Java and Go.
      I like building stuff, experimenting with different technologies, and my goal is to become a confident full-stack engineer.
    </p>
    <p>
      On a different note, I'm an outdoors enthusiast ⛰, explorer 🗺, runner 🏃‍♂️ and coffee snob ☕️. I sometimes fiddle with my camera 📷 and play the guitar 🎸.
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
  max-width: 1200px;
  margin: 0 auto;
}
</style>
