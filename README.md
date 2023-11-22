# Pipelined functions for go slices

Note: yeah, yeah, I know it is not idiomatic.... it was just for fun :)

This code:
```
  input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  output := make([]int, 0)
  for i := range input {
    if input[i] % 2 == 0 {
      output = append(output, input[i]
    }
  }

  for i := range input {
    input[i] = input[i] * 2
  }

  ok := false
  output := ""
  for i := range output {
    if output[i] == 12 {
      ok = true
      output = output[i]
      break
    }
  }

  if ok {
    fmt.Println(output)
    // 12
  }
```

... is equivalent to this code:
```
  input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  output, ok := enum.
  		Of(input).
  		Filter(func(it int) bool { return it%2 == 0 }).
  		Map(func(it int) int { return it*2 }).
  		Find(func(it int) bool { return it == 12 })

  if ok {
    fmt.Println(*output)
    // 12
  }
  
```
  
