package cloudtrail

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

// @SDKResource("aws_cloudtrail_event_data_store")
func ResourceEventDataStore() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceEventDataStoreCreate,
		ReadWithoutTimeout:   resourceEventDataStoreRead,
		UpdateWithoutTimeout: resourceEventDataStoreUpdate,
		DeleteWithoutTimeout: resourceEventDataStoreDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		CustomizeDiff: verify.SetTagsDiff,

		Schema: map[string]*schema.Schema{
			"advanced_event_selector": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_selector": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ends_with": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringLenBetween(1, 2048),
										},
									},
									"equals": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringLenBetween(1, 2048),
										},
									},
									"field": {
										Type:         schema.TypeString,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.StringInSlice(field_Values(), false),
									},
									"not_ends_with": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringLenBetween(1, 2048),
										},
									},
									"not_equals": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringLenBetween(1, 2048),
										},
									},
									"not_starts_with": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringLenBetween(1, 2048),
										},
									},
									"starts_with": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MinItems: 1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringLenBetween(1, 2048),
										},
									},
								},
							},
						},
						"name": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringLenBetween(0, 1000),
						},
					},
				},
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"multi_region_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(3, 128),
			},
			"organization_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"retention_period": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2555,
				ValidateFunc: validation.All(
					validation.IntBetween(7, 2555),
				),
			},
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
			"termination_protection_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceEventDataStoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).CloudTrailConn()
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	tags := defaultTagsConfig.MergeTags(tftags.New(ctx, d.Get("tags").(map[string]interface{})))

	name := d.Get("name").(string)
	input := &cloudtrail.CreateEventDataStoreInput{
		Name:                         aws.String(name),
		OrganizationEnabled:          aws.Bool(d.Get("organization_enabled").(bool)),
		MultiRegionEnabled:           aws.Bool(d.Get("multi_region_enabled").(bool)),
		TerminationProtectionEnabled: aws.Bool(d.Get("termination_protection_enabled").(bool)),
		RetentionPeriod:              aws.Int64(int64(d.Get("retention_period").(int))),
	}

	if _, ok := d.GetOk("advanced_event_selector"); ok {
		input.AdvancedEventSelectors = expandAdvancedEventSelector(d.Get("advanced_event_selector").([]interface{}))
	}

	if v, ok := d.GetOk("kms_key_id"); ok {
		input.KmsKeyId = aws.String(v.(string))
	}

	if len(tags) > 0 {
		input.TagsList = Tags(tags.IgnoreAWS())
	}

	output, err := conn.CreateEventDataStoreWithContext(ctx, input)

	if err != nil {
		return diag.Errorf("creating CloudTrail Event Data Store (%s): %s", name, err)
	}

	d.SetId(aws.StringValue(output.EventDataStoreArn))

	if err := waitEventDataStoreAvailable(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
		return diag.Errorf("waiting for CloudTrail Event Data Store (%s) create: %s", name, err)
	}

	return resourceEventDataStoreRead(ctx, d, meta)
}

func resourceEventDataStoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).CloudTrailConn()
	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	eventDataStore, err := FindEventDataStoreByARN(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] CloudTrail Event Data Store (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if err != nil {
		return diag.Errorf("reading CloudTrail Event Data Store (%s): %s", d.Id(), err)
	}

	if err := d.Set("advanced_event_selector", flattenAdvancedEventSelector(eventDataStore.AdvancedEventSelectors)); err != nil {
		return diag.Errorf("setting advanced_event_selector: %s", err)
	}
	d.Set("arn", eventDataStore.EventDataStoreArn)
	d.Set("kms_key_id", eventDataStore.KmsKeyId)
	d.Set("multi_region_enabled", eventDataStore.MultiRegionEnabled)
	d.Set("name", eventDataStore.Name)
	d.Set("organization_enabled", eventDataStore.OrganizationEnabled)
	d.Set("retention_period", eventDataStore.RetentionPeriod)
	d.Set("termination_protection_enabled", eventDataStore.TerminationProtectionEnabled)

	tags, err := ListTags(ctx, conn, d.Id())

	if err != nil {
		return diag.Errorf("listing tags for CloudTrail Event Data Store (%s): %s", d.Id(), err)
	}

	tags = tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig)

	//lintignore:AWSR002
	if err := d.Set("tags", tags.RemoveDefaultConfig(defaultTagsConfig).Map()); err != nil {
		return diag.Errorf("setting tags: %s", err)
	}

	if err := d.Set("tags_all", tags.Map()); err != nil {
		return diag.Errorf("setting tags_all: %s", err)
	}

	return nil
}

func resourceEventDataStoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).CloudTrailConn()

	if d.HasChangesExcept("tags", "tags_all") {
		input := &cloudtrail.UpdateEventDataStoreInput{
			EventDataStore: aws.String(d.Id()),
		}

		if d.HasChange("advanced_event_selector") {
			input.AdvancedEventSelectors = expandAdvancedEventSelector(d.Get("advanced_event_selector").([]interface{}))
		}

		if d.HasChange("multi_region_enabled") {
			input.MultiRegionEnabled = aws.Bool(d.Get("multi_region_enabled").(bool))
		}

		if d.HasChange("name") {
			input.Name = aws.String(d.Get("name").(string))
		}

		if d.HasChange("organization_enabled") {
			input.OrganizationEnabled = aws.Bool(d.Get("organization_enabled").(bool))
		}

		if d.HasChange("retention_period") {
			input.RetentionPeriod = aws.Int64(int64(d.Get("retention_period").(int)))
		}

		if d.HasChange("termination_protection_enabled") {
			input.TerminationProtectionEnabled = aws.Bool(d.Get("termination_protection_enabled").(bool))
		}

		_, err := conn.UpdateEventDataStoreWithContext(ctx, input)

		if err != nil {
			return diag.Errorf("updating CloudTrail Event Data Store (%s): %s", d.Id(), err)
		}

		if err := waitEventDataStoreAvailable(ctx, conn, d.Id(), d.Timeout(schema.TimeoutUpdate)); err != nil {
			return diag.Errorf("waiting for CloudTrail Event Data Store (%s) update: %s", d.Id(), err)
		}
	}

	if d.HasChange("tags_all") {
		o, n := d.GetChange("tags_all")

		if err := UpdateTags(ctx, conn, d.Id(), o, n); err != nil {
			return diag.Errorf("updating CloudTrail Event Data Store (%s) tags: %s", d.Id(), err)
		}
	}

	return resourceEventDataStoreRead(ctx, d, meta)
}

func resourceEventDataStoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).CloudTrailConn()

	log.Printf("[DEBUG] Deleting CloudTrail Event Data Store: %s", d.Id())
	_, err := conn.DeleteEventDataStoreWithContext(ctx, &cloudtrail.DeleteEventDataStoreInput{
		EventDataStore: aws.String(d.Id()),
	})

	if tfawserr.ErrCodeEquals(err, cloudtrail.ErrCodeEventDataStoreNotFoundException) {
		return nil
	}

	if err != nil {
		return diag.Errorf("deleting CloudTrail Event Data Store (%s): %s", d.Id(), err)
	}

	if err := waitEventDataStoreDeleted(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
		return diag.Errorf("waiting for CloudTrail Event Data Store (%s) delete: %s", d.Id(), err)
	}

	return nil
}

func FindEventDataStoreByARN(ctx context.Context, conn *cloudtrail.CloudTrail, arn string) (*cloudtrail.GetEventDataStoreOutput, error) {
	input := cloudtrail.GetEventDataStoreInput{
		EventDataStore: aws.String(arn),
	}

	output, err := conn.GetEventDataStoreWithContext(ctx, &input)

	if tfawserr.ErrCodeEquals(err, cloudtrail.ErrCodeEventDataStoreNotFoundException) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if status := aws.StringValue(output.Status); status == cloudtrail.EventDataStoreStatusPendingDeletion {
		return nil, &resource.NotFoundError{
			Message:     status,
			LastRequest: input,
		}
	}

	return output, nil
}

func statusEventDataStore(ctx context.Context, conn *cloudtrail.CloudTrail, arn string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		eventDataStore, err := FindEventDataStoreByARN(ctx, conn, arn)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return eventDataStore, aws.StringValue(eventDataStore.Status), nil
	}
}

func waitEventDataStoreAvailable(ctx context.Context, conn *cloudtrail.CloudTrail, arn string, timeout time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{cloudtrail.EventDataStoreStatusCreated},
		Target:  []string{cloudtrail.EventDataStoreStatusEnabled},
		Refresh: statusEventDataStore(ctx, conn, arn),
		Timeout: timeout,
	}

	_, err := stateConf.WaitForStateContext(ctx)

	return err
}

func waitEventDataStoreDeleted(ctx context.Context, conn *cloudtrail.CloudTrail, arn string, timeout time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{cloudtrail.EventDataStoreStatusCreated, cloudtrail.EventDataStoreStatusEnabled},
		Target:  []string{},
		Refresh: statusEventDataStore(ctx, conn, arn),
		Timeout: timeout,
	}

	_, err := stateConf.WaitForStateContext(ctx)

	return err
}