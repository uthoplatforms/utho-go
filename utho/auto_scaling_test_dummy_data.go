package utho

const dummyCreateAutoScaling = `{
    "name": "Auto-scaling-Jz8hceLN.utho",
    "os_disk_size": 80,
    "dcslug": "inmumbaizone2",
    "minsize": "1",
    "maxsize": "2",
    "desiredsize": "1",
    "planid": "10045",
    "planname": "basic",
    "instance_templateid": "none",
    "public_ip_enabled": true,
    "vpc": "f1dd58f1-1bfa-46ef-8b94-f69f312c0245",
    "load_balancers": "",
    "security_groups": "23432630",
    "policies": [
        {
            "name": "Policy-16H2jh",
            "type": "cpu",
            "compare": "above",
            "value": "90",
            "adjust": "1",
            "period": "5m",
            "cooldown": "300"
        }
    ],
    "schedules": [
        {
            "name": "Schedule-nVaZ9y",
            "desiredsize": "2",
            "start_date": "2024-05-11T17:05:01.709Z",
            "selectedTime": "20:05",
            "selectedDate": "2024-05-11"
        }
    ],
    "stackid": "6669341",
    "stackimage": "ubuntu-22.04-x86_64",
    "target_groups": "234567"
}`

const dummyCreateAutoScalingResponse = `{
	"id": 11111,
    "status": "success",
    "message": "success"
}`

const dummyReadAutoScalingRes = `{
	"id": "11111",
	"userid": "65432",
	"name": "Auto-scaling-Jz8hceLN.utho",
	"dcslug": "inmumbaizone2",
	"minsize": "1",
	"maxsize": "2",
	"desiredsize": "1",
	"planid": "10045",
	"planname": "basic",
	"instance_templateid": "none",
	"image": "ubuntu-22.04-x86_64",
	"image_name": "",
	"snapshotid": "0",
	"status": "Active",
	"created_at": "2024-05-11 22:35:06",
	"suspended_at": "0000-00-00 00:00:00",
	"stopped_at": "0000-00-00 00:00:00",
	"started_at": "0000-00-00 00:00:00",
	"deleted_at": "0000-00-00 00:00:00",
	"public_ip_enabled": "1",
	"vpc": [
		{
			"total": 254,
			"available": 248,
			"network": "10.210.100.0",
			"name": "test",
			"size": "24",
			"dcslug": "innoida",
			"dclocation": {
				"dccc": "in",
				"location": "Delhi (Noida)"
			},
			"is_default": null,
			"resources": [
				{
					"type": "cloud",
					"id": "1277623",
					"name": "cloudserver-soLSCBWM.mhc",
					"ip": "10.210.100.6"
				},
				{
					"type": "cloud",
					"id": "1277092",
					"name": "DBaaS-188417-2",
					"ip": "10.210.100.3"
				},
				{
					"type": "cloud",
					"id": "1277539",
					"name": "mks-pool-mPly8wG9-node-fpuvw",
					"ip": "10.210.100.4"
				},
				{
					"type": "cloud",
					"id": "1277538",
					"name": "MyK8S-ESF4oaIr-ie0x",
					"ip": "10.210.100.1"
				},
				{
					"type": "cloud",
					"id": "1277087",
					"name": "DBaaS-188415-2",
					"ip": "10.210.100.2"
				},
				{
					"type": "cloud",
					"id": "1277540",
					"name": "mks-pool-D1rtwPqs-node-rnsco",
					"ip": "10.210.100.5"
				}
			]
		}
	],
	"cooldown_till": "",
	"load_balancers": [
		{
			"lbid": "44444",
			"name": "hqwe",
			"ip": "103.111.111.111"
		}
	],
	"target_groups": [
		{
			"id": "666666",
			"name": "example1z2q",
			"protocol": null,
			"port": "12"
		}
	],
	"security_groups": [
		{
			"id": "55555",
			"name": "new-server"
		}
	],
	"backupid": "0",
	"stack": "6669341",
	"stack_fields": "null",
	"instances": [
		{
			"cloudid": "1277770",
			"hostname": "Auto-scaling-Jz8hceLN.utho-instance-yrlsz",
			"created_at": "2024-05-11 22:35:24",
			"ip": "103.150.136.128",
			"status": "Active"
		}
	],
	"policies": [
		{
			"id": "22222",
			"userid": "12412",
			"product": "scaling_group",
			"productid": "1241",
			"groupid": "12412",
			"name": "Policy-16H2jh",
			"type": "cpu",
			"adjust": "1",
			"period": "5m",
			"cooldown": "300",
			"cooldown_till": null,
			"compare": "above",
			"value": "90",
			"alert_id": "",
			"status": "1",
			"kubernetes_id": "0",
			"kubernetes_nodepool": null,
			"cloudid": "0",
			"maxsize": "0",
			"minsize": "0"
		}
	],
	"schedules": [
		{
			"id": "33333",
			"groupid": "123214",
			"name": "Schedule-nVaZ9y",
			"desiredsize": "2",
			"recurrence": "",
			"start_date": "2024-05-11 17:05:02",
			"status": "1",
			"timezone": "IST"
		}
	],
	"deleted_instances": [],
	"dclocation": {
		"location": "Mumbai",
		"country": "India",
		"dc": "inmumbaizone2",
		"dccc": "in"
	},
	"plan": {
		"planid": "10045",
		"ram": "4096",
		"cpu": "2",
		"disk": "80",
		"bandwidth": "1000",
		"dedicated_vcore": "0"
	}
}`

const dummyAutoScalingServerRes = `{
    "groups": [` + dummyReadAutoScalingRes + `]
}`

const dummyListAutoScalingRes = `[` + dummyReadAutoScalingRes + `]`

// Auto Scaling Policy
const dummyCreateAutoScalingPolicy = `{
    "name": "Policy-RmcXd9",
    "type": "cpu",
    "compare": "above",
    "value": "50",
    "adjust": "1",
    "period": "5m",
    "cooldown": "300",
    "product": "scaling_group",
    "productid": "23492456"
}`

const dummyReadAutoScalingPolicyRes = `{
	"id": "22222",
	"userid": "12412",
	"product": "scaling_group",
	"productid": "1241",
	"groupid": "12412",
	"name": "Policy-16H2jh",
	"type": "cpu",
	"adjust": "1",
	"period": "5m",
	"cooldown": "300",
	"cooldown_till": null,
	"compare": "above",
	"value": "90",
	"alert_id": "",
	"status": "1",
	"kubernetes_id": "0",
	"kubernetes_nodepool": null,
	"cloudid": "0",
	"maxsize": "0",
	"minsize": "0"
}`

const dummyListAutoScalingPolicyRes = `[` + dummyReadAutoScalingPolicyRes + `]`

// Auto Scaling Schedule
const dummyCreateAutoScalingSchedule = `{
    "name": "wqf",
    "desiredsize": "1",
    "recurrence": "Once",
    "start_date": "2024-05-11 20:05:00"
}`

const dummyReadAutoScalingScheduleRes = `{
	"id": "33333",
	"groupid": "123214",
	"name": "Schedule-nVaZ9y",
	"desiredsize": "2",
	"recurrence": "",
	"start_date": "2024-05-11 17:05:02",
	"status": "1",
	"timezone": "IST"
}`

const dummyListAutoScalingScheduleRes = `[` + dummyReadAutoScalingScheduleRes + `]`

// Auto Scaling Loadbalancer
const dummyReadAutoScalingLoadbalancerRes = `{
	"lbid": "44444",
	"name": "hqwe",
	"ip": "103.111.111.111"
}`

const dummyListAutoScalingLoadbalancerRes = `[` + dummyReadAutoScalingLoadbalancerRes + `]`

// Auto Scaling SecurityGroup
const dummyReadAutoScalingSecurityGroupRes = `{
	"id": "55555",
	"name": "new-server"
}`

const dummyListAutoScalingSecurityGroupRes = `[` + dummyReadAutoScalingSecurityGroupRes + `]`

// Auto Scaling Targetgroup
const dummyReadAutoScalingTargetgroupRes = `{
	"id": "666666",
	"name": "example1z2q",
	"protocol": null,
	"port": "12"
}`

const dummyListAutoScalingTargetgroupRes = `[` + dummyReadAutoScalingTargetgroupRes + `]`
