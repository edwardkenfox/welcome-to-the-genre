<template>
  <div class="container">
    <Help></Help>
    <form action="http://localhost:8081/" method="post" enctype="multipart/form-data">
      <!-- TODO: fix to allow directly uploading files -->
      <!-- <label class="button" for="file">Choose file</label>
      <input id="file" type="file" name="file" accept="audio/*" capture style="display: none;" @change="submit"> -->
      <div class="button" v-if="!isRecording" @click.prevent="startRecording">
        Start
      </div>
      <div class="button" v-else @click.prevent="stopRecording">Stop</div>
    </form>
  </div>
</template>
<script>
import Help from './Help'
export default {
  name: 'SourceSelect',

  components: {
    Help
  },

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
      // TODO: fix to allow directly uploading files
      // if (!this.$el.querySelector('#file').value && !this.audioBlob) {
      //   return false
      // }
      return true
    },
    startRecording () {
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
            this.submit()
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
      .then(res => res.json())
      .then(json => {
        this.$router.push({ path: `results/${json.genre}` })
      })
    }
  }
}
</script>
