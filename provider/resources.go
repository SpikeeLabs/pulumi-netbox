// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netbox

import (
	"unicode"

	"github.com/e-breuninger/terraform-provider-netbox/netbox"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	netboxPkg = "netbox"
	// modules:
	netboxMod = "index" // the netbox module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// netboxMember manufactures a type token for the Scaleway package and the given module and type.
func netboxMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(netboxPkg + ":" + mod + ":" + mem)
}

// netboxType manufactures a type token for the Scaleway package and the given module and type.
func netboxType(mod string, typ string) tokens.Type {
	return tokens.Type(netboxMember(mod, typ))
}

// netboxDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the Scaleway package and names the file by simply lower casing the data
// source's first character.
func netboxDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return netboxMember(mod+"/"+fn, res)
}

// netboxResource manufactures a standard resource token given a module and resource name.
// It automatically uses the Scaleway package and names the file by simply lower casing the resource's
// first character.
func netboxResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return netboxType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(netbox.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "netbox",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Netbox",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Hayden Young",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing Netbox resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "netbox", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/SpikeeLabs/pulumi-netbox",
		Repository: "https://github.com/SpikeeLabs/pulumi-netbox",

		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "e-breuninger",
		Config:    map[string]*tfbridge.SchemaInfo{
            "api_token": {
                Default: &tfbridge.DefaultInfo{
                    EnvVars: []string{"NETBOX_API_TOKEN"},
                },
            },
            "server_url": {
                Default: &tfbridge.DefaultInfo{
                    EnvVars: []string{"NETBOX_SERVER_URL"},
                },
            },

            // TODO: Add the rest: https://registry.terraform.io/providers/e-breuninger/netbox/latest/docs#schema
		},
		PreConfigureCallback: preConfigureCallback,

		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
            "netbox_aggregate": {Tok: netboxResource(netboxMod, "Aggregate")},
            "netbox_available_ip_address": {Tok: netboxResource(netboxMod, "AvailableIpAddress")},
            "netbox_available_prefix": {Tok: netboxResource(netboxMod, "AvailablePrefix")},
            "netbox_circuit": {Tok: netboxResource(netboxMod, "Circuit")},
            "netbox_circuit_provider": {Tok: netboxResource(netboxMod, "CircuitProvider")},
            "netbox_circuit_termination": {Tok: netboxResource(netboxMod, "CircuitTermination")},
            "netbox_circuit_type": {Tok: netboxResource(netboxMod, "CircuitType")},
            "netbox_cluster": {Tok: netboxResource(netboxMod, "Cluster")},
            "netbox_cluster_group": {Tok: netboxResource(netboxMod, "ClusterGroup")},
            "netbox_cluster_type": {Tok: netboxResource(netboxMod, "ClusterType")},
            "netbox_custom_field": {Tok: netboxResource(netboxMod, "CustomField")},
            "netbox_device": {Tok: netboxResource(netboxMod, "Device")},
            "netbox_device_role": {Tok: netboxResource(netboxMod, "DeviceRole")},
            "netbox_device_type": {Tok: netboxResource(netboxMod, "DeviceType")},
            "netbox_interface": {Tok: netboxResource(netboxMod, "Interface")},
            "netbox_ip_address": {
                Tok: netboxResource(netboxMod, "IpAddress"),
                Fields: map[string]*tfbridge.SchemaInfo{
                    "ip_address": {
                        CSharpName: "IpAddressOutput",
                    },
                },
            },
            "netbox_ip_range": {Tok: netboxResource(netboxMod, "IpRange")},
            "netbox_ipam_role": {Tok: netboxResource(netboxMod, "IpamRole")},
            "netbox_manufacturer": {Tok: netboxResource(netboxMod, "Manufacturer")},
            "netbox_platform": {Tok: netboxResource(netboxMod, "Platform")},
            "netbox_prefix": {
                Tok: netboxResource(netboxMod, "Prefix"),
                Fields: map[string]*tfbridge.SchemaInfo{
                    "prefix": {
                        CSharpName: "PrefixOutput",
                    },
                },
            },
            "netbox_primary_ip": {Tok: netboxResource(netboxMod, "PrimaryIp")},
            "netbox_region": {Tok: netboxResource(netboxMod, "Region")},
            "netbox_rir": {Tok: netboxResource(netboxMod, "Rir")},
            "netbox_service": {Tok: netboxResource(netboxMod, "Service")},
            "netbox_site": {Tok: netboxResource(netboxMod, "Site")},
            "netbox_tag": {Tok: netboxResource(netboxMod, "Tag")},
            "netbox_tenant": {Tok: netboxResource(netboxMod, "Tenant")},
            "netbox_tenant_group": {Tok: netboxResource(netboxMod, "TenantGroup")},
            "netbox_token": {Tok: netboxResource(netboxMod, "Token")},
            "netbox_user": {Tok: netboxResource(netboxMod, "User")},
            "netbox_virtual_machine": {Tok: netboxResource(netboxMod, "VirtualMachine")},
            "netbox_vlan": {Tok: netboxResource(netboxMod, "Vlan")},
            "netbox_vrf": {Tok: netboxResource(netboxMod, "Vrf")},
						"netbox_asn": {Tok: netboxResource(netboxMod, "Asn")},    
						"netbox_cable": {Tok: netboxResource(netboxMod, "Cable")},    
						"netbox_contact": {Tok: netboxResource(netboxMod, "Contact")},    
						"netbox_contact_assignment": {Tok: netboxResource(netboxMod, "ContactAssignment")},    
						"netbox_contact_group": {Tok: netboxResource(netboxMod, "ContactGroup")},    
						"netbox_contact_role": {Tok: netboxResource(netboxMod, "ContactRole")},    
						"netbox_custom_field_choice_set": {Tok: netboxResource(netboxMod, "CustomFieldChoiceSet")},    
						"netbox_device_console_port": {Tok: netboxResource(netboxMod, "DeviceConsolePort")},    
						"netbox_device_console_server_port": {Tok: netboxResource(netboxMod, "DeviceConsoleServerPort")},    
						"netbox_device_front_port": {Tok: netboxResource(netboxMod, "DeviceFrontPort")},    
						"netbox_device_interface": {Tok: netboxResource(netboxMod, "DeviceInterface")},    
						"netbox_device_module_bay": {Tok: netboxResource(netboxMod, "DeviceModuleBay")},    
						"netbox_device_power_outlet": {Tok: netboxResource(netboxMod, "DevicePowerOutlet")},    
						"netbox_device_power_port": {Tok: netboxResource(netboxMod, "DevicePowerPort")},    
						"netbox_device_primary_ip": {Tok: netboxResource(netboxMod, "DevicePrimaryIp")},    
						"netbox_device_rear_port": {Tok: netboxResource(netboxMod, "DeviceRearPort")},    
						"netbox_event_rule": {Tok: netboxResource(netboxMod, "EventRule")},    
						"netbox_inventory_item": {Tok: netboxResource(netboxMod, "InventoryItem")},    
						"netbox_inventory_item_role": {Tok: netboxResource(netboxMod, "InventoryItemRole")},    
						"netbox_location": {Tok: netboxResource(netboxMod, "Location")},    
						"netbox_module": {Tok: netboxResource(netboxMod, "Module")},    
						"netbox_module_type": {Tok: netboxResource(netboxMod, "ModuleType")},    
						"netbox_permission": {Tok: netboxResource(netboxMod, "Permission")},    
						"netbox_power_feed": {Tok: netboxResource(netboxMod, "PowerFeed")},    
						"netbox_power_panel": {Tok: netboxResource(netboxMod, "PowerPanel")},    
						"netbox_rack": {Tok: netboxResource(netboxMod, "Rack")},    
						"netbox_rack_reservation": {Tok: netboxResource(netboxMod, "RackReservation")},    
						"netbox_rack_role": {Tok: netboxResource(netboxMod, "RackRole")},    
						"netbox_route_target": {Tok: netboxResource(netboxMod, "RouteTarget")},    
						"netbox_site_group": {Tok: netboxResource(netboxMod, "SiteGroup")},    
						"netbox_virtual_chassis": {Tok: netboxResource(netboxMod, "VirtualChassis")},    
						"netbox_virtual_disk": {Tok: netboxResource(netboxMod, "VirtualDisk")},    
						"netbox_vlan_group": {Tok: netboxResource(netboxMod, "VlanGroup")},    
						"netbox_vpn_tunnel": {Tok: netboxResource(netboxMod, "VpnTunnel")},    
						"netbox_vpn_tunnel_group": {Tok: netboxResource(netboxMod, "VpnTunnelGroup")},    
						"netbox_vpn_tunnel_termination": {Tok: netboxResource(netboxMod, "VpnTunnelTermination")},    
						"netbox_webhook": {Tok: netboxResource(netboxMod, "Webhook")},    
            
		},

		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
            "netbox_cluster": {Tok: netboxDataSource(netboxMod, "getCluster")},
            "netbox_cluster_group": {Tok: netboxDataSource(netboxMod, "getClusterGroup")},
            "netbox_cluster_type": {Tok: netboxDataSource(netboxMod, "getClusterType")},
            "netbox_device_role": {Tok: netboxDataSource(netboxMod, "getDeviceRole")},
            "netbox_device_type": {Tok: netboxDataSource(netboxMod, "getDeviceType")},
            "netbox_interfaces": {Tok: netboxDataSource(netboxMod, "getInterfaces")},
            "netbox_ip_addresses": {Tok: netboxDataSource(netboxMod, "getIpAddresses")},
            "netbox_ip_range": {Tok: netboxDataSource(netboxMod, "getIpRange")},
            "netbox_platform": {Tok: netboxDataSource(netboxMod, "getPlatform")},
            "netbox_prefix": {Tok: netboxDataSource(netboxMod, "getPrefix")},
            "netbox_region": {Tok: netboxDataSource(netboxMod, "getRegion")},
            "netbox_site": {Tok: netboxDataSource(netboxMod, "getSite")},
            "netbox_tag": {Tok: netboxDataSource(netboxMod, "getTag")},
            "netbox_tenant": {Tok: netboxDataSource(netboxMod, "getTenant")},
            "netbox_tenant_group": {Tok: netboxDataSource(netboxMod, "getTenantGroup")},
            "netbox_tenants": {Tok: netboxDataSource(netboxMod, "getTenants")},
            "netbox_virtual_machines": {Tok: netboxDataSource(netboxMod, "getVirtualMachines")},
            "netbox_vlan": {Tok: netboxDataSource(netboxMod, "getVlan")},
            "netbox_vrf": {Tok: netboxDataSource(netboxMod, "getVrf")},
						"netbox_asn": {Tok: netboxDataSource(netboxMod, "getAsn")},     
						"netbox_asns": {Tok: netboxDataSource(netboxMod, "getAsns")},     
						"netbox_available_prefix": {Tok: netboxDataSource(netboxMod, "getAvailablePrefix")},     
						"netbox_contact": {Tok: netboxDataSource(netboxMod, "getContact")},     
						"netbox_contact_group": {Tok: netboxDataSource(netboxMod, "getContactGroup")},     
						"netbox_contact_role": {Tok: netboxDataSource(netboxMod, "getContactRole")},     
						"netbox_device_interfaces": {Tok: netboxDataSource(netboxMod, "getDeviceInterfaces")},     
						"netbox_devices": {Tok: netboxDataSource(netboxMod, "getDevices")},     
						"netbox_ipam_role": {Tok: netboxDataSource(netboxMod, "getIpamRole")},     
						"netbox_location": {Tok: netboxDataSource(netboxMod, "getLocation")},     
						"netbox_locations": {Tok: netboxDataSource(netboxMod, "getLocations")},     
						"netbox_prefixes": {Tok: netboxDataSource(netboxMod, "getPrefixes")},     
						"netbox_rack_role": {Tok: netboxDataSource(netboxMod, "getRackRole")},     
						"netbox_racks": {Tok: netboxDataSource(netboxMod, "getRacks")},     
						"netbox_route_target": {Tok: netboxDataSource(netboxMod, "getRouteTarget")},     
						"netbox_site_group": {Tok: netboxDataSource(netboxMod, "getSiteGroup")},     
						"netbox_tags": {Tok: netboxDataSource(netboxMod, "getTags")},     
						"netbox_vlan_group": {Tok: netboxDataSource(netboxMod, "getVlanGroup")},     
						"netbox_vlans": {Tok: netboxDataSource(netboxMod, "getVlans")},     
						"netbox_vrfs": {Tok: netboxDataSource(netboxMod, "getVrfs")},     
		},

	
		Python: &tfbridge.PythonInfo{
            PackageName: "spk_pulumi_netbox",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
