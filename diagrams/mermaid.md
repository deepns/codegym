# Learning Mermaid

## Flowchart

### A top down flowchart

```mermaid
flowchart TD;
    A[Start] --> B[Do This];
    B --> C[Do that];
    C --> D[Done]
```

### Left-to-right flowchart

```mermaid
flowchart LR;
    A[Open VSCode] --> B[install Markdown Preview for Mermaid Extn];
    B --> C[Write mermaid];
    C --> D[Preview]

```

### Right-to-left flowchart

Also support markdown texts within the nodes

- codespan and strikethrough are not supported

```mermaid
flowchart RL;
    A(**Bold text**) --> B(*Italics*) --> multiline("`Line1
    Line2
    Line3`")

```

- support many different shapes (hexagon, parallelogram, rhombus etc.)

```mermaid
flowchart LR
    id1[(Database)]

```

### Linking nodes

```mermaid
flowchart LR;
    A[Alice]-- send public key ---B[Bob]
```

This also works Use `-.` `.->` for dotted lines 

```mermaid
flowchart LR;
    Alice---|send public key|B
    B-. send response .->A
```

Linking multiple nodes

```mermaid
flowchart LR
   A --> B & C--> D
```

```mermaid
architecture-beta
    group api(cloud)[API]

    service db(database)[Database] in api
    service disk1(disk)[Storage] in api
    service disk2(disk)[Storage] in api
    service server(server)[Server] in api

    db:L -- R:server
    disk1:T -- B:server
    disk2:T -- B:db
```

## Charts

### PieChart

```mermaid
pie title requests by country
    "United States": 3958
    "India": 5003
    "Russia": 304
```
