// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ssm

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
		"aws_ssm_document":            DataSourceDocument,
		"aws_ssm_instances":           DataSourceInstances,
		"aws_ssm_maintenance_windows": DataSourceMaintenanceWindows,
		"aws_ssm_parameter":           DataSourceParameter,
		"aws_ssm_parameters_by_path":  DataSourceParametersByPath,
		"aws_ssm_patch_baseline":      DataSourcePatchBaseline,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_ssm_activation":                ResourceActivation,
		"aws_ssm_association":               ResourceAssociation,
		"aws_ssm_default_patch_baseline":    ResourceDefaultPatchBaseline,
		"aws_ssm_document":                  ResourceDocument,
		"aws_ssm_maintenance_window":        ResourceMaintenanceWindow,
		"aws_ssm_maintenance_window_target": ResourceMaintenanceWindowTarget,
		"aws_ssm_maintenance_window_task":   ResourceMaintenanceWindowTask,
		"aws_ssm_parameter":                 ResourceParameter,
		"aws_ssm_patch_baseline":            ResourcePatchBaseline,
		"aws_ssm_patch_group":               ResourcePatchGroup,
		"aws_ssm_resource_data_sync":        ResourceResourceDataSync,
		"aws_ssm_service_setting":           ResourceServiceSetting,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SSM
}

var ServicePackage = &servicePackage{}