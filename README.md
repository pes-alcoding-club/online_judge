# Online Judge

This creates a wrapper around the Judge0 API. The wrapper is created using GoFibre. 

To run the server, execute:

```sh
go run main.go
```

Example of a payload that can be passed:
```sh
curl localhost:8080/submission -d "{ \"language_id\": 50, \"source_code\": \"#include <stdio.h>\\n\\nint main(void) {\\n  char name[10];\\n  scanf(\\\"%s\\\", name);\\n  printf(\\\"hello %s\\\\n\\\", name);\\n  return 0;\\n}\", \"stdin\": \"world\"}" -v
```
