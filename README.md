# lamp-life-line

**Under Construction**

## Support REST Calls

### Register New Cluster
Creates a new cluster with the given `name`.

    Endpoint: /cluster

    Method: POST

    Request JSON: 

    {
        "name": "Cluster Name",
        "color": 32_bit_color_int
    }

    Response JSON:

    {
        "id": "uuid_of_cluster",
        "name": "Cluster Name",
        "color": 32_bit_color_int
    }

### Get Cluster(s)
Returns a list of clusters if no `id` is present in the json. If `id` is present return a cluster associated with the `id`.

    Endpoint: /cluster

    Method: GET

    Request JSON:

    {
        "id": "uuid_of_cluster"
    }

    Response JSON:

    {
        "id": "uuid_of_cluster",
        "name": "Cluster Name",
        "color": 32_bit_color_int
    }

### Unregister Cluster
Removes a cluster with the given `id`.

    Endpoint: /cluster

    Method: DELETE

    Request JSON:

    {
        "id": "uuid_of_cluster"
    }

### Change Color
Changes the color for a cluster with a given `id`.

    Endpoint: /color

    Method: PUT

    Request JSON:

    {
        "id": "uuid_of_cluster",
        "color": 32_bit_color_int
    }

    Response JSON:

    {
        "id": "uuid_of_cluster",
        "name": "Cluster Name",
        "color": 32_bit_color_int
    }

### Get Color
Retrieves the color for a cluster with a given `id`.

    Endpoint: /color

    Method: GET

    Request JSON:

    {
        "id": "uuid_of_cluster"
    }

    Response JSON:

    {
        "color": 32_bit_color_int
    }