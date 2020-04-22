'use strict';
import Vue from 'vue';
import Router from 'vue-router';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';
import {default as home } from '../components/home.vue';

Vue.use(Router);
Vue.use(Vuetify);
const routes = [
  {path:'/', component:home},
];

const vuetify = new Vuetify({});
const router = new Router({
  routes
});

new Vue({
  el:"#app",
  router,
  vuetify,
})