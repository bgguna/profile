<template>
  <v-app>
    <div>
        <v-tabs
      background-color="white"
      color="black"
      right
    >
      <v-tab>about me</v-tab>
      <v-tab>contact me</v-tab>

      <v-tab-item :key=1>
        <About msg="hey!"/>
      </v-tab-item>

      <v-tab-item :key=2>
        <Contact msg="let's chat!"/>
      </v-tab-item>
    </v-tabs>
    </div>

  </v-app>
</template>

<script>
import About from '@/components/About.vue'
import Contact from '@/components/Contact.vue'

export default {
  name: 'App',
  components: {
    About,
    Contact
  }
}
</script>
