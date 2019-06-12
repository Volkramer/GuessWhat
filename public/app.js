new Vue({
  el: "#app",
  data: {
    ws: null,
    roomSelected: "",
    rooms: [],
    room: "",
    username: "",
    dialog: true,
    isRegistered: false,
  },
  mounted() {
    this.ws = new WebSocket("ws://" + window.location.host + "/ws");
    this.ws.onopen = () => {};
    this.ws.onmessage = e => {
      var dataJson = JSON.parse(e.data);
      console.log(dataJson);
      if (dataJson.event === "getRooms") {
        this.rooms = JSON.parse(dataJson.content);
      }
    };
    this.ws.onclose = e => {};
  },
  methods: {
    join() {
      if (!this.username) {
        console.log("no valid username");
      }
      if (!this.roomSelected) {
        console.log("no valid game selected");
      } else {
        this.ws.send(
          JSON.stringify({
            event: "registration",
            content:
              '{"username":"' +
              this.username +
              '","room":"' +
              this.roomSelected +
              '"}',
          })
        );
      }
    },
  },
});
