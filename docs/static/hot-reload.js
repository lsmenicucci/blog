const ws = new WebSocket(`ws://${location.host}/ws`);
ws.onopen = e => {
    console.log("Mounted hot reloading system")
}
ws.onmessage = e => {
    if (e.data === "reload") {
        console.log("Reloading")
        location.reload();
    }
};
