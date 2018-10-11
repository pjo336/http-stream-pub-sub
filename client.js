const net = require('net');

const client = new net.Socket({readable: true, writable: true});

client.connect(8000, '127.0.0.1');

client.on('connect', async function() {
  console.log('Connected');
  client.setNoDelay(true);
  const isRegistered = await client.write("GET /register\n");
  if(isRegistered === false) {
  	client.destroy({ error: "Could not write register command" });
	}
});

client.on('data', function(data) {
	console.log('Received: ' + data);
});

client.on('close', function() {
	console.log('Connection closed');
});
