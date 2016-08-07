package aws

import (
	"log"

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
	log.Printf("[DEBUG] Create Cache Replication Group: %v", conn)
	return resourceAwsElasticacheReplicationGroupUpdate(d, meta)
}

func resourceAwsElasticacheReplicationGroupRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).elasticacheconn
	log.Printf("[DEBUG] Read Cache Replication Group: %v", conn)
	return nil
}

func resourceAwsElasticacheReplicationGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).elasticacheconn
	log.Printf("[DEBUG] Update Cache Replication Group: %v", conn)
	return resourceAwsElasticacheReplicationGroupRead(d, meta)
}

func resourceAwsElasticacheReplicationGroupDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Delete Cache Replication Group: %v", d)
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
