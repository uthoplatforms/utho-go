package utho

const dummyReadKubernetesRes = `{
	"id": "11111",
	"cloudid": "12345",
	"created_at": "2024-05-06 21:22:05",
	"dcslug": "inmumbaizone2",
	"ref_id": "11111",
	"nodepool": "",
	"hostname": "MyK8S-ESF4oaIr-ie0x",
	"ram": "4096",
	"cpu": "2",
	"disksize": "80",
	"app_status": "Pending",
	"ip": "103.111.111.111",
	"powerstatus": "online",
	"dclocation": {
		"location": "Mumbai",
		"country": "India",
		"dc": "inmumbaizone2",
		"dccc": "in"
	},
	"status": "Active",
	"worker_count": "2",
	"load_balancers": [
		{
			"lbid": "22222",
			"name": "hgfd",
			"ip": "103.11.11.11"
		}
	],
	"target_groups": [
		{
			"id": "33333",
			"name": "d12d4",
			"protocol": null,
			"port": "1"
		}
	],
	"security_groups": [
		{
			"id": "44444",
			"name": "new-server"
		}
	]
}`

const dummyKubernetesServerRes = `{
    "k8s": [` + dummyReadKubernetesRes + `]
}`

const dummyListKubernetesRes = `[` + dummyReadKubernetesRes + `]`

const dummyReadKubernetesLoadbalancerRes = `{
    "lbid": "22222",
    "name": "hgfd",
    "ip": "103.11.11.11"
}`

const dummyListKubernetesLoadbalancerRes = `[` + dummyReadKubernetesLoadbalancerRes + `]`

const dummyReadKubernetesSecurityGroupRes = `{
    "id": "44444",
    "name": "new-server"
}`

const dummyListKubernetesSecurityGroupRes = `[` + dummyReadKubernetesSecurityGroupRes + `]`

const dummyReadKubernetesTargetgroupRes = `{
	"id": "33333",
	"name": "d12d4",
	"protocol": null,
	"port": "1"
}`

const dummyListKubernetesTargetgroupRes = `[` + dummyReadKubernetesTargetgroupRes + `]`
