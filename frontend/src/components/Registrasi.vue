<template>
  <!-- Material form register -->
  <v-container>
        <v-row>
  <form>
    <p class="h4 text-center mb-4">Sign up</p>
    <div class="grey-text">
      <v-text-field label="Your name" icon="user" type="text"/>
      <v-text-field label="Your email" icon="envelope" type="email"/>
      <v-text-field label="Confirm your email" icon="exclamation-triangle" type="text"/>
      <v-text-field label="Your password" icon="lock" type="password"/>
    </div>
    <router-link :to="{name:'login'}">
        <v-btn
      color="success"
      class="mr-4"
      @click="validate"
    >
      Daftar
    </v-btn>
    </router-link>
  </form>
  </v-row>
  </v-container>
  <!-- Material form register -->
</template>
<script>
  //import { mdbInput, mdbBtn } from 'mdbvue';
  export default {
    name: 'Basic',
    components: {
      
    }
  }
</script>