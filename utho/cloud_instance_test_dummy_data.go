package utho

const dummyCreateCloudInstanceRequestJson = `{
    "dcslug": "innoida",
    "image": "ubuntu-18.10-x86_64",
    "planid": "10045",
    "billingcycle": "hourly",
    enable_publicip: "true"
    cpumodel: "intel"
    "cloud": [
        {
            "hostname": "cloudserver-example-name"
        }
    ]
}`

const dummyCreateCloudInstanceResponseJson = `{
    "status": "success",
    "cloudid": "1111111",
    "message": "Cloud Server deploy in process and as soon it get ready to use system will send you login detail over the email.",
    "password": "qwertuioo@111",
    "ipv4": "210.210.210.210"
}`

const dummyReadCloudInstanceServerRes = `{
    "cloud": [
        {
            "cloudid": "1111111",
            "hostname": "cloudserver-fiGwdwnGPv.mhc",
            "cpu": "2",
            "ram": "4096",
            "managed_os": null,
            "managed_full": null,
            "managed_onetime": null,
            "plan_disksize": 80,
            "disksize": 80,
            "ha": "0",
            "status": "Active",
            "iso": null,
            "ip": "201.201.201.201",
            "billingcycle": "hourly",
            "cost": 19.93,
            "vmcost": 19.93,
            "imagecost": 0,
            "backupcost": 0,
            "hourlycost": 0.03,
            "cloudhourlycost": 0.03,
            "imagehourlycost": 0,
            "backuphourlycost": 0,
            "creditrequired": 0.5,
            "creditreserved": 0,
            "nextinvoiceamount": 0.5,
            "nextinvoicehours": "17",
            "consolepassword": "Ofi4G",
            "powerstatus": "Running",
            "created_at": "2024-05-07 23:03:25",
            "updated_at": "0000-00-00 00:00:00",
            "nextduedate": "2024-05-07",
            "bandwidth": "1000",
            "bandwidth_used": 0,
            "bandwidth_free": 1000,
            "features": {
                "backups": "0"
            },
            "image": {
                "name": "Ubuntu 18.04",
                "distribution": "ubuntu",
                "version": "18.04",
                "image": "ubuntu-18.10-x86_64",
                "cost": "0"
            },
            "dclocation": {
                "location": "Delhi (Noida)",
                "country": "India",
                "dc": "innoida",
                "dccc": "in"
            },
            "v4": {
                "ip_address": "111.111.111.111",
                "netmask": "111.111.111.111",
                "gateway": "111.111.111.111",
                "type": "public",
                "nat": false,
                "primary": "1",
                "rdns": "1031272912.network.microhost.in"
            },
            "networks": {
                "public": {
                    "v4": [
                        {
                            "ip_address": "111.111.111.111",
                            "netmask": "111.111.111.111",
                            "gateway": "111.111.111.111",
                            "type": "public",
                            "nat": false,
                            "primary": "1",
                            "rdns": "1111111.network.microhost.in"
                        }
                    ]
                },
                "private": {
                    "v4": [
                        {
                            "noip": 0
                        }
                    ]
                }
            },
            "v4private": {
                "noip": 0
            },
            "storages": [
                {
                    "id": "1111111111",
                    "size": 80,
                    "disk_used": "0.00",
                    "disk_free": "0.00",
                    "disk_usedp": "0.00",
                    "created_at": null,
                    "bus": "virtio",
                    "type": "Primary"
                }
            ],
            "storage": {
                "id": "1111111",
                "size": 80,
                "disk_used": "0.00",
                "disk_free": "0.00",
                "disk_usedp": "0.00",
                "created_at": null,
                "bus": "virtio",
                "type": "Primary"
            },
            "disk_used": 0,
            "disk_free": 0,
            "disk_usedp": 0,
            "backups": [],
            "snapshots": [],
            "firewalls": [],
            "gpu_available": "0",
            "gpus": []
        }
    ]
}`

const dummyReadCloudInstanceRes = `{
	"cloudid": "1111111",
	"hostname": "cloudserver-fiGwdwnGPv.mhc",
	"cpu": "2",
	"ram": "4096",
	"managed_os": null,
	"managed_full": null,
	"managed_onetime": null,
	"plan_disksize": 80,
	"disksize": 80,
	"ha": "0",
	"status": "Active",
	"iso": null,
	"ip": "201.201.201.201",
	"billingcycle": "hourly",
	"cost": 19.93,
	"vmcost": 19.93,
	"imagecost": 0,
	"backupcost": 0,
	"hourlycost": 0.03,
	"cloudhourlycost": 0.03,
	"imagehourlycost": 0,
	"backuphourlycost": 0,
	"creditrequired": 0.5,
	"creditreserved": 0,
	"nextinvoiceamount": 0.5,
	"nextinvoicehours": "17",
	"consolepassword": "Ofi4G",
	"powerstatus": "Running",
	"created_at": "2024-05-07 23:03:25",
	"updated_at": "0000-00-00 00:00:00",
	"nextduedate": "2024-05-07",
	"bandwidth": "1000",
	"bandwidth_used": 0,
	"bandwidth_free": 1000,
	"features": {
		"backups": "0"
	},
	"image": {
		"name": "Ubuntu 18.04",
		"distribution": "ubuntu",
		"version": "18.04",
		"image": "ubuntu-18.10-x86_64",
		"cost": "0"
	},
	"dclocation": {
		"location": "Delhi (Noida)",
		"country": "India",
		"dc": "innoida",
		"dccc": "in"
	},
	"v4": {
		"ip_address": "111.111.111.111",
		"netmask": "111.111.111.111",
		"gateway": "111.111.111.111",
		"type": "public",
		"nat": false,
		"primary": "1",
		"rdns": "1031272912.network.microhost.in"
	},
	"networks": {
		"public": {
			"v4": [
				{
					"ip_address": "111.111.111.111",
					"netmask": "111.111.111.111",
					"gateway": "111.111.111.111",
					"type": "public",
					"nat": false,
					"primary": "1",
					"rdns": "1111111.network.microhost.in"
				}
			]
		},
		"private": {
			"v4": [
				{
					"noip": 0
				}
			]
		}
	},
	"v4private": {
		"noip": 0
	},
	"storages": [
		{
			"id": "1111111111",
			"size": 80,
			"disk_used": "0.00",
			"disk_free": "0.00",
			"disk_usedp": "0.00",
			"created_at": null,
			"bus": "virtio",
			"type": "Primary"
		}
	],
	"storage": {
		"id": "1111111",
		"size": 80,
		"disk_used": "0.00",
		"disk_free": "0.00",
		"disk_usedp": "0.00",
		"created_at": null,
		"bus": "virtio",
		"type": "Primary"
	},
	"disk_used": 0,
	"disk_free": 0,
	"disk_usedp": 0,
	"backups": [],
	"snapshots": [],
	"firewalls": [],
	"gpu_available": "0",
	"gpus": []
}`

const dummyListCloudInstanceRes = `{
    "cloud": [
        {
            "cloudid": "1111111",
            "hostname": "cloudserver-fiGwdwnGPv.mhc",
            "cpu": "2",
            "ram": "4096",
            "managed_os": null,
            "managed_full": null,
            "managed_onetime": null,
            "plan_disksize": 80,
            "disksize": 80,
            "ha": "0",
            "status": "Active",
            "iso": null,
            "ip": "201.201.201.201",
            "billingcycle": "hourly",
            "cost": 19.93,
            "vmcost": 19.93,
            "imagecost": 0,
            "backupcost": 0,
            "hourlycost": 0.03,
            "cloudhourlycost": 0.03,
            "imagehourlycost": 0,
            "backuphourlycost": 0,
            "creditrequired": 0.5,
            "creditreserved": 0,
            "nextinvoiceamount": 0.5,
            "nextinvoicehours": "17",
            "consolepassword": "Ofi4G",
            "powerstatus": "Running",
            "created_at": "2024-05-07 23:03:25",
            "updated_at": "0000-00-00 00:00:00",
            "nextduedate": "2024-05-07",
            "bandwidth": "1000",
            "bandwidth_used": 0,
            "bandwidth_free": 1000,
            "features": {
                "backups": "0"
            },
            "image": {
                "name": "Ubuntu 18.04",
                "distribution": "ubuntu",
                "version": "18.04",
                "image": "ubuntu-18.10-x86_64",
                "cost": "0"
            },
            "dclocation": {
                "location": "Delhi (Noida)",
                "country": "India",
                "dc": "innoida",
                "dccc": "in"
            },
            "v4": {
                "ip_address": "111.111.111.111",
                "netmask": "111.111.111.111",
                "gateway": "111.111.111.111",
                "type": "public",
                "nat": false,
                "primary": "1",
                "rdns": "1031272912.network.microhost.in"
            },
            "networks": {
                "public": {
                    "v4": [
                        {
                            "ip_address": "111.111.111.111",
                            "netmask": "111.111.111.111",
                            "gateway": "111.111.111.111",
                            "type": "public",
                            "nat": false,
                            "primary": "1",
                            "rdns": "1111111.network.microhost.in"
                        }
                    ]
                },
                "private": {
                    "v4": [
                        {
                            "noip": 0
                        }
                    ]
                }
            },
            "v4private": {
                "noip": 0
            },
            "storages": [
                {
                    "id": "1111111111",
                    "size": 80,
                    "disk_used": "0.00",
                    "disk_free": "0.00",
                    "disk_usedp": "0.00",
                    "created_at": null,
                    "bus": "virtio",
                    "type": "Primary"
                }
            ],
            "storage": {
                "id": "1111111",
                "size": 80,
                "disk_used": "0.00",
                "disk_free": "0.00",
                "disk_usedp": "0.00",
                "created_at": null,
                "bus": "virtio",
                "type": "Primary"
            },
            "disk_used": 0,
            "disk_free": 0,
            "disk_usedp": 0,
            "backups": [],
            "snapshots": [],
            "firewalls": [],
            "gpu_available": "0",
            "gpus": []
        },
        {
            "cloudid": "1111122",
            "hostname": "cloudserver-fiGwdwnGPv.mhc",
            "cpu": "2",
            "ram": "4096",
            "managed_os": null,
            "managed_full": null,
            "managed_onetime": null,
            "plan_disksize": 80,
            "disksize": 80,
            "ha": "0",
            "status": "Active",
            "iso": null,
            "ip": "201.201.201.201",
            "billingcycle": "hourly",
            "cost": 19.93,
            "vmcost": 19.93,
            "imagecost": 0,
            "backupcost": 0,
            "hourlycost": 0.03,
            "cloudhourlycost": 0.03,
            "imagehourlycost": 0,
            "backuphourlycost": 0,
            "creditrequired": 0.5,
            "creditreserved": 0,
            "nextinvoiceamount": 0.5,
            "nextinvoicehours": "17",
            "consolepassword": "Ofi4G",
            "powerstatus": "Running",
            "created_at": "2024-05-07 23:03:25",
            "updated_at": "0000-00-00 00:00:00",
            "nextduedate": "2024-05-07",
            "bandwidth": "1000",
            "bandwidth_used": 0,
            "bandwidth_free": 1000,
            "features": {
                "backups": "0"
            },
            "image": {
                "name": "Ubuntu 18.04",
                "distribution": "ubuntu",
                "version": "18.04",
                "image": "ubuntu-18.10-x86_64",
                "cost": "0"
            },
            "dclocation": {
                "location": "Delhi (Noida)",
                "country": "India",
                "dc": "innoida",
                "dccc": "in"
            },
            "v4": {
                "ip_address": "111.111.111.111",
                "netmask": "111.111.111.111",
                "gateway": "111.111.111.111",
                "type": "public",
                "nat": false,
                "primary": "1",
                "rdns": "1031272912.network.microhost.in"
            },
            "networks": {
                "public": {
                    "v4": [
                        {
                            "ip_address": "111.111.111.111",
                            "netmask": "111.111.111.111",
                            "gateway": "111.111.111.111",
                            "type": "public",
                            "nat": false,
                            "primary": "1",
                            "rdns": "1111111.network.microhost.in"
                        }
                    ]
                },
                "private": {
                    "v4": [
                        {
                            "noip": 0
                        }
                    ]
                }
            },
            "v4private": {
                "noip": 0
            },
            "storages": [
                {
                    "id": "1111111111",
                    "size": 80,
                    "disk_used": "0.00",
                    "disk_free": "0.00",
                    "disk_usedp": "0.00",
                    "created_at": null,
                    "bus": "virtio",
                    "type": "Primary"
                }
            ],
            "storage": {
                "id": "1111111",
                "size": 80,
                "disk_used": "0.00",
                "disk_free": "0.00",
                "disk_usedp": "0.00",
                "created_at": null,
                "bus": "virtio",
                "type": "Primary"
            },
            "disk_used": 0,
            "disk_free": 0,
            "disk_usedp": 0,
            "backups": [],
            "snapshots": [],
            "firewalls": [],
            "gpu_available": "0",
            "gpus": []
        },
    ]
}`

const dummyListCloudInstanceServerRes = `[
	{
		"cloudid": "1111111",
		"hostname": "cloudserver-fiGwdwnGPv.mhc",
		"cpu": "2",
		"ram": "4096",
		"managed_os": null,
		"managed_full": null,
		"managed_onetime": null,
		"plan_disksize": 80,
		"disksize": 80,
		"ha": "0",
		"status": "Active",
		"iso": null,
		"ip": "201.201.201.201",
		"billingcycle": "hourly",
		"cost": 19.93,
		"vmcost": 19.93,
		"imagecost": 0,
		"backupcost": 0,
		"hourlycost": 0.03,
		"cloudhourlycost": 0.03,
		"imagehourlycost": 0,
		"backuphourlycost": 0,
		"creditrequired": 0.5,
		"creditreserved": 0,
		"nextinvoiceamount": 0.5,
		"nextinvoicehours": "17",
		"consolepassword": "Ofi4G",
		"powerstatus": "Running",
		"created_at": "2024-05-07 23:03:25",
		"updated_at": "0000-00-00 00:00:00",
		"nextduedate": "2024-05-07",
		"bandwidth": "1000",
		"bandwidth_used": 0,
		"bandwidth_free": 1000,
		"features": {
			"backups": "0"
		},
		"image": {
			"name": "Ubuntu 18.04",
			"distribution": "ubuntu",
			"version": "18.04",
			"image": "ubuntu-18.10-x86_64",
			"cost": "0"
		},
		"dclocation": {
			"location": "Delhi (Noida)",
			"country": "India",
			"dc": "innoida",
			"dccc": "in"
		},
		"v4": {
			"ip_address": "111.111.111.111",
			"netmask": "111.111.111.111",
			"gateway": "111.111.111.111",
			"type": "public",
			"nat": false,
			"primary": "1",
			"rdns": "1031272912.network.microhost.in"
		},
		"networks": {
			"public": {
				"v4": [
					{
						"ip_address": "111.111.111.111",
						"netmask": "111.111.111.111",
						"gateway": "111.111.111.111",
						"type": "public",
						"nat": false,
						"primary": "1",
						"rdns": "1111111.network.microhost.in"
					}
				]
			},
			"private": {
				"v4": [
					{
						"noip": 0
					}
				]
			}
		},
		"v4private": {
			"noip": 0
		},
		"storages": [
			{
				"id": "1111111111",
				"size": 80,
				"disk_used": "0.00",
				"disk_free": "0.00",
				"disk_usedp": "0.00",
				"created_at": null,
				"bus": "virtio",
				"type": "Primary"
			}
		],
		"storage": {
			"id": "1111111",
			"size": 80,
			"disk_used": "0.00",
			"disk_free": "0.00",
			"disk_usedp": "0.00",
			"created_at": null,
			"bus": "virtio",
			"type": "Primary"
		},
		"disk_used": 0,
		"disk_free": 0,
		"disk_usedp": 0,
		"backups": [],
		"snapshots": [],
		"firewalls": [],
		"gpu_available": "0",
		"gpus": []
	},
	{
		"cloudid": "1111122",
		"hostname": "cloudserver-fiGwdwnGPv.mhc",
		"cpu": "2",
		"ram": "4096",
		"managed_os": null,
		"managed_full": null,
		"managed_onetime": null,
		"plan_disksize": 80,
		"disksize": 80,
		"ha": "0",
		"status": "Active",
		"iso": null,
		"ip": "201.201.201.201",
		"billingcycle": "hourly",
		"cost": 19.93,
		"vmcost": 19.93,
		"imagecost": 0,
		"backupcost": 0,
		"hourlycost": 0.03,
		"cloudhourlycost": 0.03,
		"imagehourlycost": 0,
		"backuphourlycost": 0,
		"creditrequired": 0.5,
		"creditreserved": 0,
		"nextinvoiceamount": 0.5,
		"nextinvoicehours": "17",
		"consolepassword": "Ofi4G",
		"powerstatus": "Running",
		"created_at": "2024-05-07 23:03:25",
		"updated_at": "0000-00-00 00:00:00",
		"nextduedate": "2024-05-07",
		"bandwidth": "1000",
		"bandwidth_used": 0,
		"bandwidth_free": 1000,
		"features": {
			"backups": "0"
		},
		"image": {
			"name": "Ubuntu 18.04",
			"distribution": "ubuntu",
			"version": "18.04",
			"image": "ubuntu-18.10-x86_64",
			"cost": "0"
		},
		"dclocation": {
			"location": "Delhi (Noida)",
			"country": "India",
			"dc": "innoida",
			"dccc": "in"
		},
		"v4": {
			"ip_address": "111.111.111.111",
			"netmask": "111.111.111.111",
			"gateway": "111.111.111.111",
			"type": "public",
			"nat": false,
			"primary": "1",
			"rdns": "1031272912.network.microhost.in"
		},
		"networks": {
			"public": {
				"v4": [
					{
						"ip_address": "111.111.111.111",
						"netmask": "111.111.111.111",
						"gateway": "111.111.111.111",
						"type": "public",
						"nat": false,
						"primary": "1",
						"rdns": "1111111.network.microhost.in"
					}
				]
			},
			"private": {
				"v4": [
					{
						"noip": 0
					}
				]
			}
		},
		"v4private": {
			"noip": 0
		},
		"storages": [
			{
				"id": "1111111111",
				"size": 80,
				"disk_used": "0.00",
				"disk_free": "0.00",
				"disk_usedp": "0.00",
				"created_at": null,
				"bus": "virtio",
				"type": "Primary"
			}
		],
		"storage": {
			"id": "1111111",
			"size": 80,
			"disk_used": "0.00",
			"disk_free": "0.00",
			"disk_usedp": "0.00",
			"created_at": null,
			"bus": "virtio",
			"type": "Primary"
		},
		"disk_used": 0,
		"disk_free": 0,
		"disk_usedp": 0,
		"backups": [],
		"snapshots": [],
		"firewalls": [],
		"gpu_available": "0",
		"gpus": []
	},
]`

const dummyListOsImagesRes = `[
	{
		"distro": null,
		"distribution": "Alma Linux",
		"version": "9.0 x86_64",
		"image": "almalinux-9.2-x86_64",
		"cost": 0
	},
	{
		"distro": null,
		"distribution": "Alma Linux",
		"version": "9.2 x86_64",
		"image": "almalinux-9.2-x86_64",
		"cost": 0
	}
]`

const dummyListOsImagesServerRes = `{
    "images": [
        {
            "distro": null,
            "distribution": "Alma Linux",
            "version": "9.0 x86_64",
            "image": "almalinux-9.2-x86_64",
            "cost": 0
        },
        {
            "distro": null,
            "distribution": "Alma Linux",
            "version": "9.2 x86_64",
            "image": "almalinux-9.2-x86_64",
            "cost": 0
        }
    ]
}
`

const dummyListResizePlansRes = `[
        {
            "id": "10027",
            "type": "cloud",
            "disk": "80",
            "ram": "4096",
            "cpu": "2",
            "bandwidth": "1000",
            "slug": "dedicated-cpu",
            "price": 30.59,
            "monthly": 0.41,
            "plantype": "onlyram"
        },
        {
            "id": "10028",
            "type": "cloud",
            "disk": "80",
            "ram": "8192",
            "cpu": "4",
            "bandwidth": "2000",
            "slug": "dedicated-cpu",
            "price": 61.25,
            "monthly": 0.82,
            "plantype": "onlyram"
        }
    ]`

const dummyListResizePlansServerRes = `
{
    "plans": [
        {
            "id": "10027",
            "type": "cloud",
            "disk": "80",
            "ram": "4096",
            "cpu": "2",
            "bandwidth": "1000",
            "slug": "dedicated-cpu",
            "price": 30.59,
            "monthly": 0.41,
            "plantype": "onlyram"
        },
        {
            "id": "10028",
            "type": "cloud",
            "disk": "80",
            "ram": "8192",
            "cpu": "4",
            "bandwidth": "2000",
            "slug": "dedicated-cpu",
            "price": 61.25,
            "monthly": 0.82,
            "plantype": "onlyram"
        }
    ]
}`

const dummyCreateResetPasswordResponseJson = `{
    "password": "fegtrhuy1234",
    "status": "success",
    "message": "success"
}`
