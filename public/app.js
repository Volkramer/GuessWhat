new Vue({
  el: "#app",
  data: {
    ws: null,
    newUsername: "",
    username: "",
    connected: false
  },
  mounted() {
    this.ws = new WebSocket("ws://" + window.location.host + "/ws");
    this.ws.onopen = () => {};
    this.ws.onmessage = e => {
      var dataJson = JSON.parse(e.data);
      console.log(dataJson);
    };
    this.ws.onclose = e => {
      console.log(e.data);
    };
  },
  methods: {
    join() {
      if (!(this.newUsername == null || this.newUsername == "")) {
        this.username = this.newUsername;
        this.connected = true;
        data = {
          event: "clientJoined",
          username: this.username,
        }
        console.log(data);
        this.ws.send(JSON.stringify(data));
      }
    }
  }
});