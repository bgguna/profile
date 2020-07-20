<template>
  <div class="about">
    <h1>{{ msg }}</h1>
    <p>
      I'm a 💻 Software Engineer based in 🇬🇧 London.
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
</style>
