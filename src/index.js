const ws = new WebSocket ("ws://localhost:3000/ws");

	ws.onopen = (e)=>{
		console.log("conectado");
		ws.send("ok");
	};

	ws.onmessage = (e)=>{
		let data = e.data;
		console.log(data);
	};
