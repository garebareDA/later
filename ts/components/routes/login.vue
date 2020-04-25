<template>
  <div>
    <login :login="isLogin" @login="isLogins"></login>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { default as heads } from '../parts/header.vue';
import {default as login} from '../parts/loginForm.vue';
import Login from '../../src/login'

const isLogin = new Login(true);
export default Vue.extend({
  components: {
    heads,
    login,
  },

  methods:{
    isLogins:function() {
      this.$emit('login');
    }
  },

  data:() => {
    return{
      isLogin:isLogin
    }
  }
});
</script>