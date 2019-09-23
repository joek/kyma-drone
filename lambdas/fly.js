const axios = require('axios');
 
const traceHeaders = ['x-request-id', 'x-b3-traceid', 'x-b3-spanid', 'x-b3-parentspanid', 'x-b3-sampled', 'x-b3-Flags', 'x-ot-span-context']


module.exports = { main: function (event, context) { 
    var traceCtxHeaders = extractTraceHeaders(event.extensions.request.headers);

    var takeOffUrl = `${process.env['drone-GATEWAY_URL']}/takeOff`;
    var landUrl = `${process.env['drone-GATEWAY_URL']}/land`;
    
    b = {
    }
    
    axios.post(takeOffUrl, b, {
        headers: traceCtxHeaders,
        responseType: 'json'
    }).then(function(resp){
        setTimeout(function() {
            axios.post(landUrl, b, {
            headers: traceCtxHeaders,
            responseType: 'json'
        })
        }, 10000);
    })

    

    
    return
} }


function extractTraceHeaders(headers) {
    console.log(headers)
    var map = {};
    for (var i in traceHeaders) {
        h = traceHeaders[i]
        headerVal = headers[h]
        console.log('header' + h + "-" + headerVal)
        if (headerVal !== undefined) {
            console.log('if not undefined header' + h + "-" + headerVal)
            map[h] = headerVal
        }
    }
    return map;
}
