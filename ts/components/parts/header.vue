<template>
  <div>
    <v-navigation-drawer v-model="drawer" app clipped>
      <v-list dense>
        <v-list-item link v-on:click="push('/')">
          <v-list-item-title>見る</v-list-item-title>
        </v-list-item>

        <v-list-item link v-if="isLogin == true">
          <v-list-item-title>書く</v-list-item-title>
        </v-list-item>

        <v-list-item link v-on:click="push('profile')" v-if="isLogin == true">
          <v-list-item-title>プロフィール</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app clipped-left hide-on-scroll>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Later</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn depressed v-on:click="push('login')" v-if="isLogin == false">
        <span>ログイン</span>
      </v-btn>

      <div v-if="isLogin == true" >
        <span>{{displayName}}でログイン中</span>
      </div>
    </v-app-bar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import * as firebase from "firebase/app";
import "firebase/auth";
export default Vue.extend({
  data: () => ({
    drawer: null,
    isLogin:false,
    displayName:"",
  }),

  created: function() {
    const _this = this;
    firebase.auth().onAuthStateChanged(function(user) {
      if (user) {
        _this.$data.isLogin = true;
        _this.$data.displayName = user.displayName;
      } else {
        _this.$data.isLogin = false;
      }
    });
  },

  methods: {
    push: function(url: string) {
      this.$router.push(url);
    }
  }
});
</script>