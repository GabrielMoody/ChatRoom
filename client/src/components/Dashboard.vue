<template>
  <layout-div>
     <div class="row justify-content-md-center">
       <div class="col-12">
           <nav class="navbar navbar-expand-lg navbar-light bg-light">
               <div class="container-fluid">
                   <a class="navbar-brand" href="#">Dashboard</a>
                   <div class="d-flex">
                       <ul class="navbar-nav">
                           <li class="nav-item">
                               <a @click="logoutAction()" class="nav-link " aria-current="page" href="#">Logout</a>
                           </li>
                       </ul>
                   </div>
               </div>
           </nav>
           <h2 class="text-center mt-5">Welcome, {{user}}!</h2  >
       </div>

       <form @submit.prevent="findRoom">
        <label for="room">Find Room : </label>
        <input v-model="room" class="form-control" type="text" name="room" id="room">
        <button class="btn btn-primary" type="submit">Find</button>
       </form>

       <div class="card" v-for="room in rooms" :key="room.ID">
        <p>{{ room.Name }}</p>
        <p>Authors: {{ room.CreatedBy }}</p>
        <button class="btn btn-primary">Join</button>
       </div>
     </div>
  </layout-div>
</template>
 
<script>
import axios from 'axios';
import LayoutDiv from './LayoutDiv.vue';
 
const url = 'http://localhost:8000/api/v1';

export default {
 name: 'DashboardPage',
 components: {
   LayoutDiv,
 },
 data() {
   return {
     user: {},
     rooms: [],
     room: ''
   };
 },
 created() {
   this.getUser();
   if(localStorage.getItem('token') == "" || localStorage.getItem('token') == null){
     this.$router.push('/')
   }else {
     this.getUser();
   }

 },
 methods: {
   getUser() {
       axios.get(`${url}/user`, { headers:{Authorization: 'Bearer ' + localStorage.getItem('token')}})
       .then((r) => {
          this.user = r.data.data.Name;
          return r
       })
       .catch((e) => {
          return e
       });
   },

   logoutAction () {
     axios.post('/api/logout',{}, { headers:{Authorization: 'Bearer ' + localStorage.getItem('token')}})
     .then((r) => {
         localStorage.setItem('token', "")
         this.$router.push('/')
         return r
     })
     .catch((e) => {
       return e
     });
   },

   findRoom() {
    this.rooms = [];
    axios.get(`/api/v1/rooms?room=${this.room}`, { 
      headers: {
        Authorization: 'Bearer ' + localStorage.getItem('token')
      }
    }).then((r) => {
      const data = r.data.rooms;

      data.forEach(d => {
        this.rooms.push(d);
      })

      return this.rooms;
    }).catch(e => e)
   }

 },
};
</script>