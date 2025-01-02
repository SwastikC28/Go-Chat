class SocketClient {
}

var socket = new WebSocket("ws://localhost:8080/ws")

let connect = (fn) => {
    console.log("Connecting..");

    socket.onopen = () => {
        console.log("Websocket connected");
    }

    socket.onmessage = (msg) => {
        console.log("Message from socket", msg);
        fn(msg)
    }

    socket.onclose = (ev) => {
        console.log("Websocket connection closed", ev);
    }


    socket.onerror = (err) => {
        console.log("Websocket error", err);
    }
};

let sendMsg = (msg) => {
    console.log("Msg send :- ", msg);
    socket.send(msg)
}


export {
    connect,
    sendMsg
}