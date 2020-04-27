<template>
  <div>
    <v-content>
      <v-container>
            <v-tabs centered>
              <v-tab>記事を書く</v-tab>
              <v-tab>書いた記事</v-tab>
              <v-tab>下書き</v-tab>
              <v-tab>いいねした記事</v-tab>

            <v-tab-item>
              <editor/>
            </v-tab-item>

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