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
                    <v-btn
                      class="ma-2"
                      outlined
                      v-on:click="startEdit(item.ID, item.Title, item.Content)"
                    >編集</v-btn>
                    <v-btn class="ma-2" outlined v-on:click="removeDraft(item.ID)">削除</v-btn>
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
  methods: {
    infiniteHandler: async function($state: any) {
      const user = firebase.auth().currentUser;
      if (!user) {
        return;
      }

      user.getIdToken(true).then(token => {
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
            $state.loaded();
          })
          .catch(err => {
            $state.complete();
            console.log(err);
          });
      });
    },

    startEdit: function(uuid: string, title: string, content: string) {
      this.$emit("change", uuid, title, content);
    },

    removeDraft: async function(uuid: string) {
      this.$data.posted = true;
      const user = firebase.auth().currentUser;
      if (!user) {
        return;
      }
      this.$data.error = false;
      user.getIdToken(true).then(token => {
        const url = "/draft/remove";
        const params = {
          uuid: uuid,
          token: token
        };
        axios
          .delete(url, { data: params })
          .then(res => {
            this.$emit("reload");
          })
          .catch(error => {
            this.$router.push("/");
          });
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