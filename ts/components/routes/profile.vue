<template>
  <div>
    <v-content>
      <v-container>
        <v-card class="mx-auto" max-width="500">
          <v-list dense two-line>
            <v-list-item link>
              <v-list-item-title>ユーザー名の変更</v-list-item-title>
            </v-list-item>
            <v-divider />

            <v-list-item link>
              <v-list-item-title>Eメールの変更</v-list-item-title>
            </v-list-item>
            <v-divider />

            <v-list-item link>
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
export default Vue.extend({
  methods: {
    logout: function() {
      const _this = this;
      firebase.auth().onAuthStateChanged(user => {
        firebase
          .auth()
          .signOut()
          .then(() => {
            _this.$router.push('/');
          })
          .catch(error => {
            _this.$data.error = true;
          });
      });
    },
  },

  data:() => {
    return{
      error:false,
    }
  }
});
</script>