// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package sagemaker

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
		"aws_sagemaker_prebuilt_ecr_image": DataSourcePrebuiltECRImage,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_sagemaker_app":                                       ResourceApp,
		"aws_sagemaker_app_image_config":                          ResourceAppImageConfig,
		"aws_sagemaker_code_repository":                           ResourceCodeRepository,
		"aws_sagemaker_device":                                    ResourceDevice,
		"aws_sagemaker_device_fleet":                              ResourceDeviceFleet,
		"aws_sagemaker_domain":                                    ResourceDomain,
		"aws_sagemaker_endpoint":                                  ResourceEndpoint,
		"aws_sagemaker_endpoint_configuration":                    ResourceEndpointConfiguration,
		"aws_sagemaker_feature_group":                             ResourceFeatureGroup,
		"aws_sagemaker_flow_definition":                           ResourceFlowDefinition,
		"aws_sagemaker_human_task_ui":                             ResourceHumanTaskUI,
		"aws_sagemaker_image":                                     ResourceImage,
		"aws_sagemaker_image_version":                             ResourceImageVersion,
		"aws_sagemaker_model":                                     ResourceModel,
		"aws_sagemaker_model_package_group":                       ResourceModelPackageGroup,
		"aws_sagemaker_model_package_group_policy":                ResourceModelPackageGroupPolicy,
		"aws_sagemaker_notebook_instance":                         ResourceNotebookInstance,
		"aws_sagemaker_notebook_instance_lifecycle_configuration": ResourceNotebookInstanceLifeCycleConfiguration,
		"aws_sagemaker_project":                                   ResourceProject,
		"aws_sagemaker_servicecatalog_portfolio_status":           ResourceServicecatalogPortfolioStatus,
		"aws_sagemaker_space":                                     ResourceSpace,
		"aws_sagemaker_studio_lifecycle_config":                   ResourceStudioLifecycleConfig,
		"aws_sagemaker_user_profile":                              ResourceUserProfile,
		"aws_sagemaker_workforce":                                 ResourceWorkforce,
		"aws_sagemaker_workteam":                                  ResourceWorkteam,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SageMaker
}

var ServicePackage = &servicePackage{}