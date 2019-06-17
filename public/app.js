new Vue({
  el: "#app",
  data: {
    ws: null,
    roomSelected: "",
    rooms: [],
    room: "",
    username: "",
    isRegistered: false,
  },
  mounted() {
    this.ws = new WebSocket("ws://" + window.location.host + "/ws");
    this.ws.onopen = () => {};
    this.ws.onmessage = e => {
      var dataJson = JSON.parse(e.data);
      console.log(dataJson);
      if (typeof dataJson.content === "object") {
        content = dataJson.content;
      } else {
        content = JSON.parse(dataJson.content);
      }
      console.log(content);

      switch (dataJson.event) {
        case "getRooms":
          this.rooms = content;
          break;
        case "newUser":
          this.room = content.room;
          this.username = content.username;
          this.isRegistered = true;

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
      if (this.username === "") {
        console.log("no valid username");
      }
      if (this.roomSelected === "") {
        console.log("no valid room selected");
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
