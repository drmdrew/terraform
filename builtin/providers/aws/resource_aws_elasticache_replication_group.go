package aws

import (
	"log"

	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsElasticacheReplicationGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsElasticacheReplicationGroupCreate,
		Read:   resourceAwsElasticacheReplicationGroupRead,
		Update: resourceAwsElasticacheReplicationGroupUpdate,
		Delete: resourceAwsElasticacheReplicationGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"replication_group_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"replication_group_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "Managed by Terraform",
			},
			"primary_cluster_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"automatic_failover_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"num_cache_clusters": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"preferred_cache_cluster_a_zs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cache_node_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"engine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "redis",
			},
			"engine_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cache_parameter_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cache_subnet_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cache_security_group_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cache_security_group_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"snapshot_arns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"snapshot_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"preferred_maitenance_window": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"notification_topic_arn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_minor_version_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"snapshot_retention_limit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"snapshot_window": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAwsElasticacheReplicationGroupCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).elasticacheconn
	req := elasticache.CreateReplicationGroupInput{
		ReplicationGroupId:          aws.String(d.Get("replication_group_id").(string)),
		ReplicationGroupDescription: aws.String(d.Get("replication_group_description").(string)),
		//		PrimaryClusterId:            aws.String(d.Get("primary_cluster_id")),
		//AutomaticFailoverEnabled: aws.Bool(d.Get("automatic_failover_enabled").(bool)),
		//		NumCacheClusters:            aws.Int(d.Get("num_cache_clusters")),
		//		PreferredCacheClusterAZs:    aws.String(d.Get("preferred_cache_cluster_a_zs")),
		CacheNodeType: aws.String(d.Get("cache_node_type").(string)),
		Engine:        aws.String(d.Get("engine").(string)),
		//		EngineVersion:               aws.String(d.Get("engine_version")),
		//		CacheParameterGroupName:     aws.String(d.Get("cache_parameter_group_name")),
		//		CacheSubnetGroupName: aws.String(d.Get("cache_subnet_group_name").(string)),
		//		CacheSecurityGroupNames: aws.String(d.Get("cache_security_group_names")),
		//		SnapshotArns: aws.String(d.Get("snapshot_arns")),
		//		SnapshotName: aws.String(d.Get("snapshot_name")),
		//		PreferredMaintenanceWindow: aws.String(d.Get("preferred_maitenance_window")),
		//		Port: aws.String(d.Get("port")),
		//		NotificationTopicArn: aws.String(d.Get("notification_topic_arn")),
		//		AutoMinorVersionUpgrade: aws.String(d.Get("auto_minor_version_upgrade")),
		//		SnapshotRetentionLimit: aws.String(d.Get("snapshot_retention_limit")),
		//		SnapshotWindow: aws.String(d.Get("snapshot_window")),
	}
	// TODO: cache_security_groupp_ids support??
	// CacheSecurityGroupIds: aws.String("cache_security_group_ids")),

	if v, ok := d.GetOk("primary_cluster_id"); ok {
		req.PrimaryClusterId = aws.String(v.(string))
	}

	if v, ok := d.GetOk("automatic_failover_enabled"); ok {
		req.AutomaticFailoverEnabled = aws.Bool(v.(bool))
	}

	if v, ok := d.GetOk("num_cache_clusters"); ok {
		req.NumCacheClusters = aws.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("cache_subnet_group_name"); ok {
		req.CacheSubnetGroupName = aws.String(v.(string))
	}

	if v, ok := d.GetOk("preferred_cache_cluster_a_zs"); ok && v.(*schema.Set).Len() > 0 {
		req.PreferredCacheClusterAZs = expandStringList(v.(*schema.Set).List())
	}

	log.Printf("[DEBUG] Create Cache Replication Group: %v, %v", conn, req)

	resp, err := conn.CreateReplicationGroup(&req)
	if err != nil {
		return fmt.Errorf("Error creating Elasticache: %s", err)
	}

	log.Printf("[DEBUG] Creating replication group id: %v", resp.ReplicationGroup.ReplicationGroupId)

	return resourceAwsElasticacheReplicationGroupRead(d, meta)
}

func resourceAwsElasticacheReplicationGroupRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).elasticacheconn
	log.Printf("[DEBUG] Read Cache Replication Group: %v", conn)

	req := &elasticache.DescribeReplicationGroupsInput{
		ReplicationGroupId: aws.String(d.Id()),
	}

	res, err := conn.DescribeReplicationGroups(req)
	if err != nil {
		log.Printf("[DEBUG] Replication Group error: %v", err)
		return err
	}

	if len(res.ReplicationGroups) == 1 {
		rg := res.ReplicationGroups[0]
		d.Set("replication_group_id", rg.ReplicationGroupId)
		d.Set("replication_group_description", rg.Description)
		//  d.Set("automatic_failover_enabled", rg.AutomaticFailover)
		log.Printf("[DEBUG] setting id and description: %v, %v", rg.ReplicationGroupId, rg.Description)
	}
	return nil
}

func resourceAwsElasticacheReplicationGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).elasticacheconn
	log.Printf("[DEBUG] Update Cache Replication Group: %v", conn)
	return resourceAwsElasticacheReplicationGroupRead(d, meta)
}

func resourceAwsElasticacheReplicationGroupDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).elasticacheconn

	req := &elasticache.DeleteReplicationGroupInput{
		ReplicationGroupId: aws.String(d.Id()),
	}
	_, err := conn.DeleteReplicationGroup(req)
	if err != nil {
		return err
	}
	/*
		//	log.Printf("[DEBUG] Waiting for deletion: %v", d.Id())
		//	stateConf := &resource.StateChangeConf{
		//		Pending:    []string{"creating", "available", "deleting", "incompatible-parameters", "incompatible-network", "restore-failed"},
		//		Target:     []string{},
		//		Refresh:    cacheClusterStateRefreshFunc(conn, d.Id(), "", []string{}),
		//		Timeout:    20 * time.Minute,
		//		Delay:      10 * time.Second,
		//		MinTimeout: 3 * time.Second,
		//	}
		//
		//	_, sterr := stateConf.WaitForState()
		//	if sterr != nil {
		//		return fmt.Errorf("Error waiting for elasticache (%s) to delete: %s", d.Id(), sterr)
		//	}
	*/
	d.SetId("")

	return nil
}

/*
func resourceAwsElasticacheReplicationGroupDeleteRefreshFunc(
	d *schema.ResourceData,
	meta interface{}) resource.StateRefreshFunc {
	conn := meta.(*AWSClient).elasticacheconn

	return func() (interface{}, string, error) {

		deleteOpts := elasticache.DeleteCacheReplicationGroupInput{
			CacheReplicationGroupName: aws.String(d.Id()),
		}

		if _, err := conn.DeleteCacheReplicationGroup(&deleteOpts); err != nil {
			elasticahceerr, ok := err.(awserr.Error)
			if ok && elasticahceerr.Code() == "CacheReplicationGroupNotFoundFault" {
				d.SetId("")
				return d, "error", err
			}
			return d, "error", err
		}
		return d, "destroyed", nil
	}
}

func resourceAwsElasticacheParameterHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m["name"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["value"].(string)))

	return hashcode.String(buf.String())
}
*/
