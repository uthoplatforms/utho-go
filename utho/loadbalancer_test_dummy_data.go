package utho

const dummyCreateLoadbalancerResponseJson = `{
    "status": "success",
    "loadbalancerid": "123213",
    "message": "success"
}`

const dummyReadLoadbalancerRes = `{
    "id": "11111",
    "userid": "197456",
    "ip": "103.111.88.111",
    "name": "hqwe",
    "algorithm": "",
    "cookie": "0",
    "cookiename": "",
    "redirecthttps": "0",
    "type": "application",
    "country": "India",
    "cc": "in",
    "city": "Mumbai",
    "backendcount": "0",
    "created_at": "2024-05-06 17:00:51",
    "status": "Active",
    "backends": [
        {
            "id": "22222",
            "lb": "262896",
            "ip": "103.111.111.111",
            "cloudid": "127233",
            "name": "cloudserver-fiGwdwnGPv.mhc",
            "ram": "4096",
            "cpu": "2",
            "disk": "80",
            "country": "India",
            "cc": "in",
            "city": "Delhi (Noida)"
        }
    ],
    "rules": [
        {
            "id": "247818",
            "lb": "262894",
            "src_proto": "tcp",
            "src_port": "80",
            "dst_proto": "tcp",
            "dst_port": "80",
            "timeadded": "0",
            "timeupdated": "0"
        }
    ],
    "rule": {
        "id": "247818",
        "lb": "262894",
        "src_proto": "tcp",
        "src_port": "80",
        "dst_proto": "tcp",
        "dst_port": "80",
        "timeadded": "0",
        "timeupdated": "0"
    },
    "acls": [
        {
            "id": "22222",
            "acl_condition": "innoida",
            "name": "example",
            "value": "{'frontend_id':'12324','type':'http_user_agent','data':['value1','value2']}"
        }
    ],
    "routes": [
        {
            "id": "22222",
            "acl_id": "3333",
            "acl_name": "qwer",
            "routing_condition": "true",
            "backend_id": "344555"
        }
    ],
    "scaling_groups": [],
    "frontends": [
        {
            "id": "22222",
            "name": "gfds2",
            "algorithm": "roundrobin",
            "cookie": "1",
            "cookiename": "0",
            "redirecthttps": "1",
            "certificate_id": "0",
            "port": "80",
            "proto": "http",
            "created_at": "2024-05-09 14:11:57",
            "updated_at": "2024-05-09 14:11:57",
            "backends": [
                {
                    "id": "234234525",
                    "ip": "103.209.147.133",
                    "cloudid": "1277662",
                    "status": "active",
                    "scaling_groupid": "0",
                    "backend_port": "43",
                    "frontend_id": "169"
                }
            ],
            "acls": [],
            "routes": [],
            "rules": []
        }
    ]
}
`

const dummyReadLoadbalancerServerRes = `{
    "loadbalancers": [
        {
            "id": "11111",
            "userid": "197456",
            "ip": "103.111.88.111",
            "name": "hqwe",
            "algorithm": "",
            "cookie": "0",
            "cookiename": "",
            "redirecthttps": "0",
            "type": "application",
            "country": "India",
            "cc": "in",
            "city": "Mumbai",
            "backendcount": "0",
            "created_at": "2024-05-06 17:00:51",
            "status": "Active",
            "backends": [
                {
                    "id": "22222",
                    "lb": "262896",
                    "ip": "103.111.111.111",
                    "cloudid": "127233",
                    "name": "cloudserver-fiGwdwnGPv.mhc",
                    "ram": "4096",
                    "cpu": "2",
                    "disk": "80",
                    "country": "India",
                    "cc": "in",
                    "city": "Delhi (Noida)"
                }
            ],
            "rules": [
                {
                    "id": "247818",
                    "lb": "262894",
                    "src_proto": "tcp",
                    "src_port": "80",
                    "dst_proto": "tcp",
                    "dst_port": "80",
                    "timeadded": "0",
                    "timeupdated": "0"
                }
            ],
            "rule": {
                "id": "247818",
                "lb": "262894",
                "src_proto": "tcp",
                "src_port": "80",
                "dst_proto": "tcp",
                "dst_port": "80",
                "timeadded": "0",
                "timeupdated": "0"
            },
            "acls": [
				{
					"id": "22222",
					"acl_condition": "innoida",
					"name": "example",
					"value": "{'frontend_id':'12324','type':'http_user_agent','data':['value1','value2']}"
				}
			],
            "routes": [
                {
                    "id": "22222",
                    "acl_id": "3333",
                    "acl_name": "qwer",
                    "routing_condition": "true",
                    "backend_id": "344555"
                }
            ],
            "scaling_groups": [],
            "frontends": [
                {
                    "id": "22222",
                    "name": "gfds2",
                    "algorithm": "roundrobin",
                    "cookie": "1",
                    "cookiename": "0",
                    "redirecthttps": "1",
                    "certificate_id": "0",
                    "port": "80",
                    "proto": "http",
                    "created_at": "2024-05-09 14:11:57",
                    "updated_at": "2024-05-09 14:11:57",
                    "backends": [
                        {
                            "id": "234234525",
                            "ip": "103.209.147.133",
                            "cloudid": "1277662",
                            "status": "active",
                            "scaling_groupid": "0",
                            "backend_port": "43",
                            "frontend_id": "169"
                        }
                    ],
                    "acls": [],
                    "routes": [],
                    "rules": []
                }
            ]
        },
        {
            "id": "262896",
            "userid": "197456",
            "ip": "103.189.88.182",
            "name": "hgfd",
            "algorithm": "",
            "cookie": "0",
            "cookiename": "",
            "redirecthttps": "0",
            "type": "network",
            "country": "India",
            "cc": "in",
            "city": "Mumbai",
            "backendcount": "0",
            "created_at": "2024-05-06 21:29:13",
            "status": "Active",
            "backends": [],
            "rules": [
                {
                    "id": "247820",
                    "lb": "262896",
                    "src_proto": "tcp",
                    "src_port": "80",
                    "dst_proto": "tcp",
                    "dst_port": "80",
                    "timeadded": "0",
                    "timeupdated": "0"
                }
            ],
            "rule": {
                "id": "247820",
                "lb": "262896",
                "src_proto": "tcp",
                "src_port": "80",
                "dst_proto": "tcp",
                "dst_port": "80",
                "timeadded": "0",
                "timeupdated": "0"
            },
            "acls": [],
            "routes": [],
            "scaling_groups": [],
            "frontends": [
                {
                    "id": "22222",
                    "name": "gfds2",
                    "algorithm": "roundrobin",
                    "cookie": "1",
                    "cookiename": "0",
                    "redirecthttps": "1",
                    "certificate_id": "0",
                    "port": "80",
                    "proto": "http",
                    "created_at": "2024-05-09 14:11:57",
                    "updated_at": "2024-05-09 14:11:57",
                    "backends": [
                        {
                            "id": "234234525",
                            "ip": "103.209.147.133",
                            "cloudid": "1277662",
                            "status": "active",
                            "scaling_groupid": "0",
                            "backend_port": "43",
                            "frontend_id": "169"
                        }
                    ],
                    "acls": [],
                    "routes": [],
                    "rules": []
                }
            ]
        }
    ]
}
`

// const dummyListLoadbalancerServerRes = `[
// 	{
//         "id": "11111",
//         "userid": "197456",
//         "ip": "103.111.88.111",
//         "name": "hqwe",
//         "algorithm": "",
//         "cookie": "0",
//         "cookiename": "",
//         "redirecthttps": "0",
//         "type": "application",
//         "country": "India",
//         "cc": "in",
//         "city": "Mumbai",
//         "backendcount": "0",
//         "created_at": "2024-05-06 17:00:51",
//         "status": "Active",
//         "backends": [
//             {
//                 "id": "22222",
//                 "lb": "262896",
//                 "ip": "103.111.111.111",
//                 "cloudid": "127233",
//                 "name": "cloudserver-fiGwdwnGPv.mhc",
//                 "ram": "4096",
//                 "cpu": "2",
//                 "disk": "80",
//                 "country": "India",
//                 "cc": "in",
//                 "city": "Delhi (Noida)"
//             }
//         ],
//         "rules": [
//             {
//                 "id": "247818",
//                 "lb": "262894",
//                 "src_proto": "tcp",
//                 "src_port": "80",
//                 "dst_proto": "tcp",
//                 "dst_port": "80",
//                 "timeadded": "0",
//                 "timeupdated": "0"
//             }
//         ],
//         "rule": {
//             "id": "247818",
//             "lb": "262894",
//             "src_proto": "tcp",
//             "src_port": "80",
//             "dst_proto": "tcp",
//             "dst_port": "80",
//             "timeadded": "0",
//             "timeupdated": "0"
//         },
//         "acls": [
//             {
//                 "id": "22222",
//                 "acl_condition": "innoida",
//                 "name": "example",
//                 "value": "{'frontend_id':'12324','type':'http_user_agent','data':['value1','value2']}"
//             }
//         ],
//         "routes": [
//             {
//                 "id": "22222",
//                 "acl_id": "3333",
//                 "acl_name": "qwer",
//                 "routing_condition": "true",
//                 "backend_id": "344555"
//             }
//         ],
//         "scaling_groups": [],
//         "frontends": [
//             {
//                 "id": "22222",
//                 "name": "gfds2",
//                 "algorithm": "roundrobin",
//                 "cookie": "1",
//                 "cookiename": "0",
//                 "redirecthttps": "1",
//                 "certificate_id": "0",
//                 "port": "80",
//                 "proto": "http",
//                 "created_at": "2024-05-09 14:11:57",
//                 "updated_at": "2024-05-09 14:11:57",
//                 "backends": [
//                     {
//                         "id": "234234525",
//                         "ip": "103.209.147.133",
//                         "cloudid": "1277662",
//                         "status": "active",
//                         "scaling_groupid": "0",
//                         "backend_port": "43",
//                         "frontend_id": "169"
//                     }
//                 ],
//                 "acls": [],
//                 "routes": [],
//                 "rules": []
//             }
//         ]
//     },
// 	{
// 		"id": "262896",
// 		"userid": "197456",
// 		"ip": "103.189.88.182",
// 		"name": "hgfd",
// 		"algorithm": "",
// 		"cookie": "0",
// 		"cookiename": "",
// 		"redirecthttps": "0",
// 		"type": "network",
// 		"country": "India",
// 		"cc": "in",
// 		"city": "Mumbai",
// 		"backendcount": "0",
// 		"created_at": "2024-05-06 21:29:13",
// 		"status": "Active",
// 		"backends": [],
// 		"rules": [
// 			{
// 				"id": "247820",
// 				"lb": "262896",
// 				"src_proto": "tcp",
// 				"src_port": "80",
// 				"dst_proto": "tcp",
// 				"dst_port": "80",
// 				"timeadded": "0",
// 				"timeupdated": "0"
// 			}
// 		],
// 		"rule": {
// 			"id": "247820",
// 			"lb": "262896",
// 			"src_proto": "tcp",
// 			"src_port": "80",
// 			"dst_proto": "tcp",
// 			"dst_port": "80",
// 			"timeadded": "0",
// 			"timeupdated": "0"
// 		},
// 		"acls": [],
// 		"routes": [],
// 		"scaling_groups": [],
// 		"frontends": []
// 	}
// ]`

const dummyReadLoadbalancAclRes = `{
	"id": "22222",
	"acl_condition": "innoida",
	"name": "example",
	"value": "{'frontend_id':'12324','type':'http_user_agent','data':['value1','value2']}"
}`

const dummyReadLoadbalancFrontendRes = `{
    "id": "22222",
    "name": "gfds2",
    "algorithm": "roundrobin",
    "cookie": "1",
    "cookiename": "0",
    "redirecthttps": "1",
    "certificate_id": "0",
    "port": "80",
    "proto": "http",
    "created_at": "2024-05-09 14:11:57",
    "updated_at": "2024-05-09 14:11:57",
    "backends": [
        {
            "id": "234234525",
            "ip": "103.111.147.111",
            "cloudid": "1277662",
            "status": "active",
            "scaling_groupid": "0",
            "backend_port": "43",
            "frontend_id": "169"
        }
    ],
    "acls": [],
    "routes": [],
    "rules": []
}`

const dummyReadLoadbalancBackendRes = `{
    "id": "22222",
    "lb": "262896",
    "ip": "103.111.111.111",
    "cloudid": "127233",
    "name": "cloudserver-fiGwdwnGPv.mhc",
    "ram": "4096",
    "cpu": "2",
    "disk": "80",
    "country": "India",
    "cc": "in",
    "city": "Delhi (Noida)"
}`

const dummyReadLoadbalancRouteRes = `
{
    "id": "22222",
    "acl_id": "3333",
    "acl_name": "qwer",
    "routing_condition": "true",
    "backend_id": "344555"
}`
