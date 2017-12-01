# lamp-life-line

**Under Construction**

## Support REST Calls

### Register New Cluster
Creates a new cluster with the given `name`.

    Endpoint: `/cluster`

    Method: POST

    JSON: 

    ```
    {
        "name": "Cluster Name"
    }
    ```

### Get Cluster(s)
Returns a list of clusters if no `id` is present in the json. If `id` is present return a cluster associated with the `id`.

    Endpoint: `/cluster`

    Method: GET

    JSON:

    ```
    {
        "id": "uuid_of_cluster"
    }
    ```

### Unregister Cluster
Removes a cluster with the given `id`.

    Endpoint: `/cluster`

    Method: DELETE

    JSON:

    ```
    {
        "id": "uuid_of_cluster"
    }
    ```

### Register New Lamp
Creates a new lamp with the given `id` and registers to cluster matching `clusterId`.

    Endpoint: `/lamp`

    Method: POST

    JSON:

    ```
    {
        "id": "lamp_mac_address",
        "clusterId": "uuid_of_cluster"
    }
    ```

### Update Lamp
Can be used to change the cluster a lamp with `id` belongs to.

    Endpoint: `/lamp`

    Method: PUT

    JSON:

    ```
    {
        "id": "lamp_mac_address",
        "clusterId": "uuid_of_cluster"
    }
    ```