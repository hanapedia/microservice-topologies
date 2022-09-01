db = new Mongo().getDB("chain")
for (let i = 1; i <= 9; i++){
  db.createCollection(`chain${i}`, { capped: false })
}

