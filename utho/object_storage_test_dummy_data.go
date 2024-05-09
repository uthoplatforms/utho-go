package utho

const dummyReadBucketRes = `{
    "name": "examplename",
    "access": "download",
    "dcslug": "innoida",
    "size": "10120",
    "status": "Active",
    "created_at": "2024-05-05 19:48:57",
    "object_count": "2",
    "version_enabled": false,
    "current_size": "15.51 KB",
    "permissions": [
        {
            "bucket": "apibukcet",
            "name": "apibukcetkey22",
            "accesskey": "caSu9smRk4eW3fZUpoiGTYwDBzOqhdxjK2lI",
            "permission": "write",
            "status": null,
            "created_at": "2024-05-05 20:14:40"
        }
    ],
    "dclocation": {
        "location": "Delhi (Noida)",
        "country": "India",
        "dc": "innoida",
        "dccc": "in"
    }
}
`

const dummyReadBucketServerRes = `{
    "buckets": [
        ` + dummyReadBucketRes + `
    ]
}
`
const dummyReadAccessKeyRes = `{
    "name": "access-key1",
    "accesskey": "caSu9sqwrhgfewqqqqqOqhdxjK2lI",
    "dcslug": "innoida",
    "status": "Active",
    "created_at": "2024-05-05 19:54:20"
}
`

const dummyReadAccessKeyServerRes = `{
    "accesskeys": [
        ` + dummyReadAccessKeyRes + `
    ]
}
`

const dummyCreateAccessKeyResponseJson = `{
    "status": "success",
    "message": "Object User has been created",
    "accesskey": "qqqqqqqqqqt1VaXSIsDPRM0hpWBx4wQrb",
    "secretkey": "qqqqqqqqqqUJwrVyFh0v4tb2X6locRu8O"
}
`

const dummyListBucketObjectsAndDirectoriesRes = `{
    "size": "1233",
    "type": "directory",
    "name": "pathname"
}`

const dummyListBucketObjectsAndDirectoriesServerRes = `{
    "objects": [
        ` + dummyListBucketObjectsAndDirectoriesRes + `
    ],
    "path":"/pathname"
}`

const dummyGetSharableUrlOfObjectRes = `{
    "url": "example.com/file.name"
}
`

const dummyGetSharableUrlOfObjectServerRes = `{
    "url": "example.com/file.name"
}
`

const dummyListSubscriptionPlanPricingRes = `[
    {
        "id": "10120",
        "uuid": "25GB",
        "type": "objectstorage",
        "slug": "",
        "name": "25 GB",
        "description": "",
        "disk": "25",
        "ram": "0",
        "cpu": "0",
        "bandwidth": "0",
        "is_featured": "0",
        "dedicated_vcore": "0",
        "price": "125.00",
        "monthly": "125.00"
    }
]
`

const dummyListSubscriptionPlanPricingServerRes = `{
    "pricing": [
        {
            "id": "10120",
            "uuid": "25GB",
            "type": "objectstorage",
            "slug": "",
            "name": "25 GB",
            "description": "",
            "disk": "25",
            "ram": "0",
            "cpu": "0",
            "bandwidth": "0",
            "is_featured": "0",
            "dedicated_vcore": "0",
            "price": "125.00",
            "monthly": "125.00"
        }
    ]
}
`
