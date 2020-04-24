<template>
  <div>
    <heads></heads>
    <login v-bind:login="isLogin"></login>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { default as heads } from '../parts/header.vue';
import {default as login} from '../parts/loginForm.vue';
import Login from '../../src/login';

const isLogin = new Login(false);
export default Vue.extend({
    components: {
    heads,
    login,
  },

  data:() => {
    return{
      isLogin:isLogin
    }
  }
})
</script>