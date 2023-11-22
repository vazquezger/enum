# Pipelined functions for go slices

Note: yeah, yeah, I know it is not idiomatic.... it was just for fun :)

This code:
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

... is equivalent to this code:
```

  input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  input2 := make([]int, 0)
  for i := range input {
    if input[i] % 2 == 0 {
      input2 = append(input2, input[i])
    }
  }

  input3 := make([]int, len(input2))
  for i := range input2 {
    input3[i] = input3[i] * 2
  }

  for i := range input3 {
    if input3[i] == 12 {
       fmt.Println(12)
       // 12
       break
    }
  }
  
```
  
