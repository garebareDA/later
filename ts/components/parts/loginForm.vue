<template>
  <div>
    <v-content>
      <v-container>
        <v-row align="center" justify="center">
          <v-col cols="12" sm="8" md="4">
            <v-card class="elevation-12">
              <v-toolbar>
                <v-spacer />
                <v-toolbar-title v-if="login.isLogin == true">ログイン</v-toolbar-title>
                <v-toolbar-title v-if="login.isLogin == false">新規登録</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <v-form>
                  <v-text-field
                    label="ユーザー名"
                    type="text"
                    v-model="model.userName"
                    v-if="login.isLogin == false"
                  />
                  <v-text-field label="Eメール" type="text" v-model="model.eMail" />
                  <v-text-field label="パスワード" type="password" v-model="model.passWord" />
                  <v-text-field
                    label="パスワードの再入力"
                    type="password"
                    v-model="model.reEnter"
                    v-if="login.isLogin == false"
                  />
                </v-form>
                <v-card-actions v-if="login.isLogin == true">
                  <v-btn outlined v-on:click="push('singUp')">新規登録</v-btn>
                  <v-spacer />
                  <v-btn outlined v-on:click="loginUser()">ログイン</v-btn>
                </v-card-actions>

                <v-card-actions v-show="error.isError == true">
                  <v-spacer />
                  <div class="red--text">{{error.message}}</div>
                </v-card-actions>

                <v-card-actions v-if="login.isLogin == false">
                  <v-spacer />
                  <v-btn outlined v-on:click="singUp()">登録</v-btn>
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
  name: "login",
  props: ["login"],

  methods: {
    push: function(url: string) {
      this.$router.push(url);
    },

    errors: function(message: string): void {
      this.$data.error.isError = true;
      this.$data.error.message = message;
    },

    judge: function(login: boolean): boolean {
      const email = this.$data.model.eMail;
      const passWord = this.$data.model.passWord;
      const userName = this.$data.model.userName;
      const reEnter = this.$data.model.reEnter;

      if (userName === "" && login == true) {
        this.errors("ユーザー名を入力してください");
        return false;
      }

      if (email === "") {
        this.errors("メールアドレスを入力してください");
        return false;
      }

      var mail_regex1 = new RegExp(
        "(?:[-!#-'*+/-9=?A-Z^-~]+.?(?:.[-!#-'*+/-9=?A-Z^-~]+)*|\"(?:[!#-[]-~]|\\\\[\x09 -~])*\")@[-!#-'*+/-9=?A-Z^-~]+(?:.[-!#-'*+/-9=?A-Z^-~]+)*"
      );
      var mail_regex2 = new RegExp("^[^@]+@[^@]+$");
      if (email.match(mail_regex1) && email.match(mail_regex2)) {
        if (
          email.match(
            /[^a-zA-Z0-9\!\"\#\$\%\&\'\(\)\=\~\|\-\^\\\@\[\;\:\]\,\.\/\\\<\>\?\_\`\{\+\*\} ]/
          )
        ) {
          this.errors("メールアドレスのフォーマットが不明です");
          return false;
        }

        if (!email.match(/\.[a-z]+$/)) {
          this.errors("メールアドレスのフォーマットが不明です");
          return false;
        }
      } else {
        this.errors("メールアドレスのフォーマットが不明です");
        return false;
      }

      if (passWord === "") {
        this.errors("パスワードを入力してください");
        return false;
      }

      if (passWord != reEnter && login == true) {
        this.errors("パスワードが一致しません");
        return false;
      }

      return true;
    },

    singUp: async function() {
      const email = this.$data.model.eMail;
      const passWord = this.$data.model.passWord;
      const userName = this.$data.model.userName;

      if (!this.judge(true)) {
        return;
      }

      try {
        const providers = await firebase
          .auth()
          .fetchSignInMethodsForEmail(email);
        if (
          providers.findIndex(
            p =>
              p ===
              firebase.auth.EmailAuthProvider.EMAIL_PASSWORD_SIGN_IN_METHOD
          ) !== -1
        ) {
          this.errors("すでに登録されているようです");
          return;
        }

        await firebase.auth().createUserWithEmailAndPassword(email, passWord);
        const user = firebase.auth().currentUser;
        user?.sendEmailVerification().then(() => {
          user.updateProfile({
            displayName: userName
          });
        });

        this.$router.push("/");
      } catch (error) {
        console.log("error:" + error.message);
        this.errors("エラーが発生しました");
        return;
      }
    },

    loginUser: async function() {
      const email = this.$data.model.eMail;
      const passWord = this.$data.model.passWord;
      const _this = this;

      if (!this.judge(false)) {
        return;
      }

      firebase
        .auth()
        .setPersistence(firebase.auth.Auth.Persistence.LOCAL)
        .then(() => {
          return firebase
            .auth()
            .signInWithEmailAndPassword(email, passWord)
            .then(() => {
              this.$router.push("/");
            })
            .catch((err) => {
              console.log(err);
              this.errors("エラーが発生しました");
            });
        });

        await firebase.auth().onAuthStateChanged((users) => {
        if(users){
          this.$router.push('/');
        }
      });
    }
  },

  data: () => {
    return {
      model: {
        eMail: "",
        passWord: "",
        reEnter: "",
        userName: ""
      },

      error: {
        isError: false,
        message: ""
      }
    };
  }
});
</script>