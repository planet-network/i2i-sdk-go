# i2i-sdk-go
Golang SDK for i2i

# i2i-client
i2i-client is command line utility to interact with i2i instance.
The main purpose of the application is to show possibilities of SDK.

## building i2i-client from sources
This step is not needed since i2i-client is application showing possibilities of SDK.
To manually build i2i-client you must have installed both make and go compiler.

```bash
make i2i-client
```

## running local instance of i2i
One of the ways to run i2i is tu run i2i on local machine.
To do so, you need to run console command:

```bash
i2i-client exec --i2i-path /usr/bin/i2i --name myNode
```

## running hosted instance of i2i
Other way is to order hosted instance of i2i.
Example command to perform action:

```bash
./i2i-client manager quick-order --hosting "https://v2.vphi.io" --plan "manager_plan_free_10d_1GB" --name "test1"
```

With this command i2i-client will automatically generate keychain and unlock node.
