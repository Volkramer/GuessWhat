new Vue({
    el: "#app",
    data: {
        ws: null,
        roomSelected: "",
        rooms: [],
        room: "",
        username: "",
        inGame: false,
        dialog: true,
        isRegistered: false
    },
    mounted() {
        this.ws = new WebSocket("ws://" + window.location.host + "/ws");
        this.ws.onopen = () => {
            console.log("Connected");
        };
        this.ws.onmessage = e => {
            console.log(e.data);
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
                console.log(
                    JSON.stringify({
                        event: "registration",
                        content:
                            '{"username":"' +
                            this.username +
                            '","room":"' +
                            this.roomSelected +
                            '"}'
                    })
                );

                this.ws.send(
                    JSON.stringify({
                        event: "registration",
                        content:
                            '{"username":"' +
                            this.username +
                            '","room":"' +
                            this.roomSelected +
                            '"}'
                    })
                );
            }
        }
    }
});
