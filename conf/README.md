#Config lib and example

Package config is light and loads easily configuration stored as JSON

```
{
    "maxProc" : 10,
    "path" : {
        "tmp" : {
            "linux" : "/tmp"
        }
    }
}
```

```
	var cfg config.Config
	cfg.LoadJSON("config.json")
	fmt.Println(cfg.Get("path.tmp.linux"))
```        
