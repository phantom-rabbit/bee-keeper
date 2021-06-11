package bzz

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"time"
)

func GetBeeNodeTable(ctx *context.Context) table.Table {

	beeNode := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := beeNode.GetInfo()

	//info.AddButton("支票查询", icon.Save, action.PopUp())

	info.AddField("Id", "id", db.Int8)
	info.AddField("更新时间", "updated_at", db.Datetime)
	info.AddField("Ip", "ip", db.Text)
	info.AddField("Port", "port", db.Int8)
	info.AddField("Owner", "owner", db.Text)
	info.AddField("Contract", "contract", db.Text)

	info.SetTable("bee_node").SetTitle("节点列表").SetDescription("BeeNode")

	formList := beeNode.GetForm()
	formList.AddField("created_time", "created_at", db.Datetime, form.Datetime).FieldDefault(time.Now().Format("2006-01-02 15:04:05")).FieldHide()
	formList.AddField("updated_at", "updated_at", db.Datetime, form.Datetime).FieldDefault(time.Now().Format("2006-01-02 15:04:05")).FieldHide()
	formList.AddField("Ip", "ip", db.Text, form.Ip)
	formList.AddField("Port", "port", db.Int8, form.Text)
	formList.AddField("Owner", "owner", db.Text, form.RichText)
	formList.AddField("Contract", "contract", db.Text, form.RichText)

	formList.SetTable("bee_node").SetTitle("新增节点").SetDescription("BeeNode")

	return beeNode
}
