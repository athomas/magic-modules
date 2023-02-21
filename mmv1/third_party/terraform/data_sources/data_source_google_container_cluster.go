package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGoogleContainerCluster() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := DatasourceSchemaFromResourceSchema(ResourceContainerCluster().Schema)

	// Set 'Required' schema elements
	AddRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	AddOptionalFieldsToSchema(dsSchema, "project", "location")

	return &schema.Resource{
		Read:   datasourceContainerClusterRead,
		Schema: dsSchema,
	}
}

func datasourceContainerClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	clusterName := d.Get("name").(string)

	location, err := GetLocation(d, config)
	if err != nil {
		return err
	}

	project, err := GetProject(d, config)
	if err != nil {
		return err
	}

	id := containerClusterFullName(project, location, clusterName)

	d.SetId(id)

	if err := resourceContainerClusterRead(d, meta); err != nil {
		return err
	}

	if d.Id() == "" {
		return fmt.Errorf("%s not found", id)
	}

	return nil
}