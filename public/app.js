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
    clients: [],
    strokes: [],
    otherStrokes: [],
    strokeColor: "#000000",
    isDrawing: false,
    mouse: {
      x: 0,
      y: 0
    }
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
        case "stroke":
          //console.log(dataJson);
          this.addPoint(dataJson);
          break;
        default:
          break;
      }
    };
  },
  computed: {
    mousePosition: function() {
      var canvas = document.getElementById("canvas");
      var rect = canvas.getBoundingClientRect();

      return {
        x: this.mouse.x - rect.left,
        y: this.mouse.y - rect.top
      };
    }
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
    },
    onMouseDown: function(event) {
      this.isDrawing = true;
      this.mouse = {
        x: event.pageX,
        y: event.pageY
      };
      this.sendPoint(this.mousePosition.x, this.mousePosition.y, true);
    },
    onMouseUp: function() {
      this.isDrawing = false;
    },
    onMouseMove: function(event) {
      if (this.isDrawing) {
        this.mouse = {
          x: event.pageX,
          y: event.pageY
        };
        this.sendPoint(this.mousePosition.x, this.mousePosition.y, false);
      }
    },
    onMouseLeave: function() {
      this.isDrawing = false;
    },
    sendPoint(x, y, newStroke) {
      var p = { x: x, y: y };
      data = {
        event: "stroke",
        username: this.username,
        points: [p],
        finish: newStroke
      };
      this.ws.send(JSON.stringify(data));
    },
    addPoint(data) {
      var p = { x: data.points[0].x, y: data.points[0].y };
      if (data.finish) {
        this.strokes.push([p]);
      } else {
        this.strokes[this.strokes.length - 1].push(p);
      }
      this.update();
    },
    /* addOtherPoint(data) {
      if (data.finish) {
        this.otherStrokes[data.username].push(data.points[0]);
      } else {
        this.strokes = this.otherStrokes[dataJson.username];
        this.strokes[this.strokes.length - 1] = this.strokes[
          this.strokes.length - 1
        ].concat(dataJson.points);
      }
      this.update();
    }, */
    update() {
      var canvas = document.getElementById("canvas");
      var ctx = canvas.getContext("2d");
      ctx.clearRect(0, 0, ctx.canvas.width, ctx.canvas.height);
      ctx.lineJoin = "round";
      ctx.lineWidth = 4;
      ctx.strokeStyle = this.strokeColor;
      this.drawStrokes(this.strokes);
      /* for (var i = 0; i < this.clients.length; i++) {
        var username = this.clients[i];
        this.drawStrokes(this.otherStrokes[username]);
      } */
    },
    drawStrokes(strokes) {
      var canvas = document.getElementById("canvas");
      var ctx = canvas.getContext("2d");
      for (let i = 0; i < strokes.length; i++) {
        ctx.beginPath();
        for (var j = 1; j < strokes[i].length; j++) {
          var prev = strokes[i][j - 1];
          var current = strokes[i][j];
          ctx.moveTo(prev.x, prev.y);
          ctx.lineTo(current.x, current.y);
        }
        ctx.closePath();
        ctx.stroke();
      }
    }
  }
});
