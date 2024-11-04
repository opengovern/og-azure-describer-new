package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"strings"

	"github.com/opengovern/og-azure-describer-new/provider/model"
)

func HdInsightCluster(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armhdinsight.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewClustersClient()

	monitorClientFactory, err := armmonitor.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	diagnosticClient := monitorClientFactory.NewDiagnosticSettingsClient()

	pager := client.NewListPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, cluster := range page.Value {
			resource, err := getHdInsightCluster(ctx, diagnosticClient, cluster)
			if err != nil {
				return nil, err
			}
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func getHdInsightCluster(ctx context.Context, diagnosticClient *armmonitor.DiagnosticSettingsClient, cluster *armhdinsight.Cluster) (*Resource, error) {
	resourceGroup := strings.Split(*cluster.ID, "/")[4]

	var hdinsightListOp []*armmonitor.DiagnosticSettingsResource
	pager := diagnosticClient.NewListPager(*cluster.ID, nil)
	if pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		hdinsightListOp = append(hdinsightListOp, page.Value...)
	}

	resource := Resource{
		ID:       *cluster.ID,
		Name:     *cluster.Name,
		Location: *cluster.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.HdinsightClusterDescription{
				Cluster:                     *cluster,
				DiagnosticSettingsResources: hdinsightListOp,
				ResourceGroup:               resourceGroup,
			},
		},
	}
	return &resource, nil
}