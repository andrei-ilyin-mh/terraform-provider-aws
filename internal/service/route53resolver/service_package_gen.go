// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package route53resolver

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
		"aws_route53_resolver_endpoint":                        DataSourceEndpoint,
		"aws_route53_resolver_firewall_config":                 DataSourceFirewallConfig,
		"aws_route53_resolver_firewall_domain_list":            DataSourceFirewallDomainList,
		"aws_route53_resolver_firewall_rule_group":             DataSourceFirewallRuleGroup,
		"aws_route53_resolver_firewall_rule_group_association": DataSourceFirewallRuleGroupAssociation,
		"aws_route53_resolver_firewall_rules":                  DataSourceResolverFirewallRules,
		"aws_route53_resolver_rule":                            DataSourceRule,
		"aws_route53_resolver_rules":                           DataSourceRules,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_route53_resolver_config":                          ResourceConfig,
		"aws_route53_resolver_dnssec_config":                   ResourceDNSSECConfig,
		"aws_route53_resolver_endpoint":                        ResourceEndpoint,
		"aws_route53_resolver_firewall_config":                 ResourceFirewallConfig,
		"aws_route53_resolver_firewall_domain_list":            ResourceFirewallDomainList,
		"aws_route53_resolver_firewall_rule":                   ResourceFirewallRule,
		"aws_route53_resolver_firewall_rule_group":             ResourceFirewallRuleGroup,
		"aws_route53_resolver_firewall_rule_group_association": ResourceFirewallRuleGroupAssociation,
		"aws_route53_resolver_query_log_config":                ResourceQueryLogConfig,
		"aws_route53_resolver_query_log_config_association":    ResourceQueryLogConfigAssociation,
		"aws_route53_resolver_rule":                            ResourceRule,
		"aws_route53_resolver_rule_association":                ResourceRuleAssociation,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Route53Resolver
}

var ServicePackage = &servicePackage{}