package outscale

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceOutscaleOAPIPrefixLists(t *testing.T) {
	t.Skip()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			skipIfNoOAPI(t)
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceOutscaleOAPIPrefixListsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.outscale_prefix_lists.s3_by_id", "prefix_list_set.#", "1"),
				),
			},
		},
	})
}

const testAccDataSourceOutscaleOAPIPrefixListsConfig = `
	data "outscale_prefix_lists" "s3_by_id" {
		prefix_list_id = ["pl-a14a8cdc"]
	}
`
