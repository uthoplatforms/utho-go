package utho

const dummyReadKubernetesRes = `{
  "info": {
    "cluster": {
      "id": "11111",
      "version": "v1.20.4", 
      "label": "MyK8SCluster",
      "endpoint": "https://myk8s-endpoint.example.com",
      "dcslug": "inmumbaizone2",
      "auto_upgrade": "false",
      "surge_upgrade": "false",
      "ipv4": "103.111.111.111",
      "cluster_subnet": "192.168.0.0/16",
      "service_subnet": "10.0.0.0/16",
      "tags": "tag1,tag2",
      "created_at": "2024-05-06 21:22:05",
      "updated_at": "",
      "deleted_at": "",
      "status": "Active",
      "nodepools": "default",
      "vpc": "vpc-id-123",
      "public_ip_enabled": "true",
      "load_balancers": "load_balancer_id",
      "security_groups": "security_group_id",
      "target_groups": "target_group_id",
      "userid": "123456",
      "powerstatus": "online",
      "dclocation": {
        "location": "Mumbai",
        "country": "India",
        "dc": "inmumbaizone2",
        "dccc": "in"
      }
    },
    "master": {
      "cloudid": "12345",
      "hostname": "MyK8S-ESF4oaIr-ie0x",
      "ram": "4096",
      "cpu": "2",
      "cost": "100",
      "disksize": "80",
      "app_status": "Pending",
      "dcslug": "inmumbaizone2",
      "planid": "plan-id-001",
      "ip": "103.111.111.111",
      "private_network": {
        "ip": "10.0.1.1",
        "vpc": "vpc-id-123",
        "vpc_network": "10.0.0.0/16"
      }
    }
  },
  "vpc": [
    {
      "id": "vpc-id-123",
      "vpc_network": "10.0.0.0/16"
    }
  ],
  "nodepools": {
    "default": {
      "size": "medium",
      "cost": 200,
      "planid": "plan-id-002",
      "count": "2",
      "policies": [],
      "workers": [
        {
          "cloudid": "11111",
          "nodepool": "default",
          "hostname": "worker-node-1",
          "ram": "4096",
          "cost": "50",
          "cpu": "2",
          "disksize": "80",
          "app_status": "Active",
          "ip": "103.111.112.112",
          "planid": "plan-id-003",
          "status": "Active",
          "private_network": {
            "ip": "10.0.1.2",
            "vpc": "vpc-id-123",
            "vpc_network": "10.0.0.0/16"
          }
        }
      ]
    }
  },
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
  ],
  "rcode": "0"
}
`

const dummyKubernetesServerRes = `{
  "info": {
    "cluster": {
      "id": "11111",
      "version": "v1.20.4", 
      "label": "MyK8SCluster",
      "endpoint": "https://myk8s-endpoint.example.com",
      "dcslug": "inmumbaizone2",
      "auto_upgrade": "false",
      "surge_upgrade": "false",
      "ipv4": "103.111.111.111",
      "cluster_subnet": "192.168.0.0/16",
      "service_subnet": "10.0.0.0/16",
      "tags": "tag1,tag2",
      "created_at": "2024-05-06 21:22:05",
      "updated_at": "",
      "deleted_at": "",
      "status": "Active",
      "nodepools": "default",
      "vpc": "vpc-id-123",
      "public_ip_enabled": "true",
      "load_balancers": "load_balancer_id",
      "security_groups": "security_group_id",
      "target_groups": "target_group_id",
      "userid": "123456",
      "powerstatus": "online",
      "dclocation": {
        "location": "Mumbai",
        "country": "India",
        "dc": "inmumbaizone2",
        "dccc": "in"
      }
    },
    "master": {
      "cloudid": "12345",
      "hostname": "MyK8S-ESF4oaIr-ie0x",
      "ram": "4096",
      "cpu": "2",
      "cost": "100",
      "disksize": "80",
      "app_status": "Pending",
      "dcslug": "inmumbaizone2",
      "planid": "plan-id-001",
      "ip": "103.111.111.111",
      "private_network": {
        "ip": "10.0.1.1",
        "vpc": "vpc-id-123",
        "vpc_network": "10.0.0.0/16"
      }
    }
  },
  "vpc": [
    {
      "id": "vpc-id-123",
      "vpc_network": "10.0.0.0/16"
    }
  ],
  "nodepools": {
    "default": {
      "size": "medium",
      "cost": 200,
      "planid": "plan-id-002",
      "count": "2",
      "policies": [],
      "workers": [
        {
          "cloudid": "11111",
          "nodepool": "default",
          "hostname": "worker-node-1",
          "ram": "4096",
          "cost": "50",
          "cpu": "2",
          "disksize": "80",
          "app_status": "Active",
          "ip": "103.111.112.112",
          "planid": "plan-id-003",
          "status": "Active",
          "private_network": {
            "ip": "10.0.1.2",
            "vpc": "vpc-id-123",
            "vpc_network": "10.0.0.0/16"
          }
        }
      ]
    }
  },
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
  ],
  "rcode": "0"
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
