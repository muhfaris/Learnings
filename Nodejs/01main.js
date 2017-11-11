var http = require("http");

/* Send HTTP Header
 * status code http 200 : ok
 * 
 */
http.createServer(function (request, response) {
  response.writeHead(200, {'Content-Type': 'text/plain'});
  response.end('Hello child nodejs');
}).listen(8081);

console.log('Server running at http://127.0.0.1:8081');
