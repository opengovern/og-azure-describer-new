select title, akas
from azure.azure_storage_blob
where resource_group = '{{ resourceName }}' and storage_account_name = '{{ resourceName }}' and region = '{{ output.location.value }}' and name = '{{ resourceName }}';