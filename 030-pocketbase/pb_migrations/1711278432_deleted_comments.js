/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("s8t4p8g0wsa2ren");

  return dao.deleteCollection(collection);
}, (db) => {
  const collection = new Collection({
    "id": "s8t4p8g0wsa2ren",
    "created": "2024-03-20 13:24:10.756Z",
    "updated": "2024-03-20 14:08:22.198Z",
    "name": "comments",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "btfdkjzc",
        "name": "message",
        "type": "text",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "xl9ifura",
        "name": "post",
        "type": "relation",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "collectionId": "x1qvi72e55d13uo",
          "cascadeDelete": false,
          "minSelect": null,
          "maxSelect": 1,
          "displayFields": null
        }
      },
      {
        "system": false,
        "id": "0soqn8cl",
        "name": "user",
        "type": "relation",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "collectionId": "_pb_users_auth_",
          "cascadeDelete": false,
          "minSelect": null,
          "maxSelect": 1,
          "displayFields": null
        }
      }
    ],
    "indexes": [],
    "listRule": "@request.auth.verified = true",
    "viewRule": "@request.auth.verified = true",
    "createRule": "@request.auth.verified = true",
    "updateRule": "id = @request.auth.id",
    "deleteRule": "id = @request.auth.id",
    "options": {}
  });

  return Dao(db).saveCollection(collection);
})