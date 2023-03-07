// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package rds

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
	return []func(context.Context) (resource.ResourceWithConfigure, error){
		newResourceExportTask,
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_db_cluster_snapshot":            DataSourceClusterSnapshot,
		"aws_db_event_categories":            DataSourceEventCategories,
		"aws_db_instance":                    DataSourceInstance,
		"aws_db_instances":                   DataSourceInstances,
		"aws_db_proxy":                       DataSourceProxy,
		"aws_db_snapshot":                    DataSourceSnapshot,
		"aws_db_subnet_group":                DataSourceSubnetGroup,
		"aws_rds_certificate":                DataSourceCertificate,
		"aws_rds_cluster":                    DataSourceCluster,
		"aws_rds_clusters":                   DataSourceClusters,
		"aws_rds_engine_version":             DataSourceEngineVersion,
		"aws_rds_orderable_db_instance":      DataSourceOrderableInstance,
		"aws_rds_reserved_instance_offering": DataSourceReservedOffering,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_db_cluster_snapshot":                       ResourceClusterSnapshot,
		"aws_db_event_subscription":                     ResourceEventSubscription,
		"aws_db_instance":                               ResourceInstance,
		"aws_db_instance_automated_backups_replication": ResourceInstanceAutomatedBackupsReplication,
		"aws_db_instance_role_association":              ResourceInstanceRoleAssociation,
		"aws_db_option_group":                           ResourceOptionGroup,
		"aws_db_parameter_group":                        ResourceParameterGroup,
		"aws_db_proxy":                                  ResourceProxy,
		"aws_db_proxy_default_target_group":             ResourceProxyDefaultTargetGroup,
		"aws_db_proxy_endpoint":                         ResourceProxyEndpoint,
		"aws_db_proxy_target":                           ResourceProxyTarget,
		"aws_db_security_group":                         ResourceSecurityGroup,
		"aws_db_snapshot":                               ResourceSnapshot,
		"aws_db_snapshot_copy":                          ResourceSnapshotCopy,
		"aws_db_subnet_group":                           ResourceSubnetGroup,
		"aws_rds_cluster":                               ResourceCluster,
		"aws_rds_cluster_activity_stream":               ResourceClusterActivityStream,
		"aws_rds_cluster_endpoint":                      ResourceClusterEndpoint,
		"aws_rds_cluster_instance":                      ResourceClusterInstance,
		"aws_rds_cluster_parameter_group":               ResourceClusterParameterGroup,
		"aws_rds_cluster_role_association":              ResourceClusterRoleAssociation,
		"aws_rds_global_cluster":                        ResourceGlobalCluster,
		"aws_rds_reserved_instance":                     ResourceReservedInstance,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RDS
}

var ServicePackage = &servicePackage{}