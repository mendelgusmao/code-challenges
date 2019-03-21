#! /bin/bash

DBNAME="supereasy"

for collection in addresses partners users; do
  mongo "${DBNAME}" --eval "db.${collection}.drop()"
  mongoimport --db "${DBNAME}" \
              --collection "${collection}" \
              --file "${collection}.json" \
              --jsonArray
done

mongo "${DBNAME}" --eval 'db.partners.find().forEach(function(item) {
  db.partners.update(
    { _id: item._id },
    { $set: { "location.geo": { type: "Point", coordinates: [ item.location.lat, item.location.long ] } } }
  )
});

db.partners.createIndex({ "location.geo": "2dsphere" });
'
