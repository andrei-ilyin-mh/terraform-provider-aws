// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package backup

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
		"aws_backup_framework":   DataSourceFramework,
		"aws_backup_plan":        DataSourcePlan,
		"aws_backup_report_plan": DataSourceReportPlan,
		"aws_backup_selection":   DataSourceSelection,
		"aws_backup_vault":       DataSourceVault,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_backup_framework":                ResourceFramework,
		"aws_backup_global_settings":          ResourceGlobalSettings,
		"aws_backup_plan":                     ResourcePlan,
		"aws_backup_region_settings":          ResourceRegionSettings,
		"aws_backup_report_plan":              ResourceReportPlan,
		"aws_backup_selection":                ResourceSelection,
		"aws_backup_vault":                    ResourceVault,
		"aws_backup_vault_lock_configuration": ResourceVaultLockConfiguration,
		"aws_backup_vault_notifications":      ResourceVaultNotifications,
		"aws_backup_vault_policy":             ResourceVaultPolicy,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Backup
}

var ServicePackage = &servicePackage{}