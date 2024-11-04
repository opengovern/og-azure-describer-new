package azure

import (
	"context"
	"github.com/opengovern/og-azure-describer-new/SDK/generated"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureNetworkFirewallPolicies(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_network_firewallpolicies",
		Description: "Azure Network FirewallPolicies",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetFirewallPolicy,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListFirewallPolicy,
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the firewallpolicies.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicies.ID")},
			{
				Name:        "name",
				Description: "The name of the firewallpolicies.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicy.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FirewallPolicy.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.FirewallPolicy.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.FirewallPolicy.ID").Transform(idToAkas),
			},
		}),
	}
}