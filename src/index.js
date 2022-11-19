const ws = new WebSocket ("ws://localhost:3000/ws");

ws.onopen = (e)=>{
		console.log("conectado");
		ws.send("ok");
};

window.addEventListener('beforeunload', (e)=>{
	e.preventDefault();
	ws.send("exit");
});

ws.onmessage = (e)=>{
	console.log(e.data);
	let data = JSON.parse( e.data );

	if ( data.type == "bind" ){
		isBind( data );
		return;
	}
	if ( data.type == "eval" ){
		isEval( data );
		return; 
	}
};

function isBind( data ){

	window[data.name] = ()=>{
		if ( event.type != "dragstart"){
			event.preventDefault();
		}
		console.log(event.target.value);
		ws.send( JSON.stringify({type:"event", name:data.name ,event:{type:event.type,id:event.target.id,value:event.target.value}}) );
	};
};

function isEval( data ){

	let res = eval( data.js );
		if ( typeof res != "string" ){
			res = JSON.stringify( res );
		}
		if ( data.id ){
			res = JSON.stringify( {id : data.id , body: res} );
		}
		if ( res != undefined ){
			ws.send( JSON.stringify({type:"response",name:data.name,event:{value:res}}) );
		}
};


