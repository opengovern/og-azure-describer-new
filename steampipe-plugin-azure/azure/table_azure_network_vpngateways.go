package azure

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-azure/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureNetworkVPNGateways(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_network_vpngateways",
		Description: "Azure Network VPNGateways",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetVpnGateway,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListVpnGateway,
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the vpngateways.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VPNGateways.ID")},
			{
				Name:        "name",
				Description: "The name of the vpngateways.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpnGateway.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpnGateway.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.VpnGateway.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.VpnGateway.ID").Transform(idToAkas),
			},
		}),
	}
}
