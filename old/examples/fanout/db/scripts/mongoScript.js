db = new Mongo().getDB("fanout")
for (let i = 1; i <= 3; i++){
  db.createCollection(`fanout${i}`, { capped: false })
  let entries = []
  for (let i = 0; i <= 50; i++){
    entries.push({key: i, value: Math.floor(Math.random() * 50)})
  }
    db[`fanout${i}`].insertMany(entries)
}


