<template>
  <div>
    <login :login="isLogin" @login="isLogins"></login>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { default as heads } from '../parts/header.vue';
import {default as login} from '../parts/loginForm.vue';
import Login from '../../src/login';

const isLogin = new Login("login");
export default Vue.extend({
  components: {
    heads,
    login,
  },

  methods:{
    isLogins:function() {
      this.$emit('check');
    }
  },

  data:() => {
    return{
      isLogin:isLogin
    }
  }
});
</script>