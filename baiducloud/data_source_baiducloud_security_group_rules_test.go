package baiducloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

const (
	testAccSecurityGroupRulesDataSourceName          = "data.baiducloud_security_group_rules.default"
	testAccSecurityGroupRulesDataSourceAttrKeyPrefix = "rules.0."
)

func TestAccBaiduCloudSecurityGroupRulesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,

		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupRulesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaiduCloudDataSourceId(testAccSecurityGroupRulesDataSourceName),
					resource.TestCheckResourceAttr(testAccSecurityGroupRulesDataSourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(testAccSecurityGroupRulesDataSourceName, testAccSecurityGroupRulesDataSourceAttrKeyPrefix+"direction", "ingress"),
					resource.TestCheckResourceAttr(testAccSecurityGroupRulesDataSourceName, testAccSecurityGroupRulesDataSourceAttrKeyPrefix+"protocol", "udp"),
					resource.TestCheckResourceAttr(testAccSecurityGroupRulesDataSourceName, testAccSecurityGroupRulesDataSourceAttrKeyPrefix+"port_range", "1-65523"),
					resource.TestCheckResourceAttr(testAccSecurityGroupRulesDataSourceName, testAccSecurityGroupRulesDataSourceAttrKeyPrefix+"remark", "remark"),
					resource.TestCheckResourceAttr(testAccSecurityGroupRulesDataSourceName, testAccSecurityGroupRulesDataSourceAttrKeyPrefix+"ether_type", "IPv4"),
					resource.TestCheckResourceAttr(testAccSecurityGroupRulesDataSourceName, testAccSecurityGroupRulesDataSourceAttrKeyPrefix+"source_ip", "all"),
				),
			},
		},
	})
}

func testAccSecurityGroupRulesDataSourceConfig() string {
	return fmt.Sprintf(`
resource "baiducloud_vpc" "default" {
  name = "%s"
  description = "test"
  cidr = "192.168.0.0/24"
}

resource "baiducloud_security_group" "default" {
  name        = "%s"
  description = "Baidu acceptance test"
  vpc_id      = baiducloud_vpc.default.id
}

resource "baiducloud_security_group_rule" "default" {
  security_group_id = baiducloud_security_group.default.id
  remark            = "remark"
  protocol          = "udp"
  port_range        = "1-65523"
  direction         = "ingress"
}

data "baiducloud_security_group_rules" "default" {
  security_group_id = baiducloud_security_group_rule.default.security_group_id
  vpc_id            = baiducloud_security_group.default.vpc_id
}
`, BaiduCloudTestResourceAttrNamePrefix+"VPC",
		BaiduCloudTestResourceAttrNamePrefix+"SecurityGroup")
}
