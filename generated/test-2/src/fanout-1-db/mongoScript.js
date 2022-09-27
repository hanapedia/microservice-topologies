db = new Mongo().getDB("test-2")
db.createCollection("fanout1", { capped: false })
let entries = []
for (let i = 0; i <= 50; i++){
  entries.push({key: i, value: Math.floor(Math.random() * 50)})
}
db["fanout1"].insertMany(entries)
