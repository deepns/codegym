# Client side load balancing

gRPC supports two types of load balancing policies on the client side:

1. PickFirst: The client will connect to the first server that it resolves.
2. RoundRobin: The client will connect to the servers in a round robin fashion.

The default load balancing policy is PickFirst.
To use the RoundRobin load balancing policy, we need to set the default service config.
The default service config is a JSON string that contains the load balancing policy.

```go
	fooConn, err := grpc.Dial(service,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
```

With round robin, client connects to all addresses it sees and sends RPC to each
backend address one at a time. If the connection is not ready for any reason, client
may send RPCs to the same backend address multiple times.

- The sample server takes a comma separated list of ports on the input and starts a gRPC server on each of the ports.
- Similar client takes a comma separated list of addresses to connect to. Client then uses the sample resolver (FooResolver) added in the [name_resolving example](../name_resolving/resolver/resolver.go) to resolve the sample service `foo:///resolver.foo.bar` to these backend addresses in a round robin fashion.

- Here is a sample output with debugging logs enabled on the client side. Can see that sometimes RPCs are sent to same server consecutively, depending on the connection readiness.

```console
 ✗ go run server/main.go --ports 50505,50506,50507
2023/04/25 21:17:29 Listening on port 50505
2023/04/25 21:17:29 Listening on port 50506
2023/04/25 21:17:29 Listening on port 50507
2023/04/25 21:17:32 ServerPort:50507 UnaryEcho: foo:///resolver.foo.bar
2023/04/25 21:17:32 ServerPort:50507 UnaryEcho: foo:///resolver.foo.bar
2023/04/25 21:17:32 ServerPort:50505 UnaryEcho: foo:///resolver.foo.bar
2023/04/25 21:17:32 ServerPort:50506 UnaryEcho: foo:///resolver.foo.bar
2023/04/25 21:17:32 ServerPort:50507 UnaryEcho: foo:///resolver.foo.bar
2023/04/25 21:17:32 ServerPort:50505 UnaryEcho: foo:///resolver.foo.bar
2023/04/25 21:17:32 ServerPort:50506 UnaryEcho: foo:///resolver.foo.bar
2023/04/25 21:17:32 ServerPort:50507 UnaryEcho: foo:///resolver.foo.bar
2023/04/25 21:17:32 ServerPort:50505 UnaryEcho: foo:///resolver.foo.bar
2023/04/25 21:17:32 ServerPort:50506 UnaryEcho: foo:///resolver.foo.bar

✗ GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info go run client/main.go --addr localhost:50507,localhost:50505,localhost:50506

2023/04/25 21:30:17 INFO: [core] [Channel #1] Channel created
2023/04/25 21:30:17 INFO: [core] [Channel #1] original dial target is: "foo:///resolver.foo.bar"
2023/04/25 21:30:17 INFO: [core] [Channel #1] parsed dial target is: {Scheme:foo Authority: URL:{Scheme:foo Opaque: User: Host: Path:/resolver.foo.bar RawPath:
 OmitHost:false ForceQuery:false RawQuery: Fragment: RawFragment:}}
2023/04/25 21:30:17 INFO: [core] [Channel #1] Channel authority set to "resolver.foo.bar"
2023/04/25 21:30:17 INFO: [core] [Channel #1] Resolver state updated: {
  "Addresses": [
    {
      "Addr": "localhost:50507",
      "ServerName": "",
      "Attributes": null,
      "BalancerAttributes": null,
      "Type": 0,
      "Metadata": null
    },
    {
      "Addr": "localhost:50505",
      "ServerName": "",
      "Attributes": null,
      "BalancerAttributes": null,
      "Type": 0,
      "Metadata": null
    },
    {
      "Addr": "localhost:50506",
      "ServerName": "",
      "Attributes": null,
      "BalancerAttributes": null,
      "Type": 0,
      "Metadata": null
    }
  ],
  "ServiceConfig": null,
  "Attributes": null
} (resolver returned new addresses)
2023/04/25 21:30:17 INFO: [balancer] base.baseBalancer: got new ClientConn state:  {{[{                                                              [68/11852]
  "Addr": "localhost:50507",
  "ServerName": "",
  "Attributes": null,
  "BalancerAttributes": null,
  "Type": 0,
  "Metadata": null
} {
  "Addr": "localhost:50505",
  "ServerName": "",
  "Attributes": null,
  "BalancerAttributes": null,
  "Type": 0,
  "Metadata": null
} {
  "Addr": "localhost:50506",
  "ServerName": "",
  "Attributes": null,
  "BalancerAttributes": null,
  "Type": 0,
  "Metadata": null
}] <nil> <nil>} <nil>}
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #2] Subchannel created
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #3] Subchannel created
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #4] Subchannel created
2023/04/25 21:30:17 INFO: [roundrobin] roundrobinPicker: Build called with info: {map[]}
2023/04/25 21:30:17 INFO: [core] [Channel #1] Channel Connectivity change to CONNECTING
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #2] Subchannel Connectivity change to CONNECTING
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #3] Subchannel Connectivity change to CONNECTING
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #2] Subchannel picks a new address "localhost:50507" to connect
2023/04/25 21:30:17 INFO: [balancer] base.baseBalancer: handle SubConn state change: 0x140000ae810, CONNECTING
2023/04/25 21:30:17 INFO: [balancer] base.baseBalancer: handle SubConn state change: 0x140000ae828, CONNECTING
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #3] Subchannel picks a new address "localhost:50505" to connect
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #4] Subchannel Connectivity change to CONNECTING
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #4] Subchannel picks a new address "localhost:50506" to connect
2023/04/25 21:30:17 INFO: [balancer] base.baseBalancer: handle SubConn state change: 0x140000ae840, CONNECTING
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #4] Subchannel Connectivity change to READY
2023/04/25 21:30:17 INFO: [balancer] base.baseBalancer: handle SubConn state change: 0x140000ae840, READY
2023/04/25 21:30:17 INFO: [roundrobin] roundrobinPicker: Build called with info: {map[0x140000ae840:{{
  "Addr": "localhost:50506",                                                                                                                         [29/11852]
  "ServerName": "",
  "Attributes": null,
  "BalancerAttributes": null,
  "Type": 0,
  "Metadata": null
}}]}
2023/04/25 21:30:17 INFO: [core] [Channel #1] Channel Connectivity change to READY
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #2] Subchannel Connectivity change to READY
2023/04/25 21:30:17 INFO: [balancer] base.baseBalancer: handle SubConn state change: 0x140000ae810, READY
2023/04/25 21:30:17 INFO: [roundrobin] roundrobinPicker: Build called with info: {map[0x140000ae810:{{
  "Addr": "localhost:50507",
  "ServerName": "",
  "Attributes": null,
  "BalancerAttributes": null,
  "Type": 0,
  "Metadata": null
}} 0x140000ae840:{{
  "Addr": "localhost:50506",
  "ServerName": "",
  "Attributes": null,
  "BalancerAttributes": null,
  "Type": 0,
  "Metadata": null
}}]}
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #3] Subchannel Connectivity change to READY
2023/04/25 21:30:17 INFO: [balancer] base.baseBalancer: handle SubConn state change: 0x140000ae828, READY
2023/04/25 21:30:17 INFO: [roundrobin] roundrobinPicker: Build called with info: {map[0x140000ae810:{{
  "Addr": "localhost:50507",
  "ServerName": "",
  "Attributes": null,
  "BalancerAttributes": null,
  "Type": 0,
  "Metadata": null
}} 0x140000ae828:{{
  "Addr": "localhost:50505",
  "ServerName": "",
  "Attributes": null,
  "BalancerAttributes": null,
  "Type": 0,
  "Metadata": null
}} 0x140000ae828:{{
  "Addr": "localhost:50505",
  "ServerName": "",
  "Attributes": null,
  "BalancerAttributes": null,
  "Type": 0,
  "Metadata": null
}} 0x140000ae840:{{
  "Addr": "localhost:50506",
  "ServerName": "",
  "Attributes": null,
  "BalancerAttributes": null,
  "Type": 0,
  "Metadata": null
}}]}
2023/04/25 21:30:17 echo(0): foo:///resolver.foo.bar
2023/04/25 21:30:17 echo(1): foo:///resolver.foo.bar
2023/04/25 21:30:17 echo(2): foo:///resolver.foo.bar
2023/04/25 21:30:17 echo(3): foo:///resolver.foo.bar
2023/04/25 21:30:17 echo(4): foo:///resolver.foo.bar
2023/04/25 21:30:17 echo(5): foo:///resolver.foo.bar
2023/04/25 21:30:17 echo(6): foo:///resolver.foo.bar
2023/04/25 21:30:17 echo(7): foo:///resolver.foo.bar
2023/04/25 21:30:17 echo(8): foo:///resolver.foo.bar
2023/04/25 21:30:17 echo(9): foo:///resolver.foo.bar
2023/04/25 21:30:17 INFO: [core] [Channel #1] Channel Connectivity change to SHUTDOWN
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #2] Subchannel Connectivity change to SHUTDOWN
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #2] Subchannel deleted
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #3] Subchannel Connectivity change to SHUTDOWN
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #3] Subchannel deleted
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #4] Subchannel Connectivity change to SHUTDOWN
2023/04/25 21:30:17 INFO: [core] [Channel #1 SubChannel #4] Subchannel deleted
2023/04/25 21:30:17 INFO: [core] [Channel #1] Channel deleted
```
