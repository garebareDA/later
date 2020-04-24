<template>
  <div>
    <v-navigation-drawer v-model="drawer" app clipped>
      <v-list dense>
        <v-list-item link v-on:click="push('/')">
          <v-list-item-title>見る</v-list-item-title>
        </v-list-item>

        <v-list-item link>
          <v-list-item-title>書く</v-list-item-title>
        </v-list-item>

        <v-list-item link>
          <v-list-item-title>プロフィール</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app clipped-left hide-on-scroll>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Later</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn depressed v-on:click="push('login')">
        <span>ログイン</span>
      </v-btn>
    </v-app-bar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  data: () => ({
    drawer: null
  }),

  computed:() => {
    
  },

  methods:{
    push:function(url:string) {
      this.$router.push(url);
    }
  }
});
</script>