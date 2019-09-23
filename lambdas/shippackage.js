const axios = require('axios');
 
const traceHeaders = ['x-request-id', 'x-b3-traceid', 'x-b3-spanid', 'x-b3-parentspanid', 'x-b3-sampled', 'x-b3-Flags', 'x-ot-span-context']


module.exports = { main: async function (event, context) { 
    console.log(`Event Data: ${event.data}`);
    var orderCode = event.data.orderCode;
    var traceCtxHeaders = extractTraceHeaders(event.extensions.request.headers);
    
    var orderUrl = `${process.env['order-GATEWAY_URL']}/Orders`;


    var orderBody = {
        code: orderCode,
        versionID: null,
        status: { code : 'DELIVERY' }
    }
    try {
        await axios.post(orderUrl, orderBody, {
            headers: traceCtxHeaders,
            responseType: 'json'
        })
    } catch(error) {
            console.error(error)
            console.log(orderUrl)
            return
    }


    var shippingUrl = `${process.env['drone-GATEWAY_URL']}/shipPackage`;
    
    shippingBody = {
        orderCode: orderCode
    }
    
    try {
        await axios.post(shippingUrl, shippingBody, {
            headers: traceCtxHeaders,
            responseType: 'json'
        })
    } catch(error) {
        console.error(error)
        console.log(shippingUrl)
        return
    }
    

    
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
