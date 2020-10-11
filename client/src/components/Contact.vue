<template>
  <div class="contact">
    <h1>{{ msg }}</h1>
    <p>
        Please complete the form to get in touch.
    </p>
    <div>
      <b-form @submit="onSubmit" @reset="onReset" v-if="show">
      <b-form-group id="input-group-1" label="Email:" label-for="input-1" description="Your email address stays with us.">
        <b-form-input
          id="input-1"
          v-model="form.email"
          type="email"
          required
          placeholder="enter your email address"
        ></b-form-input>
      </b-form-group>

      <b-form-group id="input-group-2" label="Name:" label-for="input-2">
        <b-form-input
          id="input-2"
          v-model="form.name"
          required
          placeholder="enter your name"
        ></b-form-input>
      </b-form-group>

      <b-form-group id="input-group-3" label="Message:">
          <b-form-textarea
            id="input-3"
            v-model="form.details"
            placeholder="write your message"
            rows="3"
            max-rows="6"
          ></b-form-textarea>

        <pre class="mt-3 mb-0">{{ text }}</pre>
      </b-form-group>

      <b-button class="button" type="submit" variant="outline-dark">Send</b-button>
      <b-button class="button" type="reset" variant="outline-dark">Clear</b-button>
    </b-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Contact',
  props: {
    msg: String
  },
  data() {
      return {
        form: {
          email: '',
          name: '',
          details: ''
        },
        show: true
      }
    },
    methods: {
      onSubmit(evt) {
        evt.preventDefault()
        alert(JSON.stringify(this.form))
      },
      onReset(evt) {
        evt.preventDefault()
        // Reset our form values
        this.form.email = ''
        this.form.name = ''
        this.form.details = ''
        this.show = false
        this.$nextTick(() => {
          this.show = true
        })
      }
    }
}
</script>

<style scoped>

.contact {
  padding-left: 10px;
  padding-right: 10px;
  padding-bottom: 20px;

  max-width: 380px;
}

.button {
  padding: 5px 15px;
  margin-right: 10px;
}

</style>
