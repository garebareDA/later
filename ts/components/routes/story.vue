<template>
  <div>
    <v-content>
      <v-container>
        <v-row align="center" justify="center">
          <v-col cols="20" sm="8" md="20">
            <v-card>
              <v-btn outlined class="ma-6" style="float: right;">いいね</v-btn>
              <v-card-title class="display-2 text--primary">{{title}}</v-card-title>
              <v-card-subtitle>by {{userName}}</v-card-subtitle>
              <div v-html="content" class="text--primary ma-3" />
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
import marked from "marked";
import createDOMPurify from "dompurify";
export default Vue.extend({
  created: function() {
    const id = this.$route.params.id;
    const url = "/story";
    const params = {
      uuid: id
    };
    axios.get(url, { params: params }).then(res => {
      const html = marked(res.data.Content, { breaks: true });
      this.$data.content = createDOMPurify.sanitize(html);
      this.$data.title = res.data.Title;
      this.$data.userName = res.data.UserName;
    });
  },

  data: () => {
    return {
      content: "",
      title: "",
      userName: "",
      error: {
        err: false,
        message: ""
      }
    };
  }
});
</script>