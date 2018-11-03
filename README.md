# lamp-life-line

**Under Construction**

## Support REST Calls

### Register New Cluster
Creates a new cluster with the given `CLUSTER_NAME`.

    Endpoint: /cluster

    Method: POST

    Request JSON: 

    {
        "name": "CLUSTER_NAME",
        "color": 32_bit_color_int
    }

    Response JSON:

    {
        "name": "CLUSTER_NAME",
        "color": 32_bit_color_int
    }

### Get Cluster
Returns a list of clusters.

    Endpoint: /cluster

    Method: GET

    Response JSON:
    [
        {
            "name": "CLUSTER_NAME",
            "color": 32_bit_color_int
        }
    ]

### Get Cluster
Returns a clusters with the given `CLUSTER_NAME`.

    Endpoint: /cluster/{{CLUSTER_NAME}}

    Method: GET

    Response JSON:

    {
        "name": "CLUSTER_NAME",
        "color": 32_bit_color_int
    }

### Unregister Cluster
Removes a cluster with the given `CLUSTER_NAME`.

    Endpoint: /cluster/{{CLUSTER_NAME}}

    Method: DELETE

### Change Color
Changes the color for a cluster with a given `CLUSTER_NAME`.

    Endpoint: /cluster/{{CLUSTER_NAME}}/color

    Method: PUT

    Request JSON:

    {
        "color": 32_bit_color_int
    }

    Response JSON:

    {
        "color": 32_bit_color_int
    }

### Get Color
Retrieves the color for a cluster with a given `CLUSTER_NAME`.

    Endpoint: /cluster/{{CLUSTER_NAME}}/color

    Method: GET

    Response JSON:

    {
        "color": 32_bit_color_int
    }