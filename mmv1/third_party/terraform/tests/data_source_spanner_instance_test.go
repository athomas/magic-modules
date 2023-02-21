package google_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSpannerInstance_basic(t *testing.T) {
	// Randomness from spanner instance
	provider.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": google.RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { TestAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckSpannerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSpannerInstanceBasic(context),
				Check: resource.ComposeTestCheckFunc(
					CheckDataSourceStateMatchesResourceState("data.google_spanner_instance.foo", "google_spanner_instance.bar"),
				),
			},
		},
	})
}

func testAccDataSourceSpannerInstanceBasic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_spanner_instance" "bar" {
	config       = "regional-us-central1"
	display_name = "Test Spanner Instance"
	num_nodes    = 2
	labels = {
		"foo" = "bar"
	}
}

data "google_spanner_instance" "foo" {
	name = google_spanner_instance.bar.name
}
`, context)
}