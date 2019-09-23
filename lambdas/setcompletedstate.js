const axios = require('axios');

const traceHeaders = ['x-request-id', 'x-b3-traceid', 'x-b3-spanid', 'x-b3-parentspanid', 'x-b3-sampled', 'x-b3-Flags', 'x-ot-span-context']


module.exports = { main: async function (event, context) { 
        console.log(`Event Data: ${event.data}`);
    var orderCode = event.data.orderCode;
    var traceCtxHeaders = extractTraceHeaders(event.extensions.request.headers);

    var url = `${process.env['order-GATEWAY_URL']}/Orders`;

    var body = {
        code: orderCode,
        versionID: null,
        status: { code : 'COMPLETED' }
    }

    try {
    resp = await axios.post(url, body, {
        headers: traceCtxHeaders,
        responseType: 'json'
    })
    } catch(error) {
        console.error(error)
        console.log(url)
    }
    
    return resp.body
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
