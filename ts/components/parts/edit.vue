<template>
  <div>
    <v-container>
      <v-text-field label="タイトル" type="text" v-model="title" />
      <v-row>
        <v-col>
          <v-textarea filled label="※MarkDown記法で書いてください" auto-grow v-model="text"></v-textarea>
        </v-col>
        <v-col>
          <div v-html="markdown"></div>
        </v-col>
      </v-row>
      <v-btn outlined style="float: right;" v-on:click="draftPost">下書き保存</v-btn>
      <v-btn outlined style="float: right;">公開</v-btn>
      <div v-if="error">{{errorMessage}}</div>
      <div v-if="post">保存しました</div>
    </v-container>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import marked from "marked";
import * as firebase from "firebase/app";
import { v4 as uuidv4 } from "uuid";
import "firebase/auth";
import axios from "axios";

export default Vue.extend({
  data: () => {
    return {
      text: "",
      title: "",
      markdown: "",
      uuid: "",
      error: false,
      errorMessage: "",
      post: false
    };
  },

  created: function() {
    const id = this.$route.params.id;
    if (id === undefined) {
      this.$data.uuid = uuidv4();
    } else if (id == "") {
      this.$data.error = true;
      this.$data.errorMessage = "IDが指定されいてません";
    } else {
      this.$data.uuid = id;
    }
  },

  methods: {
    draftPost: async function() {
      const user = firebase.auth().currentUser;
      console.log(user);
      if (!user) {
        this.$data.error = true;
        this.$data.errorMessage = "ログインしていません";
        return;
      }
      this.$data.error = false;
      const token = await user.getIdToken(true);
      const url = "/draft";
      const draftParams = {
        token: token,
        title: this.$data.title,
        content: this.$data.text,
        draftID: this.$data.uuid
      };
      axios
        .post(url, draftParams)
        .then(() => {
          this.$data.post = true;
        })
        .catch(error => {
          this.$data.error = true;
          this.$data.errorMessage = error;
        });
    }
  },

  watch: {
    text: function() {
      const html = marked(this.$data.text, { sanitize: true, breaks: true });
      this.$data.markdown = html;
      this.$data.post = false;
    }
  }
});
</script>