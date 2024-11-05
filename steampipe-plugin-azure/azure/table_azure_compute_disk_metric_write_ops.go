package azure

import (
	"context"

	"github.com/opengovern/og-describer-azure/SDK/generated"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureComputeDiskMetricWriteOps(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_disk_metric_write_ops",
		Description: "Azure Compute Disk Metrics - Write Ops",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeDiskWriteOps,
		},
		Columns: kaytuMonitoringMetricColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the disk.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.MonitoringMetric.DimensionValue").Transform(lastPathElement),
			},
		}),
	}
}
