<template>
  <div class="about">
    <About msg="hey!"/>
  </div>
</template>

<script>
// @ is an alias to /src
import About from '@/components/About.vue'

export default {
  name: 'AboutMe',
  components: {
    About
  }
}
</script>