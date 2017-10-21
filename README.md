# lamp-life-line

**Under Construction**

# Support REST Calls

## Register New Cluster
Creates a new cluster with the given `name`.

    Method: POST

    JSON: 

    ```
    {
        "name": "Cluster Name"
    }
    ```

## Get Cluster
Returns a list of clusters if no `id` is present in the json. If `id` is present return a cluster associated with the `id`.

    Method: GET

    JSON:

    ```
    {
        "id": "uuid_of_cluster"
    }
    ```