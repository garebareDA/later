<template>
  <div>
    <v-content>
      <v-container>
        <editor/>
      </v-container>
    </v-content>
    <v-content>
      <v-container>
        <v-row align="center" justify="center">
          <v-col cols="12" sm="8" md="4">
            <v-tabs>
              <v-tab>書いた記事</v-tab>
              <v-tab>下書き</v-tab>
              <v-tab>いいねした記事</v-tab>
            <v-tab-item>
              <v-card flat>
                <v-card-text>記事</v-card-text>
              </v-card>
            </v-tab-item>

            <v-tab-item>
              <v-card flat>
                <v-card-text>下書き</v-card-text>
              </v-card>
            </v-tab-item>

            <v-tab-item>
              <v-card flat>
                <v-card-text>いいね</v-card-text>
              </v-card>
            </v-tab-item>
            </v-tabs>
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import editor from "../parts/edit.vue";
export default Vue.extend({
  components:{
    editor
  },
});
</script>