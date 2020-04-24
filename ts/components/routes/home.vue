<template>
  <div>
    <headder></headder>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import {default as headder} from '../parts/header.vue';
export default Vue.extend({
  components: {
    headder
  },
})
</script>