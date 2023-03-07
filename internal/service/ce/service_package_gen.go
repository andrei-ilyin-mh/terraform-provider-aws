// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ce

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
		"aws_ce_cost_category": DataSourceCostCategory,
		"aws_ce_tags":          DataSourceTags,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_ce_anomaly_monitor":      ResourceAnomalyMonitor,
		"aws_ce_anomaly_subscription": ResourceAnomalySubscription,
		"aws_ce_cost_allocation_tag":  ResourceCostAllocationTag,
		"aws_ce_cost_category":        ResourceCostCategory,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CE
}

var ServicePackage = &servicePackage{}