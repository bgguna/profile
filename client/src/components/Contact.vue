<template>
  <div class="contact">
    <h1>{{ msg }}</h1>
    <br />
    <p>
      Please complete the form to get in touch.<br />
      Your details stay with me.
    </p>
    <div>
      <b-form @submit="onSubmit" @reset="onReset" v-if="show">
        <b-form-group id="input-group-1" label="Email:" label-for="input-1">
          <div class="input-div">
            <b-form-group-prepend is-text>
              <b-icon
                class="input-icon"
                icon="envelope"
                font-scale="1.5"
              ></b-icon>
            </b-form-group-prepend>
            <b-form-input
              id="input-1"
              v-model="form.email"
              type="email"
              required
              placeholder="enter your email address"
            ></b-form-input>
          </div>
        </b-form-group>

        <b-form-group id="input-group-2" label="Phone:" label-for="input-2">
          <div class="input-div">
            <b-form-group-prepend is-text>
              <b-icon
                class="input-icon"
                icon="telephone"
                font-scale="1.5"
              ></b-icon>
            </b-form-group-prepend>
            <b-form-input
              id="input-2"
              v-model="form.phone"
              placeholder="enter your phone number"
            ></b-form-input>
          </div>
        </b-form-group>

        <b-form-group id="input-group-3" label="Name:" label-for="input-3">
          <div class="input-div">
            <b-form-group-prepend is-text>
              <b-icon
                class="input-icon"
                icon="person-fill"
                font-scale="1.5"
              ></b-icon>
            </b-form-group-prepend>
            <b-form-input
              id="input-3"
              v-model="form.name"
              required
              placeholder="enter your name"
            ></b-form-input>
          </div>
        </b-form-group>

        <b-form-group class="input-div" id="input-group-4" label="Message:">
          <b-form-textarea
            id="input-4"
            v-model="form.message"
            placeholder="write your message"
            rows="3"
            max-rows="8"
            required
          ></b-form-textarea>
          <pre class="mt-3 mb-0">{{ text }}</pre>
        </b-form-group>

        <b-button class="button" type="submit" variant="outline-dark"
          >Send</b-button
        >
        <b-button class="button" type="reset" variant="outline-dark"
          >Clear</b-button
        >
      </b-form>
    </div>
  </div>
</template>

<script>
export default {
  name: "Contact",
  props: {
    msg: String,
  },
  data() {
    return {
      form: {
        email: "",
        phone: "",
        name: "",
        message: "",
      },
      show: true,
    };
  },
  methods: {
    onSubmit(evt) {
      evt.preventDefault();
      const uri = "http://127.0.0.1:3000/send";
      const contact = {
        name: this.form.name,
        phone: this.form.phone,
        email: this.form.email,
        message: this.form.message,
      };
      const config = {
        headers: {
          "Content-type": "application/json; charset=UTF-8",
        },
      };
      this.axios.post(uri, contact, config).then((response) => {
        if (response.status === 200) {
          console.log(
            "Successfully sent contact message: " + JSON.stringify(this.form)
          );
          this.makeToast("success");
          this.onReset();
        } else {
          console.log(
            "Failed to send contact message. HTTP Response: " +
              response.status +
              ". Status: " +
              response.statusText
          );
          this.makeToast("danger");
          this.onReset();
        }
      });
    },
    onReset(evt) {
      if (evt) {
        evt.preventDefault();
      }

      // Reset our form values
      this.form.email = "";
      this.form.phone = "";
      this.form.name = "";
      this.form.message = "";
      this.show = false;
      this.$nextTick(() => {
        this.show = true;
      });
    },
    makeToast(variant = null) {
      if (variant === "success") {
        this.$bvToast.toast("Your message has been sent!", {
          title: "Success!",
          variant: variant,
          solid: true,
        });
      }

      if (variant === "danger") {
        this.$bvToast.toast("We failed to sent your message 😢.", {
          title: "Error!",
          variant: variant,
          solid: true,
        });
      }
    },
  },
};
</script>

<style scoped>
.contact {
  padding-left: 10px;
  padding-right: 10px;
  padding-bottom: 20px;

  max-width: 450px;
  margin: 20px auto;
}

.input-div {
  display: flex;
  flex-direction: row;
}

.input-icon {
  margin-top: 5px;
  margin-right: 10px;
}

.button {
  padding: 5px 15px;
  margin-right: 10px;
}
</style>
