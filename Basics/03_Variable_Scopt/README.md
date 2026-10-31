## VARIABLE SCOPE
- Block
    - Inner blocks can access variables declared within outer block

    - Outer block cannot access variables declared within inner block

```
{
    // outer block
    {
        // inner block
    }
}
```

### Local vs Global

#### Local
- Declared inside function/block
- Not accessible outside block
- can be declared in conditional looping block

#### Global
- Declared outside the function, like where package and import