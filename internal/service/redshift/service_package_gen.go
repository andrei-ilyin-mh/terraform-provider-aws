// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package redshift

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
		"aws_redshift_cluster":             DataSourceCluster,
		"aws_redshift_cluster_credentials": DataSourceClusterCredentials,
		"aws_redshift_orderable_cluster":   DataSourceOrderableCluster,
		"aws_redshift_service_account":     DataSourceServiceAccount,
		"aws_redshift_subnet_group":        DataSourceSubnetGroup,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_redshift_authentication_profile":        ResourceAuthenticationProfile,
		"aws_redshift_cluster":                       ResourceCluster,
		"aws_redshift_cluster_iam_roles":             ResourceClusterIAMRoles,
		"aws_redshift_cluster_snapshot":              ResourceClusterSnapshot,
		"aws_redshift_endpoint_access":               ResourceEndpointAccess,
		"aws_redshift_endpoint_authorization":        ResourceEndpointAuthorization,
		"aws_redshift_event_subscription":            ResourceEventSubscription,
		"aws_redshift_hsm_client_certificate":        ResourceHSMClientCertificate,
		"aws_redshift_hsm_configuration":             ResourceHSMConfiguration,
		"aws_redshift_parameter_group":               ResourceParameterGroup,
		"aws_redshift_partner":                       ResourcePartner,
		"aws_redshift_scheduled_action":              ResourceScheduledAction,
		"aws_redshift_security_group":                ResourceSecurityGroup,
		"aws_redshift_snapshot_copy_grant":           ResourceSnapshotCopyGrant,
		"aws_redshift_snapshot_schedule":             ResourceSnapshotSchedule,
		"aws_redshift_snapshot_schedule_association": ResourceSnapshotScheduleAssociation,
		"aws_redshift_subnet_group":                  ResourceSubnetGroup,
		"aws_redshift_usage_limit":                   ResourceUsageLimit,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Redshift
}

var ServicePackage = &servicePackage{}