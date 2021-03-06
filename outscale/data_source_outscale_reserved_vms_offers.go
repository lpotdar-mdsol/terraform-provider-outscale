package outscale

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-outscale/osc/fcu"
	"github.com/terraform-providers/terraform-provider-outscale/utils"
)

func dataSourceOutscaleOAPIReservedVMOffers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOutscaleOAPIReservedVMOffersRead,

		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"sub_region_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenancy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"offering_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reserved_vms_offering_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"reserved_vms_offering": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"currency_code": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"marketplace": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"recurring_charge": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"frequency": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"request_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOutscaleOAPIReservedVMOffersRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).FCU

	az, azok := d.GetOk("sub_region_name")
	it, itok := d.GetOk("tenancy")
	ity, ityok := d.GetOk("type")
	pd, pdok := d.GetOk("product_type")
	ot, otok := d.GetOk("offering_type")
	ri, riok := d.GetOk("reserved_vms_offering_id")
	filter, filterOk := d.GetOk("filter")

	req := &fcu.DescribeReservedInstancesOfferingsInput{}

	if azok {
		req.AvailabilityZone = aws.String(az.(string))
	}
	if otok {
		req.OfferingType = aws.String(ot.(string))
	}
	if itok {
		req.InstanceTenancy = aws.String(it.(string))
	}
	if ityok {
		req.InstanceTenancy = aws.String(ity.(string))
	}
	if pdok {
		req.InstanceTenancy = aws.String(pd.(string))
	}
	if riok {
		var ids []*string
		for _, v := range ri.([]interface{}) {
			ids = append(ids, aws.String(v.(string)))
		}
		req.ReservedInstancesOfferingIds = ids
	}
	if filterOk {
		req.Filters = buildOutscaleDataSourceFilters(filter.(*schema.Set))
	}

	var resp *fcu.DescribeReservedInstancesOfferingsOutput
	err := resource.Retry(60*time.Second, func() *resource.RetryError {
		var err error
		resp, err = conn.VM.DescribeReservedInstancesOfferings(req)
		if err != nil {
			if strings.Contains(err.Error(), "RequestLimitExceeded") {
				return resource.RetryableError(err)
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	if resp == nil || len(resp.ReservedInstancesOfferingsSet) == 0 {
		return fmt.Errorf("no matching reserved VMS Offer found")
	}

	utils.PrintToJSON(resp, "OFFERS")

	d.SetId(resource.UniqueId())

	rsi := make([]map[string]interface{}, len(resp.ReservedInstancesOfferingsSet))

	for k, v := range resp.ReservedInstancesOfferingsSet {
		r := make(map[string]interface{})
		r["sub_region_name"] = *v.AvailabilityZone
		r["currency_code"] = *v.CurrencyCode
		r["tenancy"] = *v.InstanceTenancy
		r["type"] = *v.InstanceType
		r["marketplace"] = *v.Martketplace
		r["offering_type"] = *v.OfferingType
		r["product_type"] = *v.ProductDescription
		r["reserved_vms_offering_id"] = *v.ReservedInstancesOfferingId

		rcs := make([]map[string]interface{}, len(v.RecurringCharges))
		for k1, v1 := range v.RecurringCharges {
			rc := make(map[string]interface{})
			rc["frequency"] = v1.Frequency
			rcs[k1] = rc
		}

		r["recurring_charge"] = rcs

		pds := make([]map[string]interface{}, len(v.PricingDetailsSet))
		for k1, v1 := range v.PricingDetailsSet {
			rc := make(map[string]interface{})
			rc["count"] = v1.Count
			rcs[k1] = rc
		}

		r["pricing_details_set"] = pds

		rsi[k] = r
	}

	d.Set("reserved_vms_offering", rsi)
	d.Set("request_id", resp.RequestId)

	return nil
}
