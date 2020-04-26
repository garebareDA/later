<template>
  <div>
    <v-content>
      <v-container>
        <v-row align="center" justify="center">
          <v-col cols="12" sm="8" md="4">
            <v-card class="elevation-12">
              <v-toolbar>
                <v-spacer />
                <v-toolbar-title>変更</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <v-form v-if="thisPage == 'name'">
                  <v-text-field label="新しいユーザー名" type="text" v-model="text" />
                </v-form>

                <v-form v-if="thisPage == 'email'">
                  <v-text-field label="新しいEメール" type="text" v-model="text" />
                </v-form>

                <v-form v-if="thisPage == 'password'">
                  <v-text-field label="新しいパスワード" type="text" v-model="text" />
                </v-form>

                <v-card-actions v-show="errors.isError == true">
                  <v-spacer />
                  <div class="red--text">{{errors.message}}</div>
                </v-card-actions>

                <v-card-actions>
                  <v-spacer />
                  <v-btn outlined v-on:click="changeDone()">変更</v-btn>
                </v-card-actions>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import * as firebase from "firebase/app";
import "firebase/auth";
export default Vue.extend({
  created: function() {
    this.thisPage = this.$route.params.state;
    const user = firebase.auth().currentUser;
    if (!user) {
      this.$router.push("/");
    }
  },

  methods: {
    changeDone: async function() {
      const state = this.$route.params.state;
      const user = firebase.auth().currentUser;
      const text = this.$data.text;

      if (text === "") {
        this.error("入力してください");
        return;
      }

      if (!user) {
        this.$router.push("/");
        return;
      }

      if (state === "name") {
        try {
          await user.updateProfile({
            displayName: text
          });
          this.$emit("check");
          this.$router.push("/");
        } catch (error) {
          this.error("エラーが発生しました");
        }
      } else if (state === "email") {
        try {
          await user.updateEmail(text);
          this.$router.push("/");
        } catch (error) {
          this.error("エラーが発生しました");
        }
      } else if (state === "password") {
        try {
          await user.updatePassword(text);
          this.$router.push("/");
        } catch (error) {
          this.error("エラーが発生しました");
        }
      } else {
        this.error("エラーが発生しました");
      }
    },

    error: function(message: string) {
      this.$data.error.isError = true;
      this.$data.error.message = message;
    }
  },

  data: () => {
    return {
      text: "",
      thisPage: "",
      errors: {
        isError: false,
        message: ""
      }
    };
  }
});
</script>