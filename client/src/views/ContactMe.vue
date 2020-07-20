<template>
  <div class="contact">
    <Contact msg="Please reach out!"/>
  </div>
</template>

<script>
// @ is an alias to /src
import Contact from '@/components/Contact.vue'

export default {
  name: 'ContactMe',
  components: {
    Contact
  }
}
</script>
