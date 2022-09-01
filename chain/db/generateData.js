'use strict'

const fs = require('fs')

for (let i = 1; i <= 9; i++){
  var entries = []
  for (let i = 0; i <= 50; i++){
    var entry = {
      key: i,
      value: Math.floor(Math.random() * 50),
    }
  entries.push(entry)
  }
  fs.writeFileSync(`chain_${i}.json`, JSON.stringify(entries))
}
