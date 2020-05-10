<template>
  <div>
    <v-content>
      <v-container>
        <v-row align="center" justify="center">
          <v-col>
            <v-card class="mb-6 mx-auto" max-width="80%">
              <v-card-text>
                <v-list two-line>
                  <div class="posts" v-for="(item, index) in list" :key="index">
                    <v-list-item link :to="'/story/' + item.ID">
                      <v-list-item-title>{{item.Title}}</v-list-item-title>
                      <div>
                        <v-list-item-subtitle>by {{item.UserName}}</v-list-item-subtitle>
                      </div>
                      <v-divider />
                    </v-list-item>
                  </div>
                </v-list>
                <div v-if="errorMessage != ''">{{errorMessage}}</div>
                <infinite-loading spinner="spiral" @infinite="infiniteHandler">
                  <div slot="no-more"></div>
                  <div slot="no-results"></div>
                </infinite-loading>
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
import axios from "axios";
export default Vue.extend({
  methods: {
    infiniteHandler: function($state: any) {
      const url = "/homes";
      this.$data.getPublic += 10;
      const params = { number: this.$data.getPublic };
      axios
        .get(url, { params: params })
        .then(res => {
          if (res.data != "empty") {
            this.$data.list.push(...res.data.get);
          }

          if (res.data.continue === false) {
            $state.complete();
            return;
          }

          $state.loaded();
        })
        .catch(error => {
          console.log(error.response?.data.error);
          if (error.response?.data.error != undefined) {
            this.$data.errorMessage = error.response?.data.error;
          }
          $state.complete();
        });
    }
  },

  data: () => {
    return {
      getPublic: 0,
      list: [],
      errorMessage: ""
    };
  }
});
</script>