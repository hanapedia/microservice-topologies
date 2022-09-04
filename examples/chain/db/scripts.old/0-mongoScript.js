db = new Mongo().getDB("chain")
for (let i = 1; i <= 9; i++){
  db.createCollection(`chain${i}`, { capped: false })
  let entries = []
  for (let i = 0; i <= 50; i++){
    entries.push({key: i, value: Math.floor(Math.random() * 50)})
  }
    db[`chain${i}`].insertMany(entries)
}

