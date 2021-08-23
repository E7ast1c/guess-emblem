<template>
  <main-template />
</template>

<script lang="ts">
  import Vue from 'vue'
  import MainTemplate from '../components/Template.vue'

  export default Vue.extend({
    name: 'Home',

    components: {
      MainTemplate,
    },
  })
</script>
