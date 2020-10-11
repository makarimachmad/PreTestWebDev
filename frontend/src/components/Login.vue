<template>
  <!-- Material form login -->
  <v-container>
        <v-row>
            <div class="mx-auto">
  <form>
    <p class="h4 text-center mb-4">Sign in</p>
    <div class="grey-text">
      <v-text-field label="Your email" icon="envelope" type="email"/>
      <v-text-field label="Your password" icon="lock" type="password"/>
    </div>
    <router-link :to="{name:'akun'}">
        <v-btn
      color="success"
      class="mr-4"
      @click="validate"
    >
      Login
    </v-btn>
    </router-link>
    
  </form>
    <router-link :to="{name:'registrasi'}">
        <v-btn
      color="grey"
      class="mr-4"
      @click="validate"
    >
      Registrasi
    </v-btn>
    </router-link>
    <router-link :to="{name:'admin'}">
        <v-btn
      color="grey"
      class="mr-4"
      @click="validate"
    >
      Admin
    </v-btn>
    </router-link>
  </div>
   </v-row>
  </v-container>
       
  <!-- Material form login -->
  
</template>
<script>
  //import { mdbInput, mdbBtn } from 'mdbvue';
  export default {
    name: 'Basic',
    components: {
      
    },
  }
</script>