<template>
  <div class="contact">
    <h1>{{ msg }}</h1>
    <p>
        Complete this form.
    </p>
  </div>
</template>

<script>
export default {
  name: 'Contact',
  props: {
    msg: String
  }
}
</script>

<style scoped>
</style>
