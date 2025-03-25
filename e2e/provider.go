package e2e

import (
	"github.com/Jatinkoli15/terraform-provider-jello/e2e/privateCluster"

	"github.com/Jatinkoli15/terraform-provider-jello/e2e/notebook"

	"github.com/Jatinkoli15/terraform-provider-jello/e2e/modelRepo"

	"github.com/Jatinkoli15/terraform-provider-jello/e2e/modelEndpoint"

	"github.com/Jatinkoli15/terraform-provider-jello/e2e/integration"

	"github.com/Jatinkoli15/terraform-provider-jello/e2e/dataset"

	"github.com/Jatinkoli15/terraform-provider-jello/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider function defines the schema for authentication.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "https://api-loki.e2enetworks.net/myaccount/api/v1/gpu",
				Description: "The API endpoint for requests",
				//https://api-loki.e2enetworks.net/myaccount/api/v1/gpu/teams/{team_id}/projects/{project_id}/notebooks/?api_key={}&active_iam={}
			},
			"auth_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Authentication token",
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "API Key for authentication",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"e2e_notebook":        notebook.ResourceNode(),
			"e2e_eos":             dataset.ResourceEOS(),
			"e2e_modelRepository": modelRepo.ResourceModelRepo(),
			"e2e_modelEndpoint":   modelEndpoint.ResourceModel(),
			"e2e_integration":     integration.ResourceModelRepo(),
			"e2e_privateCluster":  privateCluster.ResourcePrivateCluster(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"e2e_notebook":       notebook.DataSourceImages(),
			"e2e_notebook_plans": notebook.DataSourceSKUPlans(),
		},
		ConfigureFunc: providerConfigure, // setup the API Client
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	api_key := d.Get("api_key").(string)
	auth_token := d.Get("auth_token").(string)
	api_endpoint := d.Get("api_endpoint").(string)
	return client.NewClient(api_key, auth_token, api_endpoint), nil
}
