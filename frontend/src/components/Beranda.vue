<template>
    <v-container>
      <v-row dense>
        <v-col
          v-for="data in datas"
          :key="data.i"
          cols="12"
        >
          <v-card
            color="#1F7087"
            dark
          >
            <div class="d-flex flex-no-wrap justify-space-between">
              <div>
                <v-card-title
                  class="headline"
                  v-text="data.title"
                ></v-card-title>

                <v-card-subtitle v-text="data.artist"></v-card-subtitle>

                  <router-link :to="{name:'detail'}">
                    <v-btn  
                        class="ml-2 mt-3"
                        fab
                        icon
                        height="40px"
                        right
                        width="40px"
                        @click="sementara(data)"
                    >
                        <v-icon>mdi-play</v-icon>
                    </v-btn>
                  </router-link>
                  
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
                
              </div>

              <v-avatar
                class="ma-3"
                size="125"
                tile
              >
                <v-img :src="data.src"></v-img>
              </v-avatar>
            </div>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
</template>
<script>
import axios from 'axios'
  export default {
    data: () => ({
      datas:[],
      tampung:{
        id:1,
        src:'',
        title: '',
        artist: '',
        deskripsi: '',
        sifat: false,
      },
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
            axios.get('http://localhost:8090')
            .then(response => {
              this.datas = response.data
              console.log(this.datas)
            }).catch( err => {
              console.log(err.response)
            })
        },
        sementara(data){
          //this.tampung.id=data.id,
          this.tampung.src=data.src,
          this.tampung.title=data.title,
          this.tampung.artist=data.artist,
          this.tampung.sifat=data.sifat,
          this.tampung.deskripsi=data.deskripsi

          console.log(this.tampung)
          return axios.put('http://localhost:3000/videos/'+this.tampung.id,
          {
            id: this.tampung.id,
            sumber: this.tampung.src,
            title: this.tampung.title,
            artist: this.tampung.artist,
            sifat: this.tampung.sifat,
            deskripsi: this.tampung.deskripsi
          })
          .then(response => {
            console.log(response.data)
          })
          .catch((error) => {
            console.log(error)
          })
        }
    },
  }
</script>