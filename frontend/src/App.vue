<template>
  <div id="app">
    <img src="./assets/logo.png">
    <input id="file" type="file" name="file" value="">
    <button id="submit" type="button" @click="submit">送信</button>
  </div>
</template>
<script>
export default {
  name: 'app',
  methods: {
    submit () {
      const input = document.getElementById('file')
      if (input.files.length === 0) return false;
      const url = 'http://localhost:8080/file'
      fetch(url, {
        method: 'POST',
        body: {
          file: input.files[0]
        }
      }).then(function (response) {
        return response.json()
      }).then(function (json) {
        console.log(json)
      })
    }
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
