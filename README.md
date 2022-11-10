## Config Template (for GoLang projects)

About
I happen to use .json configs in my project, which I would like to have updated with ability of rollbacking. 
So this project is dedicated purely for this task.


### Features, which now are one-liners

- Load, Dump
- Backup
- Rollback
- Query-like update, i.e. `util.Update("substruct", "field", 42)`

> Note: by default backup file is called `data/config.json` --> `data/.backup_config.json`


### Example 

```go
func main() {
    util, err := NewConfigUtil("config.json")
    if err != nil {...}
    util.Backup()
	
    err = util.Update("example", "easy")
    if err != nil {...}

    err = util.Update("missing", "field", "oh no")
    if err != nil {
        log.Print(err)
        util.Rollback()
    }
}
```


---
### Caveats

Generics in GoLang lack a lot of features provided in Rust/C++, so purely for simplicity (I made this in like an hour)
this project don't use generics and assumes you are ok with copy-pasting the code, as long as you have ~one type 
of configs in your project.
