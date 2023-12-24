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
			"id": "o2re3duol3dwz5c",
			"created": "2023-12-24 12:39:50.134Z",
			"updated": "2023-12-24 12:39:50.134Z",
			"name": "transactions",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "zq9kvavz",
					"name": "paid_by",
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
				},
				{
					"system": false,
					"id": "tsdzhmqq",
					"name": "amount",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "myadlz39",
					"name": "proportions",
					"type": "json",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSize": 2000000
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_tLnxTua` + "`" + ` ON ` + "`" + `transactions` + "`" + ` (` + "`" + `paid_by` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("o2re3duol3dwz5c")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
