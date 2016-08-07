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
				ForceNew: true,
				Required: true,
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
