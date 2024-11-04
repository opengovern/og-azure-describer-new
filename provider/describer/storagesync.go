package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
	"strings"

	"github.com/opengovern/og-azure-describer-new/provider/model"
)

func StorageSync(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *StreamSender) ([]Resource, error) {
	clientFactory, err := armstoragesync.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewServicesClient()

	pager := client.NewListBySubscriptionPager(nil)
	var values []Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetStorageSync(ctx, v)
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

func GetStorageSync(ctx context.Context, storage *armstoragesync.Service) *Resource {
	resourceGroup := strings.Split(*storage.ID, "/")[4]

	resource := Resource{
		ID:       *storage.ID,
		Name:     *storage.Name,
		Location: *storage.Location,
		Description: JSONAllFieldsMarshaller{
			Value: model.StorageSyncDescription{
				Service:       *storage,
				ResourceGroup: resourceGroup,
			},
		}}
	return &resource
}
