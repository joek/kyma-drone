const axios = require("axios");

// SLACK_URL https://hooks.slack.com/services/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

module.exports = { main: async function (event, context) {
    await axios.post(process.env.SLACK_URL, {text: "Bye bye!"});
    
} }