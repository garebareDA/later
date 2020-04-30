<template>
  <div>
    <v-container>
      <v-row align="center" justify="center">
        <v-col>
          <v-card class="elevation-12">
            <v-card-text>
              <v-list-item>
                <v-list-item-title>
                  title
                </v-list-item-title>
                <v-spacer />
                <v-btn outlined>編集</v-btn>
                <v-btn outlined>公開</v-btn>
              </v-list-item>
              <v-divider />
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({});
</script>