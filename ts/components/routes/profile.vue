<template>
  <div>
    <loginProfile v-if="auth == true" :login="isLogin" @auth="push()"></loginProfile>
    <v-content v-if="auth == false">
      <v-container>
        <v-card class="mx-auto" max-width="500">
          <v-list two-line>
            <v-list-item link v-on:click="changeName()">
              <v-list-item-title>ユーザー名の変更</v-list-item-title>
            </v-list-item>
            <v-divider />

            <v-list-item link v-on:click="changeEmail()">
              <v-list-item-title>Eメールの変更</v-list-item-title>
            </v-list-item>
            <v-divider />

            <v-list-item link v-on:click="changePassword()">
              <v-list-item-title>パスワードの変更</v-list-item-title>
            </v-list-item>
            <v-divider />

            <v-list-item link v-on:click="logout()">
              <v-list-item-title>ログアウト</v-list-item-title>
            </v-list-item>
            <v-list-item v-if="error">
              <v-list-item-title>エラーが発生しました</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-card>
      </v-container>
    </v-content>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import * as firebase from "firebase/app";
import "firebase/auth";
import Login from '../../src/login';
import loginProfile from '../parts/loginForm.vue';

const isLogin = new Login("recertification");
export default Vue.extend({
  methods: {
    logout:async function() {
      const user = firebase.auth().currentUser;
      if (user) {
        try {
          await firebase.auth().signOut();
          this.$emit("check");
          this.$router.push("/");
        } catch (error) {
          this.$data.error = true;
        }
      } else {
        this.$data.error = true;
      }
    },

    changeName:function() {
      this.$data.change = "name";
      this.$data.auth = true;
    },

    changeEmail:function() {
      this.$data.change = "email";
      this.$data.auth = true;
    },

    changePassword:function() {
      this.$data.change = "password";
      this.$data.auth = true;
    },

    push:function() {
      this.$router.push("change/" + this.$data.change);
    }
  },

  created:function(){
    const user = firebase.auth().currentUser;
    if(!user){
      this.$router.push('/');
    }
  },

  components: {
    loginProfile,
  },

  data: () => {
    return {
      error: false,
      auth:false,
      change:"",
      isLogin:isLogin,
    };
  }
});
</script>