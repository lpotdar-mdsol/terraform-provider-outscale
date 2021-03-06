package outscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider ...
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OUTSCALE_ACCESSKEYID", nil),
				Description: "The Access Key ID for API operations.",
			},
			"secret_key_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OUTSCALE_SECRETKEYID", nil),
				Description: "The Secret Key ID for API operations.",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OUTSCALE_REGION", nil),
				Description: "The Region for API operations.",
			},
			"oapi": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OUTSCALE_OAPI", false),
				Description: "Enable oAPI Usage",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"outscale_vm":                            resourceOutscaleOApiVM(),
			"outscale_keypair":                       resourceOutscaleOAPIKeyPair(),
			"outscale_image":                         resourceOutscaleOAPIImage(),
			"outscale_internet_service_link":         resourceOutscaleOAPIInternetServiceLink(),
			"outscale_internet_service":              resourceOutscaleOAPIInternetService(),
			"outscale_net":                           resourceOutscaleOAPINet(),
			"outscale_security_group":                resourceOutscaleOAPISecurityGroup(),
			"outscale_outbound_rule":                 resourceOutscaleOAPIOutboundRule(),
			"outscale_security_group_rule":           resourceOutscaleOAPIOutboundRule(),
			"outscale_tag":                           resourceOutscaleOAPITags(),
			"outscale_public_ip":                     resourceOutscaleOAPIPublicIP(),
			"outscale_public_ip_link":                resourceOutscaleOAPIPublicIPLink(),
			"outscale_volume":                        resourceOutscaleOAPIVolume(),
			"outscale_volumes_link":                  resourceOutscaleOAPIVolumeLink(),
			"outscale_net_attributes":                resourceOutscaleOAPILinAttributes(),
			"outscale_nat_service":                   resourceOutscaleOAPINatService(),
			"outscale_subnet":                        resourceOutscaleOAPISubNet(),
			"outscale_route":                         resourceOutscaleOAPIRoute(),
			"outscale_route_table":                   resourceOutscaleOAPIRouteTable(),
			"outscale_route_table_link":              resourceOutscaleOAPILinkRouteTable(),
			"outscale_image_copy":                    resourceOutscaleOAPIImageCopy(),
			"outscale_vpn_connection":                resourceOutscaleOAPIVpnConnection(),
			"outscale_vpn_gateway":                   resourceOutscaleOAPIVpnGateway(),
			"outscale_image_tasks":                   resourceOutscaleOAPIVpnGateway(),
			"outscale_vpn_connection_route":          resourceOutscaleOAPIVpnConnectionRoute(),
			"outscale_vpn_gateway_route_propagation": resourceOutscaleOAPIVpnGatewayRoutePropagation(),
			"outscale_vpn_gateway_link":              resourceOutscaleOAPIVpnGatewayLink(),
			"outscale_nic":                           resourceOutscaleOAPINic(),
			"outscale_snapshot_export_tasks":         resourceOutscaleOAPIImageExportTasks(),
			"outscale_snapshot":                      resourceOutscaleOAPISnapshot(),
			"outscale_image_register":                resourceOutscaleOAPIImageRegister(),
			"outscale_keypair_importation":           resourceOutscaleOAPIKeyPairImportation(),
			"outscale_image_launch_permission":       resourceOutscaleOAPIImageLaunchPermission(),
			"outscale_net_peering":                   resourceOutscaleOAPILinPeeringConnection(),
			"outscale_net_peering_acceptation":       resourceOutscaleOAPILinPeeringConnectionAccepter(),
			"outscale_nic_link":                      resourceOutscaleOAPINetworkInterfaceAttachment(),
			"outscale_nic_private_ip":                resourceOutscaleOAPINetworkInterfacePrivateIP(),
			"outscale_snapshot_attributes":           resourcedOutscaleOAPISnapshotAttributes(),
			"outscale_net_api_access":                resourceOutscaleOAPIVpcEndpoint(),
			"outscale_snapshot_import":               resourcedOutscaleOAPISnapshotImport(),
			"outscale_client_endpoint":               resourceOutscaleOAPICustomerGateway(),
			"outscale_dhcp_option":                   resourceOutscaleDHCPOption(),     //TODO: OAPI
			"outscale_dhcp_option_link":              resourceOutscaleDHCPOptionLink(), //TODO: OAPI
		},
		DataSourcesMap: map[string]*schema.Resource{
			"outscale_vm":                      dataSourceOutscaleOAPIVM(),
			"outscale_vms":                     datasourceOutscaleOApiVMS(),
			"outscale_security_group":          dataSourceOutscaleOAPISecurityGroup(),
			"outscale_security_groups":         dataSourceOutscaleOAPISecurityGroups(),
			"outscale_image":                   dataSourceOutscaleOAPIImage(),
			"outscale_images":                  dataSourceOutscaleOAPIImages(),
			"outscale_tag":                     dataSourceOutscaleOAPITag(),
			"outscale_tags":                    dataSourceOutscaleOAPITags(),
			"outscale_public_ip":               dataSourceOutscaleOAPIPublicIP(),
			"outscale_public_ips":              dataSourceOutscaleOAPIPublicIPS(),
			"outscale_volume":                  datasourceOutscaleOAPIVolume(),
			"outscale_volumes":                 datasourceOutscaleOAPIVolumes(),
			"outscale_nat_service":             dataSourceOutscaleOAPINatService(),
			"outscale_nat_services":            dataSourceOutscaleOAPINatServices(),
			"outscale_keypair":                 datasourceOutscaleOAPIKeyPair(),
			"outscale_keypairs":                datasourceOutscaleOAPIKeyPairs(),
			"outscale_vm_state":                dataSourceOutscaleOAPIVMState(),
			"outscale_vms_state":               dataSourceOutscaleOAPIVMSState(),
			"outscale_internet_service":        datasourceOutscaleOAPIInternetService(),
			"outscale_internet_services":       datasourceOutscaleOAPIInternetServices(),
			"outscale_subnet":                  dataSourceOutscaleOAPISubnet(),
			"outscale_subnets":                 dataSourceOutscaleOAPISubnets(),
			"outscale_net":                     dataSourceOutscaleOAPIVpc(),
			"outscale_nets":                    dataSourceOutscaleOAPIVpcs(),
			"outscale_net_attributes":          dataSourceOutscaleOAPIVpcAttr(),
			"outscale_client_endpoint":         dataSourceOutscaleOAPICustomerGateway(),
			"outscale_client_endpoints":        dataSourceOutscaleOAPICustomerGateways(),
			"outscale_route_table":             dataSourceOutscaleOAPIRouteTable(),
			"outscale_route_tables":            dataSourceOutscaleOAPIRouteTables(),
			"outscale_vpn_gateway":             dataSourceOutscaleOAPIVpnGateway(),
			"outscale_vpn_gateways":            dataSourceOutscaleOAPIVpnGateways(),
			"outscale_vpn_connection":          dataSourceOutscaleVpnConnection(), //TODO: OAPI
			"outscale_sub_region":              dataSourceOutscaleOAPIAvailabilityZone(),
			"outscale_prefix_list":             dataSourceOutscaleOAPIPrefixList(),
			"outscale_quota":                   dataSourceOutscaleOAPIQuota(),
			"outscale_quotas":                  dataSourceOutscaleOAPIQuotas(),
			"outscale_prefix_lists":            dataSourceOutscaleOAPIPrefixLists(),
			"outscale_region":                  dataSourceOutscaleOAPIRegion(),
			"outscale_sub_regions":             dataSourceOutscaleOAPIAvailabilityZone(),
			"outscale_regions":                 dataSourceOutscaleOAPIRegions(),
			"outscale_vpn_connections":         dataSourceOutscaleOAPIVpnConnections(),
			"outscale_product_types":           dataSourceOutscaleOAPIProductTypes(),
			"outscale_reserved_vms":            dataSourceOutscaleOAPIReservedVMS(),
			"outscale_vm_type":                 dataSourceOutscaleOAPIVMType(),
			"outscale_vm_types":                dataSourceOutscaleOAPIVMTypes(),
			"outscale_reserved_vms_offer":      dataSourceOutscaleOAPIReservedVMOffer(),
			"outscale_reserved_vms_offers":     dataSourceOutscaleOAPIReservedVMOffers(),
			"outscale_snapshot":                dataSourceOutscaleOAPISnapshot(),
			"outscale_snapshots":               dataSourceOutscaleOAPISnapshots(),
			"outscale_net_peering":             dataSourceOutscaleOAPILinPeeringConnection(),
			"outscale_net_peerings":            dataSourceOutscaleOAPILinPeeringsConnection(),
			"outscale_nics":                    dataSourceOutscaleOAPINics(),
			"outscale_nic":                     dataSourceOutscaleOAPINic(),
			"outscale_net_api_access":          dataSourceOutscaleOAPIVpcEndpoint(),
			"outscale_net_api_accesses":        dataSourceOutscaleOAPIVpcEndpoints(),
			"outscale_net_api_access_services": dataSourceOutscaleOAPIVpcEndpointServices(),
		},

		ConfigureFunc: providerConfigureClient,
	}
}

func providerConfigureClient(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessKeyID: d.Get("access_key_id").(string),
		SecretKeyID: d.Get("secret_key_id").(string),
		Region:      d.Get("region").(string),
		OApi:        d.Get("oapi").(bool),
	}
	return config.Client()
}
