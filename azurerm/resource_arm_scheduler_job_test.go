package azurerm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

// check: start time
// check recurring basic -> recurring each type?
// check: basic everything -> complete everything?
// check : base + web + error action
// check : base + web + auth_basic
// check : base + web + auth_cert
// check : base + web + auth_ad

func TestAccAzureRMSchedulerJob_web_basic(t *testing.T) {
	ri := acctest.RandInt()
	resourceName := "azurerm_scheduler_job.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSchedulerJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMSchedulerJob_base(ri, testLocation(),
					testAccAzureRMSchedulerJob_block_actionWeb_basic(),
					"", "", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					checkAccAzureRMSchedulerJob_base(resourceName),
					checkAccAzureRMSchedulerJob_web_basic(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureRMSchedulerJob_web_put(t *testing.T) {
	ri := acctest.RandInt()
	resourceName := "azurerm_scheduler_job.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSchedulerJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMSchedulerJob_base(ri, testLocation(),
					testAccAzureRMSchedulerJob_block_actionWeb_basic(),
					"", "", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					checkAccAzureRMSchedulerJob_base(resourceName),
					checkAccAzureRMSchedulerJob_web_basic(resourceName),
				),
			},
			{
				Config: testAccAzureRMSchedulerJob_base(ri, testLocation(),
					testAccAzureRMSchedulerJob_block_actionWeb_put(),
					"", "", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					checkAccAzureRMSchedulerJob_base(resourceName),
					checkAccAzureRMSchedulerJob_web_put(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureRMSchedulerJob_web_recurring(t *testing.T) {
	ri := acctest.RandInt()
	resourceName := "azurerm_scheduler_job.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSchedulerJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMSchedulerJob_base(ri, testLocation(),
					testAccAzureRMSchedulerJob_block_actionWeb_basic(),
					testAccAzureRMSchedulerJob_block_recurrence_basic(),
					"", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					checkAccAzureRMSchedulerJob_base(resourceName),
					checkAccAzureRMSchedulerJob_web_basic(resourceName),
					checkAccAzureRMSchedulerJob_recurrence_basic(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureRMSchedulerJob_web_recurringDaily(t *testing.T) {
	ri := acctest.RandInt()
	resourceName := "azurerm_scheduler_job.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSchedulerJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMSchedulerJob_base(ri, testLocation(),
					testAccAzureRMSchedulerJob_block_actionWeb_basic(),
					testAccAzureRMSchedulerJob_block_recurrence_daily(),
					"", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					checkAccAzureRMSchedulerJob_base(resourceName),
					checkAccAzureRMSchedulerJob_web_basic(resourceName),
					checkAccAzureRMSchedulerJob_recurrence_daily(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureRMSchedulerJob_web_recurringWeekly(t *testing.T) {
	ri := acctest.RandInt()
	resourceName := "azurerm_scheduler_job.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSchedulerJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMSchedulerJob_base(ri, testLocation(),
					testAccAzureRMSchedulerJob_block_actionWeb_basic(),
					testAccAzureRMSchedulerJob_block_recurrence_weekly(),
					"", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					checkAccAzureRMSchedulerJob_base(resourceName),
					checkAccAzureRMSchedulerJob_web_basic(resourceName),
					checkAccAzureRMSchedulerJob_recurrence_weekly(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureRMSchedulerJob_web_retry(t *testing.T) {
	ri := acctest.RandInt()
	resourceName := "azurerm_scheduler_job.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSchedulerJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMSchedulerJob_base(ri, testLocation(),
					testAccAzureRMSchedulerJob_block_actionWeb_basic(),
					testAccAzureRMSchedulerJob_block_retry_empty(),
					"", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					checkAccAzureRMSchedulerJob_base(resourceName),
					checkAccAzureRMSchedulerJob_web_basic(resourceName),
					checkAccAzureRMSchedulerJob_retry_empty(resourceName),
				),
			},
			{
				Config: testAccAzureRMSchedulerJob_base(ri, testLocation(),
					testAccAzureRMSchedulerJob_block_actionWeb_basic(),
					testAccAzureRMSchedulerJob_block_retry_complete(),
					"", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					checkAccAzureRMSchedulerJob_base(resourceName),
					checkAccAzureRMSchedulerJob_web_basic(resourceName),
					checkAccAzureRMSchedulerJob_retry_complete(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureRMSchedulerJob_web_basic_onceToRecurring(t *testing.T) {
	ri := acctest.RandInt()
	resourceName := "azurerm_scheduler_job.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureRMSchedulerJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMSchedulerJob_base(ri, testLocation(),
					testAccAzureRMSchedulerJob_block_actionWeb_basic(),
					"", "", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					checkAccAzureRMSchedulerJob_base(resourceName),
					checkAccAzureRMSchedulerJob_web_basic(resourceName),
				),
			},
			{
				Config: testAccAzureRMSchedulerJob_base(ri, testLocation(),
					testAccAzureRMSchedulerJob_block_actionWeb_basic(),
					"", "", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					checkAccAzureRMSchedulerJob_base(resourceName),
					checkAccAzureRMSchedulerJob_web_basic(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testCheckAzureRMSchedulerJobDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azurerm_scheduler_job.test" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]
		jobCollection := rs.Primary.Attributes["job_collection_name"]

		client := testAccProvider.Meta().(*ArmClient).schedulerJobsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		resp, err := client.Get(ctx, resourceGroup, jobCollection, name)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Scheduler Job Collection still exists:\n%#v", resp)
	}

	return nil
}

func testCheckAzureRMSchedulerJobExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup, hasResourceGroup := rs.Primary.Attributes["resource_group_name"]
		jobCollection := rs.Primary.Attributes["job_collection_name"]

		if !hasResourceGroup {
			return fmt.Errorf("Bad: no resource group found in state for Scheduler Job: %q", name)
		}

		client := testAccProvider.Meta().(*ArmClient).schedulerJobsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		resp, err := client.Get(ctx, resourceGroup, jobCollection, name)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Scheduler Job  %q (resource group: %q) was not found: %+v", name, resourceGroup, err)
			}

			return fmt.Errorf("Bad: Get on schedulerJobsClient: %+v", err)
		}

		return nil
	}
}

func testAccAzureRMSchedulerJob_base(rInt int, location, block1, block2, block3, block4 string) string {
	return fmt.Sprintf(` 
resource "azurerm_resource_group" "rg" { 
  name     = "acctestRG-%[1]d" 
  location = "%[2]s" 
} 
 
resource "azurerm_scheduler_job_collection" "jc" {
    name                = "acctestRG-%[1]d-job_collection"
    location            = "${azurerm_resource_group.rg.location}"
    resource_group_name = "${azurerm_resource_group.rg.name}"
    sku                 = "standard"
}

resource "azurerm_scheduler_job" "test" {
    name                = "acctestRG-%[1]d-job"
    resource_group_name = "${azurerm_resource_group.rg.name}"
    job_collection_name = "${azurerm_scheduler_job_collection.jc.name}"

    %[3]s

    %[4]s

    %[5]s

    %[6]s
} 
`, rInt, location, block1, block2, block3, block4)
}
func checkAccAzureRMSchedulerJob_base(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		testCheckAzureRMSchedulerJobExists(resourceName),
		resource.TestCheckResourceAttrSet(resourceName, "name"),
		resource.TestCheckResourceAttrSet(resourceName, "resource_group_name"),
		resource.TestCheckResourceAttrSet(resourceName, "job_collection_name"),
		resource.TestCheckResourceAttr(resourceName, "state", "enabled"),
	)
}

func testAccAzureRMSchedulerJob_block_actionWeb_basic() string {
	return ` 
  action_web {
    url = "http://this.get.url.fails"
  } 
`
}
func checkAccAzureRMSchedulerJob_web_basic(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "action_web.0.url", "http://this.get.url.fails"),
	)
}

func testAccAzureRMSchedulerJob_block_actionWeb_put() string {
	return `
  action_web {
    url    = "http://this.put.url.fails"
    method = "put"
    body   = "this is some text"
    headers = {
      Content-Type = "text"
	}
  } 
`
}
func checkAccAzureRMSchedulerJob_web_put(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "action_web.0.url", "http://this.put.url.fails"),
		resource.TestCheckResourceAttr(resourceName, "action_web.0.method", "put"),
		resource.TestCheckResourceAttr(resourceName, "action_web.0.body", "this is some text"),
		resource.TestCheckResourceAttrSet(resourceName, "headers"),
	)
}

func testAccAzureRMSchedulerJob_block_recurrence_basic() string {
	return ` 
  recurrence {
    frequency  = "minute"
    interval   = 5
    count      = 10
  //end_time  = "2019-07-17T07:07:07-07:00"
  } 
`
}
func checkAccAzureRMSchedulerJob_recurrence_basic(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.frequency", "minute"),
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.interval", "5"),
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.count", "10"),
	)
}

func testAccAzureRMSchedulerJob_block_recurrence_daily() string {
	return ` 
  recurrence {
    frequency = "day"
    count     = 100 
    hours     = [0,12]
    minutes   = [0,15,30,45] 
  } 
`
}
func checkAccAzureRMSchedulerJob_recurrence_daily(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.frequency", "day"),
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.count", "100"),
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.hours.#", "2"),
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.minutes.#", "4"),
	)
}

func testAccAzureRMSchedulerJob_block_recurrence_weekly() string {
	return ` 
  recurrence {
     frequency    = "week"
     count        = 100 
     week_days = ["Sunday", "Saturday"] 
  } 
`
}
func checkAccAzureRMSchedulerJob_recurrence_weekly(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.frequency", "week"),
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.count", "100"),
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.week_days.#", "2"),
	)
}

func testAccAzureRMSchedulerJob_block_recurrence_monthly() string {
	return ` 
  recurrence {
    frequency  = "month"
    count      = 100 
    month_days = [-11,-1,1,11]
  } 
`
}

func checkAccAzureRMSchedulerJob_recurrence_monthly(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.frequency", "month"),
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.count", "100"),
		resource.TestCheckResourceAttr(resourceName, "recurrence.0.month_days.#", "4"),
	)
}

func testAccAzureRMSchedulerJob_block_retry_empty() string {
	return ` 
  retry { 
  } 
`
}
func checkAccAzureRMSchedulerJob_retry_empty(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "retry.0.interval", "00:00:30"),
		resource.TestCheckResourceAttr(resourceName, "retry.0.count", "4"),
	)
}

func testAccAzureRMSchedulerJob_block_retry_complete() string {
	return ` 
  retry { 
    interval = "00:05:00" //5 min
    count    =  10
  } 
`
}
func checkAccAzureRMSchedulerJob_retry_complete(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "retry.0.interval", "00:05:00"),
		resource.TestCheckResourceAttr(resourceName, "retry.0.count", "10"),
	)
}
