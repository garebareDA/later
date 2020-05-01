<template>
  <div>
    <v-container>
      <v-row align="center" justify="center">
        <v-col>
          <v-card class="elevation-12">
            <v-card-text>
              <v-list two-line>
                <div class="posts" v-for="(item, index) in list" :key="index">
                  <v-list-item>
                    <v-list-item-title>{{item.Title}}</v-list-item-title>
                    <v-spacer />
                    <v-btn class="ma-2" outlined>編集</v-btn>
                    <v-btn class="ma-2" outlined>削除</v-btn>
                  </v-list-item>
                  <v-divider />
                </div>
              </v-list>
              <infinite-loading ref="infiniteLoading" spinner="spiral" @infinite="infiniteHandler">
                <div slot="no-more"></div>
                <div slot="no-results"></div>
              </infinite-loading>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>
<script lang="ts">
import Vue from "vue";
import axios from "axios";
import * as firebase from "firebase/app";
import "firebase/auth";

export default Vue.extend({
  created: function() {},
  methods: {
    infiniteHandler: async function($state: any) {
      const user = firebase.auth().currentUser;
      if (!user) {
        return;
      }
      const token = await user.getIdToken(true);
      this.$data.getNumber += 10;
      const url = "/drafts";
      const params = {
        number: this.$data.getNumber,
        token: token
      };
      axios
        .get(url, { params: params })
        .then(res => {
          this.$data.list.push(...res.data);
          if (res.data.length < 10) {
            $state.complete();
          }
        })
        .catch(err => {
          console.log(err);
        });
    }
  },

  data: function() {
    return {
      getNumber: 0,
      list: []
    };
  }
});
</script>