// This file is generated by GoAdmin CLI adm.
package tables

import (
	"bee-keeper/tables/bzz"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

// Generators is a map of table models.
//
// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "transfers" => http://localhost:9033/admin/info/transfers
//
// example end
//
var Generators = map[string]table.Generator{
	"node": bzz.GetBeeNodeTable,
	// generators end
}
