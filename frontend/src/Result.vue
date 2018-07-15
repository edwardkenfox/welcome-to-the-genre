<template>
  <div class="container">
    <Help></Help>
    <div class="result-text">
      <div>That's a...</div>
      <div class="result-genre" v-text="genre"></div>
      <div>song!</div>
    </div>
    <router-link class="button" to="/source">Try again</router-link>
  </div>
</template>
<script>
import Help from './Help.vue'
export default {
  name: 'Result',
  components: {
    Help
  },
  props: {
    genre: {
      type: String,
      required: true
    }
  }
}
</script>
