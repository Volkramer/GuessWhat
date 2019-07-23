new Vue({
  el: "#app",
  data: {
    ws: null,
    newUsername: "",
    username: "",
    connected: false,
    messages: [],
    message: "",
    newMessage: "",
    clients: []
  },
  mounted() {
    this.ws = new WebSocket("ws://" + window.location.host + "/ws");
    this.ws.onopen = () => {};
    this.ws.onmessage = e => {
      var dataJson = JSON.parse(e.data);
      //console.log(dataJson);
      //console.log(this.messages);
      switch (dataJson.event) {
        case "clientJoined":
          this.username = dataJson.username;
          this.connected = true;
          break;
        case "system":
          this.message = "SYSTEM: " + dataJson.systemMessage;
          this.messages.push(this.message);
          break;
        case "message":
          this.message = dataJson.username + ": " + dataJson.message;
          this.messages.push(this.message);
          break;
        case "clientList":
          this.clients = dataJson.clients;
          break;
        default:
          break;
      }
    };
    this.ws.onclose = e => {
      console.log(e.data);
    };
  },
  methods: {
    join() {
      if (!(this.newUsername == null || this.newUsername == "")) {
        data = {
          event: "clientJoined",
          username: this.newUsername
        };
        //console.log(data);
        this.ws.send(JSON.stringify(data));
      }
    },
    postMessage() {
      if (!(this.newMessage == null || this.newMessage == "")) {
        data = {
          event: "message",
          username: this.username,
          message: this.newMessage
        };
        this.ws.send(JSON.stringify(data));
        this.newMessage = "";
      }
    }
  }
});