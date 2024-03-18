## Tate - structured concurrency
- Manual gorroutine handle:
```golang
h := tate.Go(func())
h.Join()
```


- Scope with SubScopes:
```golang
sc := tate.NewScope()
sc.Go(func())
sc.SubScope(func(sub *Scope){
    sub.Go(func())
})
sc.Join()
```

- Tree Nursery (uses simple Go or scopes)

```golang
nr := tate.NewNursery()
nr.Go(func())
// Scopes and inner subscopes
nr.Scope(func(sc *Scope))
nr.Join() 
```
- Repeater
```golang
rp := tate.NewRepeater()
// Repeats in infinite cycle each routine
rp.Go(func())
rp.Go(func())
// Instantly cancels all cycles and joins routines
rp.Join() 
```


## Reference
+ [Go statement considered harmful](https://vorpus.org/blog/notes-on-structured-concurrency-or-go-statement-considered-harmful/)
+ [Tree structure](https://blog.yoshuawuyts.com/tree-structured-concurrency/#pattern-managed-background-tasks)
+ [Rust scope](https://doc.rust-lang.org/std/thread/fn.scope.html)