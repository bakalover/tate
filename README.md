## Tate - structured concurrency

Never write `go` statement again!


- Manual gorroutine handle:
```golang
h := tate.Go(func(){})
h.Join() // Ok
h.Join() // panic! -> Double join
```


- Fixed Scope:
```golang
tate.FixScope(func(s *Scope){
    s.Go(func(){})
    s.Go(func(){})
    s.Go(func(){})
}) // Synchronous wait of all goroutines
```

- Dynamic Scope:
```golang
j := tate.DynScope(func(s *Scope){
    s.Go(func(){})
    s.Go(func(){})
    s.Go(func(){})
}) 

j.Join() // Only at that point all goroutines are done
j.Join() // panic! -> Double join

```
- Repeater
```golang
rp := tate.NewRepeater()
// Repeats in infinite cycle each routine
rp.Go(func(){})
rp.Go(func(){})
// Instantly cancels all cycles and joins goroutines
rp.CancelJoin() 

// We can do it again and again
rp.Go(func(){})
rp.Go(func(){})
rp.CancelJoin() 
```
## Todo
+ Tree Design
+ Inject Cancellation
+ Maybe some functional style

## Reference
+ [Go statement considered harmful](https://vorpus.org/blog/notes-on-structured-concurrency-or-go-statement-considered-harmful/)
+ [Tree structure](https://blog.yoshuawuyts.com/tree-structured-concurrency/#pattern-managed-background-tasks)
+ [Rust scope](https://doc.rust-lang.org/std/thread/fn.scope.html)