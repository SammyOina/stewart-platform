## Library Modifications
comment out `_client.extraHeaders = WEBSOCKETS_STRING("Origin: file://");` line 32 from the `WebSocketsClient.cpp` located in the `.pio/libdeps` dir 

UNCOMMENT OUT 
```
OFFSETS = (long *)malloc(COUNT * sizeof(long));
	for (int i = 0; i < COUNT; ++i)
	{
		OFFSETS[i] = 0;
	}
```
in `HX711-multi.cpp` line 20-24 located in `.pio/libdeps` dir