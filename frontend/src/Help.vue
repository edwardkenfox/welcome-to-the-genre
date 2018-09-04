<template>
  <div style="background-color: inherit;">
    <div class="help" v-if="!overlayShown" @click.stop="showOverlay">?</div>
    <div class="overlay" v-if="overlayShown">
      <div class="close" v-if="overlayShown" @click.stop="hideOverlay">☓</div>
      Lorem ipsum dolor sit amet, consectetur adipisicing elit. Quaerat, neque magnam explicabo dolorem necessitatibus delectus inventore quidem doloremque iure at. Quibusdam aperiam dolore fuga nihil, ut voluptate architecto harum alias.
    </div>
  </div>
</template>
<script>
export default {
  name: 'Help',
  data () {
    return {
      overlayShown: false
    }
  },
  methods: {
    showOverlay () {
      this.overlayShown = true
    },
    hideOverlay () {
      this.overlayShown = false
    }
  }
}
</script>
