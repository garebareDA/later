<template>
  <div>
    <v-container>
      <v-text-field label="タイトル" type="text" />
      <v-row>
        <v-col>
          <v-textarea filled label="※MarkDown記法で書いてください" auto-grow v-model="text"></v-textarea>
        </v-col>
        <v-col>
          <div v-html="markdown"></div>
        </v-col>
      </v-row>
      <v-btn outlined style="float: right;">下書き保存</v-btn>
      <v-btn outlined style="float: right;">公開</v-btn>
    </v-container>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import marked from "marked";
export default Vue.extend({
  data: () => {
    return {
      text: "",
      markdown: ""
    };
  },

  watch: {
    text: function() {
      const html = marked(this.$data.text, { sanitize: true, breaks: true });
      this.$data.markdown = html;
    }
  }
});
</script>