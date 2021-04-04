<template>
  <div class="about">
    <h1>{{ msg }}</h1>
    <br />
    <p>My name is Bogdan and I'm a Software Engineer 💻 based in London 🇬🇧.</p>
    <p>
      I graduated from The University of Manchester with a Master of Engineering
      in Computer Science in 2015. Since then, I've worked as a backend engineer
      in the fintech industry, designing, building and maintaining enterprise
      applications and services in Java and Go.
    </p>
    <p>
      I'm also a great outdoors enthusiast ⛰, cyclist 🚲, runner 🏃‍♂️ and
      somewhat, a coffee snob ☕️. On most days, you can also find me fiddling
      with my camera 📷 or my guitar 🎸.
    </p>
  </div>
</template>

<script>
export default {
  name: "About",
  props: {
    msg: String,
  },
};
</script>

<style scoped>
.about {
  padding-left: 10px;
  padding-right: 10px;
  padding-bottom: 20px;
  max-width: 1200px;
  margin: 20px auto;
}
</style>
