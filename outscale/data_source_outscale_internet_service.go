package outscale

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/antihax/optional"
	oscgo "github.com/marinsalinas/osc-sdk-go"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func datasourceOutscaleOAPIInternetService() *schema.Resource {
	return &schema.Resource{
		Read: datasourceOutscaleOAPIInternetServiceRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"net_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"internet_service_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsOAPIListSchemaComputed(),
		},
	}
}

func datasourceOutscaleOAPIInternetServiceRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	filters, filtersOk := d.GetOk("filter")
	internetID, insternetIDOk := d.GetOk("internet_service_id")

	if filtersOk == false && insternetIDOk == false {
		return fmt.Errorf("One of filters, or instance_id must be assigned")
	}

	// Build up search parameters
	params := oscgo.ReadInternetServicesRequest{}

	if insternetIDOk {
		params.Filters = &oscgo.FiltersInternetService{
			InternetServiceIds: &[]string{internetID.(string)},
		}

	}

	if filtersOk {
		params.Filters = buildOutscaleOSCAPIDataSourceInternetServiceFilters(filters.(*schema.Set))
	}

	var resp oscgo.ReadInternetServicesResponse
	var err error
	err = resource.Retry(120*time.Second, func() *resource.RetryError {
		r, _, err := conn.InternetServiceApi.ReadInternetServices(context.Background(), &oscgo.ReadInternetServicesOpts{ReadInternetServicesRequest: optional.NewInterface(params)})

		if err != nil {
			if strings.Contains(err.Error(), "RequestLimitExceeded:") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}

		resp = r

		return resource.RetryableError(err)
	})

	var errString string

	if err != nil {
		errString = err.Error()

		return fmt.Errorf("[DEBUG] Error reading Internet Service id (%s)", errString)
	}

	if !resp.HasInternetServices() || len(resp.GetInternetServices()) == 0 {
		return fmt.Errorf("Error reading Internet Service: Internet Services is not found with the seatch criteria")
	}

	result := resp.GetInternetServices()[0]

	log.Printf("[DEBUG] Setting OAPI Internet Service id (%s)", err)

	d.Set("request_id", resp.ResponseContext.GetRequestId())
	d.Set("internet_service_id", result.GetInternetServiceId())
	d.Set("state", result.GetState())
	d.Set("net_id", result.GetNetId())

	d.SetId(result.GetInternetServiceId())

	return d.Set("tags", tagsOSCAPIToMap(result.GetTags()))
}

func buildOutscaleOSCAPIDataSourceInternetServiceFilters(set *schema.Set) *oscgo.FiltersInternetService {
	var filters oscgo.FiltersInternetService
	for _, v := range set.List() {
		m := v.(map[string]interface{})
		var filterValues []string
		for _, e := range m["values"].([]interface{}) {
			filterValues = append(filterValues, e.(string))
		}

		switch name := m["name"].(string); name {
		case "internet_service_ids":
			filters.InternetServiceIds = &filterValues
		default:
			log.Printf("[Debug] Unknown Filter Name: %s.", name)
		}
	}
	return &filters
}
