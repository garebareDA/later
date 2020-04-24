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
                  <v-text-field label="Eメール" type="text" />
                  <v-text-field label="パスワード" type="password" />
                </v-form>
                <v-card-actions v-if="login.isLogin == true">
                  <v-btn outlined>新規登録</v-btn>
                  <v-spacer />
                  <v-btn outlined>ログイン</v-btn>
                </v-card-actions>

                <v-card-actions v-if="login.isLogin == false">
                  <v-spacer />
                  <v-btn outlined>登録</v-btn>
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
import Vue from 'vue'
export default Vue.extend({
  name:'login',
  props: ['login']
})
</script>