<template>
  <v-app>
    <div>
      <div class="side-profile">
        <b-button v-b-toggle.profile variant="outline-dark" size="sm"><b-icon-list></b-icon-list></b-button>
        <b-sidebar id="profile" title="bogdan guna" backdrop shadow>
          <div class="sidebar-info">
            <img src="@/assets/images/profile_bw.jpg" width="135px" height="200px">
            <p>
              Hello there!
            </p>
          </div>
        </b-sidebar>
      </div>
      <div class="tabs">
        <v-tabs
          background-color="white"
          color="black"
          right
        >
          <v-tab>about me</v-tab>
          <v-tab>contact me</v-tab>

          <v-tab-item :key=1>
            <About class="all" msg="hey!"/>
          </v-tab-item>

          <v-tab-item :key=2>
            <Contact msg="let's chat!"/>
          </v-tab-item>
        </v-tabs>
      </div>
    </div>
  </v-app>
</template>

<script>
import Vue from 'vue'
import axios from 'axios';
import VueAxios from 'vue-axios';

import About from '@/components/About.vue'
import Contact from '@/components/Contact.vue'
import { BootstrapVue, BootstrapVueIcons } from 'bootstrap-vue'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

// Install BootstrapVue
Vue.use(BootstrapVue)
Vue.use(BootstrapVueIcons)

// Use Axios
Vue.use(VueAxios, axios);

export default {
  name: 'App',
  components: {
    About,
    Contact
  }
}
</script>

<style scoped>
.all {
  min-width: 300px;
  max-width: 600px;
}

.sidebar-info {
  padding-left: 25px;
  padding-right: 25px;
}
</style>