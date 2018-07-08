<template>
  <div id="app">
    <form action="http://localhost:8081/" method="post" enctype="multipart/form-data">
      <img src="./assets/logo.png">
      <audio id="player" controls></audio>
      <button v-if="!isRecording" @click.prevent="startRecording">Start</button>
      <button v-else @click.prevent="stopRecording">Stop</button>
      <input id="file" type="file" name="file" accept="audio/*" capture>
      <button id="submit" type="submit" @click.prevent="submit">送信</button>
    </form>
  </div>
</template>
<script>
export default {
  name: 'app',

  data () {
    return {
      isRecording: false,
      audioChunks: [],
      mediaRecorder: null,
      audioBlob: null
    }
  },

  methods: {
    audioExists () {
      if (!this.$el.querySelector('#file').value && !this.audioBlob) {
        return false
      }
      return true
    },
    startRecording () {
      const player = document.getElementById('player')
      player.src = null

      this.isRecording = true

      navigator.mediaDevices.getUserMedia({ audio: true, video: false })
        .then(stream => {
          console.log('start recording')

          this.mediaRecorder = new MediaRecorder(stream)
          this.mediaRecorder.start()

          this.mediaRecorder.addEventListener('dataavailable', event => {
            this.audioChunks.push(event.data)
          })

          this.mediaRecorder.addEventListener('stop', () => {
            this.audioBlob = new Blob(this.audioChunks)
            const audioUrl = URL.createObjectURL(this.audioBlob)
            player.src = audioUrl
          })
        })
    },
    stopRecording () {
      if (!this.mediaRecorder) return false
      console.log('stop recording')
      this.isRecording = false
      this.mediaRecorder.stop()
    },
    submit () {
      if (!this.audioExists()) {
        console.error('You must record audio or select an audio file to submit')
        return false
      }

      var fd = new FormData()
      fd.append('file', this.audioBlob, 'myaudio.txt')

      fetch('http://localhost:8081/', {
        method: 'POST',
        body: fd
      })
    }
  },

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
