<template>
  <div class="container">
    <Help></Help>
    <div class="title">Welcome<br>to<br>the<br>genre</div>
    <p>Let us hear music, you shall know the genre of it.</p>
    <router-link class="button" to="/source">Play</router-link>
  </div>
</template>
<script>
import Help from './Help.vue'
export default {
  name: 'Top',
  components: {
    Help
  }
}
</script>
