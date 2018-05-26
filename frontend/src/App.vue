<template>
  <div id="app">
    <form action="http://localhost:8081/" method="post" enctype="multipart/form-data">
      <img src="./assets/logo.png">
      <input id="file" type="file" name="file" value="">
      <button id="submit" type="submit">送信</button>
    </form>
  </div>
</template>
<script>
// Vueインスタンスの存在意義...
export default {
  name: 'app',
  mounted () {
    window.vm = this
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
