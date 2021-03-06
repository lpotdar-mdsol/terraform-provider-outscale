package outscale

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccOutscaleOAPIDSLinAPIAccesses_basic(t *testing.T) {
	t.Skip()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			skipIfNoOAPI(t)
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOutscaleOAPIVpcEndpointsDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleOAPIVpcEndpointsDataSourceID("data.outscale_net_api_accesses.test"),
					resource.TestCheckResourceAttr("data.outscale_net_api_accesses.test", "vpc_endpoint_set.0.service_name", "com.outscale.eu-west-2.osu"),
				),
			},
		},
	})
}

func testAccCheckOutscaleOAPIVpcEndpointsDataSourceID(n string) resource.TestCheckFunc {
	// Wait for IAM role
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find OMI data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Key Pair data source ID not set")
		}
		return nil
	}
}

const testAccCheckOutscaleOAPIVpcEndpointsDataSourceConfig = `
	resource "outscale_net" "foo" {
		ip_ranges = "10.1.0.0/16"
	}

	resource "outscale_route_table" "foo" {
		net_id = "${outscale_net.foo.id}"
	}

	resource "outscale_net_api_access" "link" {
		net_id = "${outscale_net.foo.id}"

		route_table_id = [
			"${outscale_route_table.foo.id}",
		]

		prefix_list_name = "com.outscale.eu-west-2.osu"
	}

	data "outscale_net_api_accesses" "test" {
		filter {
			name   = "service-name"
			values = ["${outscale_net_api_access.link.service_name}"]
		}
	}
`
