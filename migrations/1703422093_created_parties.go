package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "azrqika42gy4juj",
			"created": "2023-12-24 12:48:13.832Z",
			"updated": "2023-12-24 12:48:13.832Z",
			"name": "parties",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "gt08qkgq",
					"name": "name",
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
					"id": "o8rtwc67",
					"name": "users",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "u6333ypt",
					"name": "transactions",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "o2re3duol3dwz5c",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_Z4c8na5` + "`" + ` ON ` + "`" + `parties` + "`" + ` (` + "`" + `name` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_JnfgW9B` + "`" + ` ON ` + "`" + `parties` + "`" + ` (` + "`" + `users` + "`" + `)"
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("azrqika42gy4juj")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
