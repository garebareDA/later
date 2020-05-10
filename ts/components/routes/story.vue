<template>
  <div>
    <v-content>
      <v-container>
        <v-row align="center" justify="center">
          <v-col cols="20" sm="8" md="20">
            <v-card>
              <h1 v-if="errorMessage != ''">{{errorMessage}}</h1>
              <v-btn
                color="pink"
                v-if="isLogin"
                :outlined="!like"
                class="ma-6"
                style="float: right;"
                v-on:click="likes()"
                dark
              >いいね</v-btn>
              <v-card-title class="display-2 text--primary">{{title}}</v-card-title>
              <v-card-subtitle>by {{userName}}</v-card-subtitle>
              <div v-html="content" class="text--primary ma-3" />
              <v-divider />
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import axios from "axios";
import marked from "marked";
import createDOMPurify from "dompurify";
import * as firebase from "firebase/app";
import "firebase/auth";
export default Vue.extend({
  created: function() {
    const id = this.$route.params.id;
    const url = "/story";
    const params = {
      uuid: id
    };
    axios.get(url, { params: params }).then(res => {
      const html = marked(res.data.Content, { breaks: true });
      this.$data.content = createDOMPurify.sanitize(html);
      this.$data.title = res.data.Title;
      this.$data.userName = res.data.UserName;
    });

    const _this = this;
    firebase.auth().onAuthStateChanged(user => {
      _this.$data.isLogin = user ? true : false;
      if (!user) {
        return;
      }
      user.getIdToken(true).then(token => {
        const url = "/like";
        const params = {
          uuid: id,
          token: token
        };
        axios
          .get(url, { params: params })
          .then(res => {
            console.log(res);
            if (res.data.status === "like") {
              _this.$data.like = true;
            } else {
              _this.$data.like = false;
            }
          })
          .catch(error => {
            _this.$data.errorMessage = error.response?.data.error;
          });
      });
    });
  },

  methods: {
    likes: function() {
      const id = this.$route.params.id;
      const user = firebase.auth().currentUser;
      const url = "/like";

      if (!user) {
        this.$data.error = true;
        this.$data.errorMessage = "ログインしていません";
        return;
      }

      user.getIdToken(true).then(token => {
        const params = {
          token: token,
          uuid: id
        };

        if (this.$data.like == true) {
          this.likeRemove(params);
          return;
        }

        axios
          .post(url, params)
          .then(res => {
            this.$data.like = true;
          })
          .catch(err => {
            console.log(err);
          });
      });
    },

    likeRemove(params:Object) {
      const url = "/like"
      axios
        .delete(url, {data:params})
        .then(res => {
          this.$data.like = false;
        })
        .catch(err => {
          console.log(err);
        });
    }
  },

  data: () => {
    return {
      content: "",
      title: "",
      userName: "",
      isLogin: false,
      like: false,
      error: {
        err: false,
        message: ""
      }
    };
  }
});
</script>