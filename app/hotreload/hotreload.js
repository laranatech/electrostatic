const wsPath = "ws://localhost:3030/ws"

const ws = new WebSocket(wsPath)

ws.onmessage = (e) => {
	if (e.data === "reload") {
		console.log("reload")
		window.location.reload()
	}
}

