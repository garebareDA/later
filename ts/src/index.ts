'use strict';
import Vue from 'vue';
import Router from 'vue-router';
import { default as home } from '../components/home.vue';

Vue.use(Router);
const routes = [
  {path:'/', component:home},
];

const router = new Router({
  routes
});

new Vue({
  el:"#app",
  router,
})