<template>
  
  <v-card
    class="mx-auto justify-center"
    max-width="1020"
  >
  <div v-for="data in datas" :key="data.id">
  
    <video width="100%" controls :src="data.sumber">
        <source  :src="data.sumber" type="video/webm">
        Your browser does not support the video tag.
    </video>
    <v-card-title v-text="data.title"></v-card-title>

    <v-card-subtitle v-text="data.artist"></v-card-subtitle>

    <v-card-actions>
      <v-btn
        v-for="(social, i) in socials"
          :key="i"
          :color="social.color"
          class="white--text ml-2 mt-3"
          fab
          icon
          small
      >
        <v-icon>{{ social.icon }}</v-icon>
      </v-btn>

      <v-spacer></v-spacer>
    </v-card-actions>

    
      <div>
        <v-divider></v-divider>

        <v-card-text v-text="data.deskripsi"></v-card-text>
      </div>
      </div>
  </v-card>
</template>
<script>
import axios from 'axios'
  export default {
    data: () => ({
      show: false,
      datas:[],
      socials: [
        {
          icon: 'mdi-facebook',
          color: 'indigo',
        },
        {
          icon: 'mdi-linkedin',
          color: 'cyan darken-1',
        },
        {
          icon: 'mdi-instagram',
          color: 'red lighten-3',
        },
      ],
    }),
    mounted(){
      this.load()
    },
    methods: {
      load(){
         axios.get('http://localhost:3000/videos')
         .then(response => {
           this.datas = response.data
           console.log(this.datas)
         })
         .catch(error => {
           console.log(error)
         })
      }
    }
  }
</script>