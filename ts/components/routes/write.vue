<template>
  <div>
    <v-content>
      <v-container>
        <v-tabs centered v-model="tabs">
          <v-tab href="#edit">記事を書く</v-tab>
          <v-tab href="#item">書いた記事</v-tab>
          <v-tab href="#draft">下書き</v-tab>
          <v-tab heref="#like">いいねした記事</v-tab>
        </v-tabs>

        <v-tabs-items v-model="tabs">
          <v-tab-item value="edit">
            <editor :titles="title" :texts="text" :uuids="uuid" />
          </v-tab-item>

          <v-tab-item value="item">
            <v-card flat>
              <v-card-text>記事</v-card-text>
            </v-card>
          </v-tab-item>

          <v-tab-item value="draft">
            <v-card flat>
              <drafts @change="changeToEdit" />
            </v-card>
          </v-tab-item>

          <v-tab-item value="like">
            <v-card flat>
              <v-card-text>いいね</v-card-text>
            </v-card>
          </v-tab-item>
        </v-tabs-items>
      </v-container>
    </v-content>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import editor from "../parts/edit.vue";
import drafts from "../parts/drafts.vue";
export default Vue.extend({
  data: () => {
    return {
      title: "",
      text: "",
      uuid: null,
      tabs: "edit"
    };
  },

  methods: {
    changeToEdit: function(uuid: string, title: string, content: string) {
      this.$data.title = title;
      this.$data.text = content;
      this.$data.uuid = uuid;
      this.$data.tabs = "edit";
    }
  },

  components: {
    editor,
    drafts
  }
});
</script>