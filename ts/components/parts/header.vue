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

        <v-list-item link v-on:click="push('profile')" v-if="isLogin.is == true">
          <v-list-item-title>プロフィール</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app clipped-left hide-on-scroll>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Later</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn depressed v-on:click="push('login')" v-if="isLogin.is == false">
        <span>ログイン</span>
      </v-btn>

      <div v-if="isLogin.is == true" >
        <span>{{isLogin.displayName}}でログイン中</span>
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
  }),

  props:['isLogin'],

  methods: {
    push: function(url: string) {
      this.$router.push(url);
    }
  }
});
</script>