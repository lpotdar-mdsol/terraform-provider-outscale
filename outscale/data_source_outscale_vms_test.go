package outscale

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccOutscaleOAPIVMSDataSource_basic(t *testing.T) {
	region := os.Getenv("OUTSCALE_REGION")
	omi := getOMIByRegion(region, "ubuntu").OMI

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			skipIfNoOAPI(t)
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOAPIVMSDataSourceConfig(omi, "t2.micro"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.outscale_vms.basic_web", "vms.0.image_id", omi),
					resource.TestCheckResourceAttr(
						"data.outscale_vms.basic_web", "vms.0.vm_type", "t2.micro"),
				),
			},
		},
	})
}

func testAccOAPIVMSDataSourceConfig(omi, vmType string) string {
	return fmt.Sprintf(`
		resource "outscale_vm" "basic" {
			image_id			= "%s"
			vm_type				= "%s"
			keypair_name	= "terraform-basic"
		}

		data "outscale_vms" "basic_web" {
			filter {
				name   = "vm_ids"
				values = ["${outscale_vm.basic.id}"]
			}
		}`, omi, vmType)
}
