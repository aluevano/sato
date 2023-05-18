package arm

import (
	_ "embed" //required for embed
)

//go:embed resources/azurerm_virtual_machine.template
var azurermVirtualMachine []byte

//go:embed resources/azurerm_virtual_machine_extension.template
var azurermVirtualMachineExtension []byte

//go:embed resources/azurerm_network_interface.template
var azurermNetworkInterface []byte

//go:embed resources/azurerm_network_security_group.template
var azurermNetworkSecurityGroup []byte

//go:embed resources/azurerm_public_ip.template
var azurermPublicIP []byte

//go:embed resources/azurerm_virtual_network.template
var azurermVirtualNetwork []byte

//go:embed resources/azurerm_storage_account.template
var azurermStorageAccount []byte

//go:embed resources/azurerm_subnet.template
var azurermSubnet []byte

//go:embed resources/azurerm_analysis_services_server.template
var azurermAnalysisServicesServer []byte

//go:embed resources/azurerm_api_management.template
var azurermApiManagement []byte
