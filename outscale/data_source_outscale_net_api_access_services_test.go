package outscale

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccOutscaleOAPIVpcEndpointServicesDataSource_Instance(t *testing.T) {
	t.Skip()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			skipIfNoOAPI(t)
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOutscaleOAPIVpcEndpointServicesDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPIVpcEndpointServicesDataSourceID("data.outscale_net_api_access_services.test"),
					resource.TestCheckResourceAttrSet("data.outscale_net_api_access_services.test", "request_id"),
				),
			},
		},
	})
}

func testAccCheckOutscaleOAPIVpcEndpointServicesDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find OMI data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Key Pair data source ID not set")
		}

		if i, err := strconv.Atoi(rs.Primary.Attributes["prefix_list_name.#"]); err != nil || i < 1 {
			return fmt.Errorf("prefix_list_name attribute doesn't have elements")
		}

		return nil
	}
}

const testAccCheckOutscaleOAPIVpcEndpointServicesDataSourceConfig = `
data "outscale_net_api_access_services" "test" {}
`
