package aws

import (
	"testing"
	/*
		"github.com/aws/aws-sdk-go/service/elasticache"
		"github.com/hashicorp/terraform/helper/resource"
		"github.com/hashicorp/terraform/terraform"
		"fmt"
		"github.com/aws/aws-sdk-go/aws/awserr"
	*/)

func TestAccAWSElasticacheReplicationGroup_basic(t *testing.T) {
	/*
		var v elasticache.CacheReplicationGroup

		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testAccPreCheck(t) },
			Providers:    testAccProviders,
			CheckDestroy: testAccCheckAWSElasticacheReplicationGroupDestroy,
			Steps: []resource.TestStep{
				resource.TestStep{
					Config: testAccAWSElasticacheReplicationGroupConfig,
					Check: resource.ComposeTestCheckFunc(
						testAccCheckAWSElasticacheReplicationGroupExists("aws_elasticache_parameter_group.bar", &v),
						testAccCheckAWSElasticacheReplicationGroupAttributes(&v),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "name", "parameter-group-test-terraform"),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "family", "redis2.8"),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "description", "Managed by Terraform"),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "parameter.283487565.name", "appendonly"),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "parameter.283487565.value", "yes"),
					),
				},
				resource.TestStep{
					Config: testAccAWSElasticacheReplicationGroupAddParametersConfig,
					Check: resource.ComposeTestCheckFunc(
						testAccCheckAWSElasticacheReplicationGroupExists("aws_elasticache_parameter_group.bar", &v),
						testAccCheckAWSElasticacheReplicationGroupAttributes(&v),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "name", "parameter-group-test-terraform"),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "family", "redis2.8"),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "description", "Test parameter group for terraform"),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "parameter.283487565.name", "appendonly"),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "parameter.283487565.value", "yes"),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "parameter.2196914567.name", "appendfsync"),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "parameter.2196914567.value", "always"),
					),
				},
			},
		})
	*/
}

func TestAccAWSElasticacheReplicationGroupOnly(t *testing.T) {
	/*
		var v elasticache.CacheReplicationGroup

		resource.Test(t, resource.TestCase{
			PreCheck:     func() { testAccPreCheck(t) },
			Providers:    testAccProviders,
			CheckDestroy: testAccCheckAWSElasticacheReplicationGroupDestroy,
			Steps: []resource.TestStep{
				resource.TestStep{
					Config: testAccAWSElasticacheReplicationGroupOnlyConfig,
					Check: resource.ComposeTestCheckFunc(
						testAccCheckAWSElasticacheReplicationGroupExists("aws_elasticache_parameter_group.bar", &v),
						testAccCheckAWSElasticacheReplicationGroupAttributes(&v),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "name", "parameter-group-test-terraform"),
						resource.TestCheckResourceAttr(
							"aws_elasticache_parameter_group.bar", "family", "redis2.8"),
					),
				},
			},
		})
	*/
}

/*
func testAccCheckAWSElasticacheReplicationGroupDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).elasticacheconn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_elasticache_parameter_group" {
			continue
		}

		// Try to find the Group
		resp, err := conn.DescribeCacheReplicationGroups(
			&elasticache.DescribeCacheReplicationGroupsInput{
				CacheReplicationGroupName: aws.String(rs.Primary.ID),
			})

		if err == nil {
			if len(resp.CacheReplicationGroups) != 0 &&
				*resp.CacheReplicationGroups[0].CacheReplicationGroupName == rs.Primary.ID {
				return fmt.Errorf("Cache Parameter Group still exists")
			}
		}

		// Verify the error
		newerr, ok := err.(awserr.Error)
		if !ok {
			return err
		}
		if newerr.Code() != "CacheReplicationGroupNotFound" {
			return err
		}
	}
	return nil
}
*/
/*
func testAccCheckAWSElasticacheReplicationGroupAttributes(v *elasticache.CacheReplicationGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *v.CacheReplicationGroupName != "parameter-group-test-terraform" {
			return fmt.Errorf("bad name: %#v", v.CacheReplicationGroupName)
		}

		if *v.CacheReplicationGroupFamily != "redis2.8" {
			return fmt.Errorf("bad family: %#v", v.CacheReplicationGroupFamily)
		}
		return nil
	}
}
*/
/*
func testAccCheckAWSElasticacheReplicationGroupExists(n string, v *elasticache.CacheReplicationGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cache Parameter Group ID is set")
		}

		conn := testAccProvider.Meta().(*AWSClient).elasticacheconn

		opts := elasticache.DescribeCacheReplicationGroupsInput{
			CacheReplicationGroupName: aws.String(rs.Primary.ID),
		}

		resp, err := conn.DescribeCacheReplicationGroups(&opts)

		if err != nil {
			return err
		}

		if len(resp.CacheReplicationGroups) != 1 ||
			*resp.CacheReplicationGroups[0].CacheReplicationGroupName != rs.Primary.ID {
			return fmt.Errorf("Cache Parameter Group not found")
		}

		*v = *resp.CacheReplicationGroups[0]
		return nil
	}
}
*/

/*
const testAccAWSElasticacheReplicationGroupConfig = `
resource "aws_elasticache_parameter_group" "bar" {
	name = "parameter-group-test-terraform"
	family = "redis2.8"
	parameter {
	  name = "appendonly"
	  value = "yes"
	}
}
`

const testAccAWSElasticacheReplicationGroupAddParametersConfig = `
resource "aws_elasticache_parameter_group" "bar" {
	name = "parameter-group-test-terraform"
	family = "redis2.8"
	description = "Test parameter group for terraform"
	parameter {
	  name = "appendonly"
	  value = "yes"
	}
	parameter {
	  name = "appendfsync"
	  value = "always"
	}
}
`

const testAccAWSElasticacheReplicationGroupOnlyConfig = `
resource "aws_elasticache_parameter_group" "bar" {
	name = "parameter-group-test-terraform"
	family = "redis2.8"
	description = "Test parameter group for terraform"
}
`
*/
