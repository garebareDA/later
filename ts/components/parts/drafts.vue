<template>
  <div>
    <v-container>
      <v-row align="center" justify="center">
        <v-col>
          <v-card class="elevation-12">
            <v-card-text>
              <v-list two-line>
                <div class="posts" v-for="(item, index) in list" :key="index">
                  <v-list-item :link="parts == 'item'">
                    <v-list-item-title>{{item.Title}}</v-list-item-title>
                    <v-spacer />
                    <v-btn class="ma-2" outlined v-if="parts == 'item'">見る</v-btn>
                    <v-btn
                      class="ma-2"
                      outlined
                      v-on:click="startEdit(item.ID, item.Title, item.Content)"
                    >編集</v-btn>
                    <v-btn class="ma-2" outlined v-on:click="removeDraft(item.ID, parts)">削除</v-btn>
                  </v-list-item>
                  <v-divider />
                </div>
              </v-list>
              <infinite-loading
                spinner="spiral"
                @infinite="infiniteHandler"
                v-if="parts === 'draft'"
              >
                <div slot="no-more"></div>
                <div slot="no-results"></div>
              </infinite-loading>

              <infinite-loading
                spinner="spiral"
                @infinite="infiniteHandlerItem"
                v-if="parts === 'item'"
              >
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
  props: ["parts"],
  methods: {
    infiniteHandler: function($state: any) {
      const url = "/drafts";
      this.infiniteGet(url, $state);
    },

    infiniteHandlerItem: function($state: any) {
      const url = "/publics";
      this.infiniteGet(url, $state);
    },

    infiniteGet: function(url: string, $state: any) {
      const user = firebase.auth().currentUser;
      if (!user) {
        return;
      }

      user.getIdToken(true).then(token => {
        this.$data.getNumber += 10;
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

    removeDraft: async function(uuid: string, parts: any) {
      if (parts === "draft") {
        const url = "/draft/remove";
        this.delete(url, uuid);
      }

      if (parts === "item") {
        const url = "/public/remove";
        this.delete(url, uuid);
      }
    },

    delete: function(url: string, uuid: string) {
      this.$data.posted = true;
      const user = firebase.auth().currentUser;
      if (!user) {
        return;
      }
      this.$data.error = false;
      user.getIdToken(true).then(token => {
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
    },

    startEdit: function(uuid: string, title: string, content: string) {
      this.$emit("change", uuid, title, content);
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