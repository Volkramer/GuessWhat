new Vue({
  el: "#app",
  data: {
    ws: null,
    roomSelected: "",
    rooms: [],
    room: "",
    username: "",
    isRegistered: false,
    users: [],
    user: "",
    messages: [],
    message: "",
    newMessage: "",
  },
  mounted() {
    this.ws = new WebSocket("ws://" + window.location.host + "/ws");
    this.ws.onopen = () => {};
    this.ws.onmessage = e => {
      var dataJson = JSON.parse(e.data);
      console.log(dataJson);
      content = dataJson.content;
      switch (dataJson.event) {
        case "getRooms":
          this.rooms = JSON.parse(content);
          break;
        case "newUser":
          this.room = content.room;
          this.username = content.username;
          this.isRegistered = true;
          break;
        case "userJoin":
          this.users.push(content);
          break;
        case "message":
          console.log(content);
          this.messages.push(content);
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

    Leave() {},
    comment() {
      if (this.newMessage !== "") {
        this.ws.send(
          JSON.stringify({
            event: "message",
            content:
              '{"username":"' +
              this.username +
              '","text":"' +
              this.newMessage +
              '"}',
          })
        );
      }
    },
  },
});
