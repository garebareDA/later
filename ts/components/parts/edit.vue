<template>
  <div>
    <v-text-field label="タイトル" type="text" v-model="title" />
    <v-row>
      <v-col>
        <v-textarea filled label="※MarkDown記法で書いてください" auto-grow v-model="text" class="ma-2"></v-textarea>
      </v-col>
      <v-col>
        <div v-html="markdown" lass="ma-3"></div>
      </v-col>
    </v-row>
    <v-row style="float: right;">
      <v-btn
        outlined
        style="float: right;"
        v-on:click="draftPost"
        class="ma-2"
        :disabled="posted == true"
      >下書き保存</v-btn>
      <v-btn
        outlined
        style="float: right;"
        class="ma-2"
        v-on:click="publicPost"
        :disabled="posted == true"
      >公開</v-btn>
    </v-row>

    <v-row style="float: left;">
      <v-btn outlined style="float: left;" class="ma-2" v-on:click="newEdit">新規作成</v-btn>
    </v-row>

    <div v-if="error" style="float: right;" class="ma-2">{{errorMessage}}</div>
    <div v-if="post" style="float: right;" class="ma-2">{{foundMessage}}</div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import marked from "marked";
import createDOMPurify from "dompurify";
import * as firebase from "firebase/app";
import "firebase/auth";
import { v4 as uuidv4 } from "uuid";
import axios, { AxiosResponse, AxiosError } from "axios";

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
      user.getIdToken(true).then(token => {
        const url = "/draft";
        const draftParams = {
          token: token,
          title: this.$data.title,
          content: this.$data.text,
          draftID: this.$data.uuid
        };
        axios
          .post(url, draftParams)
          .then((res: AxiosResponse) => {
            this.$data.post = true;
            this.$data.posted = false;
            this.$data.foundMessage = res.data.status;
            this.$data.errorMessage = "";
            this.$emit("drafts");
          })
          .catch((error: AxiosError) => {
            this.$data.error = true;
            this.$data.posted = false;
            this.$data.foundMessage = "";
            if (error.response?.data.error != undefined) {
              this.$data.errorMessage = error.response?.data.error;
            }else{
              this.$data.errorMessage = "不明なエラー"
            }
          });
      });
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
      user.getIdToken(true).then(token => {
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
            this.$data.posted = false;
            this.$data.foundMessage = res.data.status;
            this.$data.errorMessage = "";
            this.$emit("publics");
          })
          .catch(error => {
            this.$data.error = true;
            this.$data.posted = false;
            this.$data.foundMessage = "";
            if (error.response?.data.error != undefined) {
              this.$data.errorMessage = error.response?.data.error;
            }else{
              this.$data.errorMessage = "不明なエラー"
            }
          });
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
      this.$data.markdown = createDOMPurify.sanitize(html);
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