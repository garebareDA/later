<template>
  <div>
      <v-text-field label="タイトル" type="text" v-model="title" />
      <v-row>
        <v-col>
          <v-textarea filled label="※MarkDown記法で書いてください" auto-grow v-model="text"></v-textarea>
        </v-col>
        <v-col>
          <div v-html="markdown"></div>
        </v-col>
      </v-row>
      <v-btn v-if="posted == false" outlined style="float: right;" v-on:click="draftPost" class="ma-2" >下書き保存</v-btn>
      <v-btn v-if="posted == true" outlined style="float: right;" disabled class="ma-2" >下書き保存</v-btn>
      <v-btn outlined style="float: right;" class="ma-2" >公開</v-btn>
      <div v-if="error">{{errorMessage}}</div>
      <div v-if="post">保存しました</div>
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
      post: false,
      posted:false,
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
      this.$data.posted = true;
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
        this.$data.posted = false;
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