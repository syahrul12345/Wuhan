const passwords = require('./password')
const axios = require('axios')
const TelegramBot = require('node-telegram-bot-api');
const bot = new TelegramBot(passwords.BOTAPIKEY, {polling: true});
// Matches "/echo [whatever]"
bot.onText(/\/update global (\d+$)/, (msg, match) => {
    const chatId = msg.chat.id;
    const resp = match[1]; // the captured "whatever"
    const payload = {
        "country":"global",
        "deaths":parseInt(resp)
    }
    axios.post(passwords.URL_END_POINT,payload)
        .then((res) => {
            bot.sendMessage(chatId,res.data.message)
        })
        .catch((err) => {
            bot.sendMessage("Failed to update the total number of deaths")
        })
})
console.log("Running bot")