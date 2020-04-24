'use strict';
import Vue from 'vue';
import Router from 'vue-router';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';
import {default as home } from '../components/routes/home.vue';
import {default as login} from '../components/routes/login.vue';
import {default as signUp} from '../components/routes/singUp.vue';
import {default as heads} from '../components/parts/header.vue'

Vue.use(Router);
Vue.use(Vuetify);
const routes = [
  {path:'/', component:home},
  {path:'/login', component:login},
  {path:'/singUp', component:signUp}
];

const vuetify = new Vuetify({});
const router = new Router({
  routes
});

new Vue({
  components:{
    heads
  },
  el:"#app",
  router,
  vuetify,
})