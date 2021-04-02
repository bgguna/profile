<template>
  <v-app>
    <div>
      <div class="side-profile">
        <b-button v-b-toggle.profile variant="outline-dark" size="sm">
          <span class="when-closed"><b-icon icon="chevron-right"/></span>
        </b-button>
        <b-sidebar class="profile" id="profile" shadow>
          <div class="sidebar-info">
            <br>
            <div class="profile-img" align="center">
              <img src="@/assets/images/profile_bw.jpg" width="100px" height="100px">
            </div>  
            <div class="profile-description" align="center">
              <p>
                <b>Bogdan Guna</b><br>
                Software Engineer<br>
                London
              </p>
            </div>
            <div class="profile-urls" align="center">
              <b-button href="https://github.com/bgguna" pill variant="outline-dark" size="sm"><b-icon icon="github" variant="secondary"/></b-button>
              <b-button href="https://instagram.com/bgguna" pill variant="outline-dark" size="sm"><b-icon icon="instagram" variant="secondary"/></b-button>
              <b-button href="https://t.me/bgguna" pill variant="outline-dark" size="sm"><b-icon icon="envelope" variant="secondary"/></b-button>
            </div>
          </div>
        </b-sidebar>
      </div>
      <div class="tabs">
        <v-tabs background-color="white" color="black" right>
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

.profile {
  max-width: 200px;
}

.collapse-info {
  padding-left: 25px;
  padding-right: 25px;
}

</style>