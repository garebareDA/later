<template>
  <div>
    <v-navigation-drawer v-model="drawer" app clipped>
      <v-list dense>
        <v-list-item link to="/">
          <v-list-item-title>見る</v-list-item-title>
        </v-list-item>

        <v-list-item link v-if="islogin.is == true"  to="/write">
          <v-list-item-title>書く</v-list-item-title>
        </v-list-item>

        <v-list-item link v-if="islogin.is == true" to="/profile">
          <v-list-item-title>プロフィール</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app clipped-left hide-on-scroll>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Later</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn depressed v-if="islogin.is == false" to="/login">
        <span>ログイン</span>
      </v-btn>

      <div v-if="islogin.is == true" >
        <span>{{islogin.displayName}}でログイン中</span>
      </div>
    </v-app-bar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import * as firebase from "firebase/app";
import "firebase/auth";
import logins from "../../src/isLogin";
import IsLogin from "../../src/isLogin";
export default Vue.extend({
  name:'islogin',
  props:{
    islogin:{
      default:new logins(false, null)
    }
  },

  data: () => ({
    drawer: false,
  }),
});
</script>