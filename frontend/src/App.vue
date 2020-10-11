<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png"
          transition="scale-transition"
          width="40"
        />
        Kamu Tube
      </div>

      <v-spacer></v-spacer>
      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>
      <router-link :to="{name:'login'}">
        <v-btn icon>
          <v-icon>mdi-account</v-icon>
        </v-btn>
      </router-link>
      
      <!-- <v-btn
        href="https://github.com/vuetifyjs/vuetify/releases/latest"
        target="_blank"
        text
      >
        <span class="mr-2">Latest Release</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn> -->
    </v-app-bar>

    <v-main>
      <!-- <HelloWorld/> -->
      <br>
      <router-view></router-view><br><br><br>

      <Footer/>
    </v-main>
  </v-app>
</template>

<script>
// import HelloWorld from './components/HelloWorld';
import Footer from './components/Footer'

export default {
  name: 'App',

  components: {
    // HelloWorld,
    Footer
  },

  data: () => ({
    //
  }),
};
</script>
