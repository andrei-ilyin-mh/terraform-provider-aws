// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package appstream

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
	return map[string]func() *schema.Resource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_appstream_directory_config":        ResourceDirectoryConfig,
		"aws_appstream_fleet":                   ResourceFleet,
		"aws_appstream_fleet_stack_association": ResourceFleetStackAssociation,
		"aws_appstream_image_builder":           ResourceImageBuilder,
		"aws_appstream_stack":                   ResourceStack,
		"aws_appstream_user":                    ResourceUser,
		"aws_appstream_user_stack_association":  ResourceUserStackAssociation,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppStream
}

var ServicePackage = &servicePackage{}