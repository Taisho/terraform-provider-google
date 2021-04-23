// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceMemcacheInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceMemcacheInstanceCreate,
		Read:   resourceMemcacheInstanceRead,
		Update: resourceMemcacheInstanceUpdate,
		Delete: resourceMemcacheInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMemcacheInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource name of the instance.`,
			},
			"node_config": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `Configuration for memcache nodes.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_count": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `Number of CPUs per node.`,
						},
						"memory_size_mb": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `Memory size in Mebibytes for each memcache node.`,
						},
					},
				},
			},
			"node_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Number of nodes in the memcache instance.`,
			},
			"authorized_network": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `The full name of the GCE network to connect the instance to.  If not provided,
'default' will be used.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `A user-visible name for the instance.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Resource labels to represent user-provided metadata.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"memcache_parameters": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `User-specified parameters for this memcache instance.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"params": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: `User-defined set of parameters to use in the memcache process.`,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `This is a unique ID associated with this set of parameters.`,
						},
					},
				},
			},
			"memcache_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"MEMCACHE_1_5", ""}, false),
				Description: `The major version of Memcached software. If not provided, latest supported version will be used.
Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
determined by our system based on the latest supported minor version. Default value: "MEMCACHE_1_5" Possible values: ["MEMCACHE_1_5"]`,
				Default: "MEMCACHE_1_5",
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the Memcache instance. If it is not provided, the provider region is used.`,
			},
			"zones": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `Zones where memcache nodes should be provisioned.  If not
provided, all zones will be used.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"discovery_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Endpoint for Discovery API`,
			},
			"memcache_full_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The full version of memcached server running on this instance.`,
			},
			"memcache_nodes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Additional information about the instance state, if available.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Hostname or IP address of the Memcached node used by the clients to connect to the Memcached server on this node.`,
						},
						"node_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Identifier of the Memcached node. The node id does not include project or location like the Memcached instance name.`,
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `The port number of the Memcached server on this node.`,
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Current state of the Memcached node.`,
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Location (GCP Zone) for the Memcached node.`,
						},
					},
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceMemcacheInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandMemcacheInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandMemcacheInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	zonesProp, err := expandMemcacheInstanceZones(d.Get("zones"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zones"); !isEmptyValue(reflect.ValueOf(zonesProp)) && (ok || !reflect.DeepEqual(v, zonesProp)) {
		obj["zones"] = zonesProp
	}
	authorizedNetworkProp, err := expandMemcacheInstanceAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("authorized_network"); !isEmptyValue(reflect.ValueOf(authorizedNetworkProp)) && (ok || !reflect.DeepEqual(v, authorizedNetworkProp)) {
		obj["authorizedNetwork"] = authorizedNetworkProp
	}
	nodeCountProp, err := expandMemcacheInstanceNodeCount(d.Get("node_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_count"); !isEmptyValue(reflect.ValueOf(nodeCountProp)) && (ok || !reflect.DeepEqual(v, nodeCountProp)) {
		obj["nodeCount"] = nodeCountProp
	}
	memcacheVersionProp, err := expandMemcacheInstanceMemcacheVersion(d.Get("memcache_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("memcache_version"); !isEmptyValue(reflect.ValueOf(memcacheVersionProp)) && (ok || !reflect.DeepEqual(v, memcacheVersionProp)) {
		obj["memcacheVersion"] = memcacheVersionProp
	}
	nodeConfigProp, err := expandMemcacheInstanceNodeConfig(d.Get("node_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_config"); !isEmptyValue(reflect.ValueOf(nodeConfigProp)) && (ok || !reflect.DeepEqual(v, nodeConfigProp)) {
		obj["nodeConfig"] = nodeConfigProp
	}
	parametersProp, err := expandMemcacheInstanceMemcacheParameters(d.Get("memcache_parameters"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("memcache_parameters"); !isEmptyValue(reflect.ValueOf(parametersProp)) && (ok || !reflect.DeepEqual(v, parametersProp)) {
		obj["parameters"] = parametersProp
	}

	url, err := replaceVars(d, config, "{{MemcacheBasePath}}projects/{{project}}/locations/{{region}}/instances?instanceId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = memcacheOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Instance", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Instance: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceMemcacheInstanceRead(d, meta)
}

func resourceMemcacheInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{MemcacheBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("MemcacheInstance %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	if err := d.Set("display_name", flattenMemcacheInstanceDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("memcache_nodes", flattenMemcacheInstanceMemcacheNodes(res["memcacheNodes"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("create_time", flattenMemcacheInstanceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("discovery_endpoint", flattenMemcacheInstanceDiscoveryEndpoint(res["discoveryEndpoint"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("labels", flattenMemcacheInstanceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("memcache_full_version", flattenMemcacheInstanceMemcacheFullVersion(res["memcacheFullVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("zones", flattenMemcacheInstanceZones(res["zones"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("authorized_network", flattenMemcacheInstanceAuthorizedNetwork(res["authorizedNetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("node_count", flattenMemcacheInstanceNodeCount(res["nodeCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("memcache_version", flattenMemcacheInstanceMemcacheVersion(res["memcacheVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("node_config", flattenMemcacheInstanceNodeConfig(res["nodeConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("memcache_parameters", flattenMemcacheInstanceMemcacheParameters(res["parameters"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceMemcacheInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandMemcacheInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	labelsProp, err := expandMemcacheInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	nodeCountProp, err := expandMemcacheInstanceNodeCount(d.Get("node_count"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_count"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nodeCountProp)) {
		obj["nodeCount"] = nodeCountProp
	}
	memcacheVersionProp, err := expandMemcacheInstanceMemcacheVersion(d.Get("memcache_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("memcache_version"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, memcacheVersionProp)) {
		obj["memcacheVersion"] = memcacheVersionProp
	}

	url, err := replaceVars(d, config, "{{MemcacheBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Instance %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("node_count") {
		updateMask = append(updateMask, "nodeCount")
	}

	if d.HasChange("memcache_version") {
		updateMask = append(updateMask, "memcacheVersion")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Instance %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Instance %q: %#v", d.Id(), res)
	}

	err = memcacheOperationWaitTime(
		config, res, project, "Updating Instance", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceMemcacheInstanceRead(d, meta)
}

func resourceMemcacheInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{MemcacheBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Instance %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Instance")
	}

	err = memcacheOperationWaitTime(
		config, res, project, "Deleting Instance", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceMemcacheInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/instances/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenMemcacheInstanceDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMemcacheInstanceMemcacheNodes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"node_id": flattenMemcacheInstanceMemcacheNodesNodeId(original["nodeId"], d, config),
			"zone":    flattenMemcacheInstanceMemcacheNodesZone(original["zone"], d, config),
			"port":    flattenMemcacheInstanceMemcacheNodesPort(original["port"], d, config),
			"host":    flattenMemcacheInstanceMemcacheNodesHost(original["host"], d, config),
			"state":   flattenMemcacheInstanceMemcacheNodesState(original["state"], d, config),
		})
	}
	return transformed
}
func flattenMemcacheInstanceMemcacheNodesNodeId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMemcacheInstanceMemcacheNodesZone(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMemcacheInstanceMemcacheNodesPort(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenMemcacheInstanceMemcacheNodesHost(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMemcacheInstanceMemcacheNodesState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMemcacheInstanceCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMemcacheInstanceDiscoveryEndpoint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMemcacheInstanceLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMemcacheInstanceMemcacheFullVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMemcacheInstanceZones(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenMemcacheInstanceAuthorizedNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMemcacheInstanceNodeCount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenMemcacheInstanceMemcacheVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMemcacheInstanceNodeConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["cpu_count"] =
		flattenMemcacheInstanceNodeConfigCpuCount(original["cpuCount"], d, config)
	transformed["memory_size_mb"] =
		flattenMemcacheInstanceNodeConfigMemorySizeMb(original["memorySizeMb"], d, config)
	return []interface{}{transformed}
}
func flattenMemcacheInstanceNodeConfigCpuCount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenMemcacheInstanceNodeConfigMemorySizeMb(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenMemcacheInstanceMemcacheParameters(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["id"] =
		flattenMemcacheInstanceMemcacheParametersId(original["id"], d, config)
	transformed["params"] =
		flattenMemcacheInstanceMemcacheParametersParams(original["params"], d, config)
	return []interface{}{transformed}
}
func flattenMemcacheInstanceMemcacheParametersId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenMemcacheInstanceMemcacheParametersParams(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandMemcacheInstanceDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMemcacheInstanceZones(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandMemcacheInstanceAuthorizedNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceNodeCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMemcacheVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceNodeConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCpuCount, err := expandMemcacheInstanceNodeConfigCpuCount(original["cpu_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuCount); val.IsValid() && !isEmptyValue(val) {
		transformed["cpuCount"] = transformedCpuCount
	}

	transformedMemorySizeMb, err := expandMemcacheInstanceNodeConfigMemorySizeMb(original["memory_size_mb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMemorySizeMb); val.IsValid() && !isEmptyValue(val) {
		transformed["memorySizeMb"] = transformedMemorySizeMb
	}

	return transformed, nil
}

func expandMemcacheInstanceNodeConfigCpuCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceNodeConfigMemorySizeMb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMemcacheParameters(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandMemcacheInstanceMemcacheParametersId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
		transformed["id"] = transformedId
	}

	transformedParams, err := expandMemcacheInstanceMemcacheParametersParams(original["params"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedParams); val.IsValid() && !isEmptyValue(val) {
		transformed["params"] = transformedParams
	}

	return transformed, nil
}

func expandMemcacheInstanceMemcacheParametersId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMemcacheInstanceMemcacheParametersParams(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}