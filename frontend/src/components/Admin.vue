<template>
    <v-container>
        <v-row>

        
  <div class="mx-auto">
      <br>
      <br>
    <h2>Admin</h2>
  <br><br>
    <h5>Daftar Video</h5>
    <ul v-for="data in datas" :key="data.id">
      <li>
        <span>{{data.title}}</span> &#160;
        <button @click="edit(data)">|| Edit</button> ||  <router-link :to="'/detail/'+data.id">Detail</router-link> || <button @click.prevent="del(data.id)">Delete</button>
      </li>
    </ul>
  </div>
</v-row>
</v-container>
</template>
<script>
/* eslint-disable */ 
import axios from 'axios'
export default {
  data(){
    return{
        datas: [],
        form: {
          id: '',
          title: '',
          artist: '',
          deskripsi: '',
          sifat:true,
        },
        file: '',
        updateSubmit: false
    }
  },
  mounted() {
    this.load()
  },
  methods: {
    load(){
        axios.get('http://localhost:8090')
        .then(res => {
          this.datas = res.data //respon dari rest api dimasukan ke users
          console.log(res.data)
        }).catch ((err) => {
          console.log(err)
        })
    },
    add(file, onUploadProgress){
        this.file = this.$refs.file.files[0];
        console.log(this.form)
        axios.post('http://localhost:8090/tambah', {
          nama: this.form.nama, 
          nim: this.form.nim, 
          email: this.form.email
        },{
          headers:{
            'Content-Type':'application/json'
          }
        })
        .then(res => {
          this.load()
          console.log(res.data)
          this.form.nim=""
          this.form.nama=""
          this.form.email=""
        })
        .catch(error => {
          console.log(error.response)
        })
    },
    del(user){
      axios.delete('http://localhost:8090/hapus?id=' + user)
      .then(res =>{
          this.load()
          let index = this.users.indexOf(form.nama)
          this.users.splice(index,1)
          console.log(res.data)
      }).catch(error => {
          console.log(error) 
      })
    },
    edit(user){ 
        this.updateSubmit = true
        this.form.id = user.id 
        this.form.nama = user.nama
        this.form.nim = user.nim
        this.form.email = user.email 
    },
    update(form){
      return axios.put('http://localhost:8090/update', {
          id: parseInt(this.form.id),
          nama: this.form.nama, 
          nim: this.form.nim, 
          email: this.form.email
        },{
          headers:{
            'Content-Type':'application/json'
          }
        })
        .then(res => {
          this.load()
          console.log(res.data)
          this.form.nim=""
          this.form.nama=""
          this.form.email=""
        }).catch((err) => {
          console.log(err)
      })
    },
  }
}
</script>