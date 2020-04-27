<template>
  <div>
      <textarea :value="text" ></textarea>
      <div v-html="compiledMarkdown"></div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import marked from "marked";
export default Vue.extend({
  data() {
    return {
      text: ""
    };
  },

  computed: {
    compiledMarkdown: function() {
      const html = marked(this.$data.text, { sanitize: true });
      console.log(html);
      return html;
    }
  },


});
</script>