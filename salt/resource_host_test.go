package salt

import (
	"fmt"
	"testing"

	r "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestSaltHost(t *testing.T) {
	r.Test(t, r.TestCase{
		Providers: testAccProviders,
		Steps: []r.TestStep{
			r.TestStep{
				Config: testSaltHostConfig,
				Check: func(s *terraform.State) error {
					rID := s.RootModule().Outputs["host_id"].Value
					if "example.medstack.net" != rID {
						return fmt.Errorf("Unexpected value for resource ID: %s", rID)
					}

					return nil
				},
			},
		},
	})
}

var testSaltHostConfig = `
resource "salt_host" "test" {
	host = "example.medstack.net"
}

output "host_id" {
	value = "${salt_host.test.id}"
}
`
