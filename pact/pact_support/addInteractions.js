const fs = require('fs')
const path = require('path')
module.exports = provider => {
  const interactionsStore = path.resolve(__dirname, './interactions.json')
  const interactions = JSON.parse(fs.readFileSync(interactionsStore))
  interactions.forEach(interaction => provider.addInteraction(interaction))
}