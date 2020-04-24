'use strict';
import * as firebase from 'firebase/app';
import Vue from 'vue';
import Router from 'vue-router';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';
import {default as home } from '../components/routes/home.vue';
import {default as login} from '../components/routes/login.vue';
import {default as signUp} from '../components/routes/singUp.vue';
import {default as profile} from '../components/routes/profile.vue';
import {default as heads} from '../components/parts/header.vue';


const firebaseConfig = {
  apiKey: "AIzaSyDm9xIFZ1Xul9Jbg19ocK6XzeFvU9gje04",
  authDomain: "later-d4187.firebaseapp.com",
  databaseURL: "https://later-d4187.firebaseio.com",
  projectId: "later-d4187",
  storageBucket: "later-d4187.appspot.com",
  messagingSenderId: "386886264700",
  appId: "1:386886264700:web:3f228e273448f4551aa6f9",
  measurementId: "G-CFNR8RNLDZ"
};
firebase.initializeApp(firebaseConfig);

Vue.use(Router);
Vue.use(Vuetify);
const routes = [
  {path:'/', component:home},
  {path:'/login', component:login},
  {path:'/singUp', component:signUp},
  {path:'/profile', component:profile}
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
});