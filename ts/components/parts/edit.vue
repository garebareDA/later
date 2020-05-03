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
    <v-row style="float: right;">
      <div v-if="error" style="float: right; ma-2">{{errorMessage}}</div>
      <div v-if="post" style="float: right; ma-2">{{foundMessage}}</div>
      <v-btn
        v-if="posted == false"
        outlined
        style="float: right;"
        v-on:click="draftPost"
        class="ma-2"
      >下書き保存</v-btn>
      <v-btn v-if="posted == true" outlined style="float: right;" disabled class="ma-2">下書き保存</v-btn>
      <v-btn
        v-if="posted == false"
        outlined
        style="float: right;"
        class="ma-2"
        v-on:click="publicPost"
      >公開</v-btn>
      <v-btn v-if="posted == true" disabled  outlined style="float: right;" class="ma-2">公開</v-btn>
    </v-row>

    <v-row style="float: left;">
      <v-btn outlined style="float: left;" class="ma-2" v-on:click="newEdit">新規作成</v-btn>
    </v-row>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import marked from "marked";
import sanitize from "sanitize-html";
import * as firebase from "firebase/app";
import "firebase/auth";
import { v4 as uuidv4 } from "uuid";
import axios from "axios";

export default Vue.extend({
  props: ["texts", "titles", "uuids"],

  data: () => {
    return {
      text: "",
      title: "",
      uuid: null,
      markdown: "",
      error: false,
      errorMessage: "",
      post: false,
      posted: false,
      foundMessage: ""
    };
  },

  created: function() {
    const id = this.uuids;
    this.$data.text = this.texts;
    this.$data.title = this.titles;

    if (id === null) {
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
        .then(res => {
          this.$data.post = true;
          this.$data.foundMessage = res.data.status;
        })
        .catch(error => {
          this.$data.error = true;
          this.$data.errorMessage = error;
        });
      this.$data.posted = false;
    },

    publicPost: async function() {
      this.$data.posted = true;
      const user = firebase.auth().currentUser;
      if (!user) {
        this.$data.error = true;
        this.$data.errorMessage = "ログインしていません";
        return;
      }
      this.$data.error = false;
      const token = await user.getIdToken(true);

      const url = "/public";
      const draftParams = {
        token: token,
        title: this.$data.title,
        content: this.$data.text,
        uuid: this.$data.uuid
      };
      axios
        .post(url, draftParams)
        .then(res => {
          this.$data.post = true;
          this.$data.foundMessage = res.data.status;
        })
        .catch(error => {
          this.$data.error = true;
          this.$data.posted = false;
          this.$data.errorMessage = error;
        });
    },
    newEdit: function() {
      this.$data.uuid = uuidv4();
      this.$data.text = "";
      this.$data.title = "";
    }
  },

  watch: {
    text: function() {
      const html = marked(this.$data.text, { breaks: true });
      this.$data.markdown = sanitize(html);
      this.$data.post = false;
    },

    texts: function() {
      this.$data.uuid = this.uuids;
      this.$data.text = this.texts;
      this.$data.title = this.titles;
    }
  }
});
</script>