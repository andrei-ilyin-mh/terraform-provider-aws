// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_guardduty_detector": DataSourceDetector,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_guardduty_detector":                   ResourceDetector,
		"aws_guardduty_filter":                     ResourceFilter,
		"aws_guardduty_invite_accepter":            ResourceInviteAccepter,
		"aws_guardduty_ipset":                      ResourceIPSet,
		"aws_guardduty_member":                     ResourceMember,
		"aws_guardduty_organization_admin_account": ResourceOrganizationAdminAccount,
		"aws_guardduty_organization_configuration": ResourceOrganizationConfiguration,
		"aws_guardduty_publishing_destination":     ResourcePublishingDestination,
		"aws_guardduty_threatintelset":             ResourceThreatIntelSet,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.GuardDuty
}

var ServicePackage = &servicePackage{}