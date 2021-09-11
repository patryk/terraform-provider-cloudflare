package cloudflare

import (
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCloudflareOriginCARootCertificate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudflareOriginCARootCertConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCloudflareOriginCARootCert("data.cloudflare_origin_ca_root_certificate.ecc"),
					testAccCloudflareOriginCARootCert("data.cloudflare_origin_ca_root_certificate.rsa"),
				),
			},
		},
	})
}

func testAccCloudflareOriginCARootCert(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		cert := r.Primary.Attributes["cert_pem"]

		p, _ := pem.Decode([]byte(cert))
		if p == nil || p.Type != "CERTIFICATE" {
			return fmt.Errorf("invalid certificate: %s", cert)
		}

		return nil
	}
}

const testAccCloudflareOriginCARootCertConfig = `
data "cloudflare_origin_ca_root_certificate" "rsa" {
	algorithm = "RSA"
}

data "cloudflare_origin_ca_root_certificate" "ecc" {
	algorithm = "ecc"
}
`
