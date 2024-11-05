package azure

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-azure/pkg/SDK/generated"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureMariaDBDatabases(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_mariadb_databases",
		Description: "Azure MariaDB Databases",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetMariadbDatabase,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListMariadbDatabase,
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the databases.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Databases.ID")},
			{
				Name:        "name",
				Description: "The name of the databases.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.Database.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.Database.ID").Transform(idToAkas),
			},
		}),
	}
}
