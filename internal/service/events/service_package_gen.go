// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package events

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
		"aws_cloudwatch_event_bus":        DataSourceBus,
		"aws_cloudwatch_event_connection": DataSourceConnection,
		"aws_cloudwatch_event_source":     DataSourceSource,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_cloudwatch_event_api_destination": ResourceAPIDestination,
		"aws_cloudwatch_event_archive":         ResourceArchive,
		"aws_cloudwatch_event_bus":             ResourceBus,
		"aws_cloudwatch_event_bus_policy":      ResourceBusPolicy,
		"aws_cloudwatch_event_connection":      ResourceConnection,
		"aws_cloudwatch_event_permission":      ResourcePermission,
		"aws_cloudwatch_event_rule":            ResourceRule,
		"aws_cloudwatch_event_target":          ResourceTarget,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Events
}

var ServicePackage = &servicePackage{}